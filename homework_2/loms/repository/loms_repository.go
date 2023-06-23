package repository

import (
	"context"
	"fmt"
	"math/rand"
	"route256/loms/internal/protoc/loms"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) CreateOrder(ctx context.Context, req *loms.CreateOrderRequest) (*loms.CreateOrderResponse, error) {

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction createorder: %v", err)
	}
	defer tx.Rollback(ctx)

	orderID, err := r.createOrder(ctx, tx, req.User, req.Item)
	if err != nil {

		return nil, fmt.Errorf("failed to create order createorder: %v", err)
	}

	if err := r.reserveItems(ctx, tx, req.Item, orderID); err != nil {
		return nil, fmt.Errorf("failed to reserve items: %v", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction createorder: %v", err)
	}

	response := &loms.CreateOrderResponse{
		OrderID: orderID,
	}

	return response, nil
}

func (r *Repository) ListOrder(ctx context.Context, req *loms.ListOrderRequest) (*loms.ListOrderResponse, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction ListOrder: %v", err)
	}
	defer tx.Rollback(ctx)
	result := &loms.ListOrderResponse{}
	status, userID, err := r.getListStatusAndUser(ctx, tx, req.OrderID)
	if err != nil {
		return nil, fmt.Errorf("failed to exec listorder query: %v", err)
	}
	result.User = uint64(userID)
	result.Status = status
	result.Items, err = r.getListCartItems(ctx, tx, req.OrderID)

	return result, nil

}

func (r *Repository) OrderPaid(ctx context.Context, req *loms.OrderPayedRequest) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	query := "UPDATE Orders SET status = 'payed' WHERE orderID = $1"
	_, err = tx.Exec(ctx, query, req.OrderID)
	if err != nil {
		return fmt.Errorf("failed to update order status: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	return nil
}

func (r *Repository) CancelOrder(ctx context.Context, req *loms.CancelOrderRequest) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction cancelorder: %v", err)
	}
	defer tx.Rollback(ctx)

	if err := r.cancelOrder(ctx, tx, req.OrderId); err != nil {
		return fmt.Errorf("failed to cancel order: %v", err)
	}

	if err := r.releaseReservedItems(ctx, tx, req.OrderId); err != nil {
		return fmt.Errorf("failed to release reserved items: %v", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction cancelorder: %v", err)
	}

	return nil
}

func (r *Repository) Stocks(ctx context.Context, req *loms.StocksRequest) (*loms.StocksResponse, error) {
	query := "SELECT warehouseID, count FROM Stocks WHERE sku = $1"
	rows, err := r.pool.Query(ctx, query, req.Sku)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var response loms.StocksResponse
	for rows.Next() {
		var warehouseID int64
		var count int64
		if err := rows.Scan(&warehouseID, &count); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		stock := loms.Stock{
			WarehouseID: warehouseID,
			Count:       uint64(count),
		}
		response.Stocks = append(response.Stocks, &stock)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("encountered an error while iterating over rows: %v", err)
	}

	return &response, nil
}

func (r *Repository) createOrder(ctx context.Context, tx pgx.Tx, userID uint64, items []*loms.Items) (int64, error) {
	var orderID int64
	orderID = rand.Int63n(1000) + 1
	query := squirrel.Insert("Orders").
		Columns("orderID", "status", "user_id").
		Values(orderID, "new", userID).
		PlaceholderFormat(squirrel.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build SQL query: %v", err)
	}

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return 0, fmt.Errorf("failed to execute SQL query: %v", err)
	}
	return orderID, nil
}

func (r *Repository) reserveItems(ctx context.Context, tx pgx.Tx, items []*loms.Items, orderID int64) error {
	for _, item := range items {
		insertBuilder := squirrel.Insert("OrderItems").
			Columns("orderID", "sku", "count").
			Values(orderID, item.Sku, item.Count).
			PlaceholderFormat(squirrel.Dollar)
		sql, args, err := insertBuilder.ToSql()
		if err != nil {
			return fmt.Errorf("failed to build SQL query: %v", err)
		}

		_, err = tx.Exec(ctx, sql, args...)
		if err != nil {
			return fmt.Errorf("failed to execute SQL query: %v", err)
		}

	}
	updateBuilder := squirrel.Update("Orders").
		Set("status", "reserved").
		Where(squirrel.Eq{"orderID": orderID})

	sql, args, err := updateBuilder.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query: %v", err)
	}

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed to execute SQL query: %v", err)
	}

	return nil
}

func (r *Repository) updateReservedItemsToPurchased(ctx context.Context, tx pgx.Tx, orderID int64) error {
	updateBuilder := squirrel.Update("Stocks").
		Set("count", squirrel.Expr("count + OrderItems.count")).
		From("OrderItems").
		Where(squirrel.Eq{"Stocks.sku": squirrel.Expr("OrderItems.sku")}).
		Where(squirrel.Eq{"OrderItems.orderID": orderID}).
		PlaceholderFormat(squirrel.Dollar)

	sql, args, err := updateBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query: %v", err)
	}

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed to execute SQL query: %v", err)
	}

	return nil
}

func (r *Repository) cancelOrder(ctx context.Context, tx pgx.Tx, orderID int64) error {
	query := "DELETE FROM Orders WHERE orderID = $1"
	_, err := tx.Exec(ctx, query, orderID)
	if err != nil {
		return fmt.Errorf("failed to cancel order: %v", err)
	}
	query1 := "DELETE FROM OrderItems WHERE orderID = $1"
	_, err = tx.Exec(ctx, query1, orderID)
	if err != nil {
		return fmt.Errorf("failed to delete order items: %v", err)
	}

	return nil
}

func (r *Repository) releaseReservedItems(ctx context.Context, tx pgx.Tx, orderID int64) error {
	selectBuilder := squirrel.Select("sku", "count").
		From("OrderItems").
		Where(squirrel.Eq{"orderID": orderID}).
		PlaceholderFormat(squirrel.Dollar)

	sql, args, err := selectBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query: %v", err)
	}

	rows, err := tx.Query(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed to execute SQL query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var sku uint32
		var count uint16
		err := rows.Scan(&sku, &count)
		if err != nil {
			return fmt.Errorf("failed to scan row: %v", err)
		}

		updateBuilder := squirrel.Update("Stocks").
			Set("count", squirrel.Expr("count + ?", count)).
			Where(squirrel.Eq{"sku": sku}).
			PlaceholderFormat(squirrel.Dollar)

		sql, args, err := updateBuilder.ToSql()
		if err != nil {
			return fmt.Errorf("failed to build SQL query: %v", err)
		}

		_, err = tx.Exec(ctx, sql, args...)
		if err != nil {
			return fmt.Errorf("failed to execute SQL query: %v", err)
		}
	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("error iterating over rows: %v", err)
	}

	return nil
}

func (r *Repository) getListStatusAndUser(ctx context.Context, tx pgx.Tx, orderID int64) (string, int64, error) {
	query := "SELECT status, user_id FROM Orders WHERE orderID = $1"

	var status string
	var userID int64

	err := tx.QueryRow(ctx, query, orderID).Scan(&status, &userID)
	if err != nil {
		return "", 0, fmt.Errorf("failed to exec SQL query listorder: %v", err)
	}

	return status, userID, nil
}

func (r *Repository) getListCartItems(ctx context.Context, tx pgx.Tx, orderID int64) ([]*loms.Items, error) {
	var items []*loms.Items

	queryItems, args, err := squirrel.
		Select("*").
		From("OrderItems").
		Where(squirrel.Eq{"orderID": orderID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		panic(err)
	}

	rows, err := tx.Query(ctx, queryItems, args...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var orderID int64
		var sku int64
		var count int32

		err := rows.Scan(&orderID, &sku, &count)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rows: %v", err)
		}

		item := &loms.Items{
			Sku:   uint32(sku),
			Count: uint32(count),
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %v", err)
	}
	return items, nil
}
