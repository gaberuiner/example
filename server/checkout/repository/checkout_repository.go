package repository

import (
	"context"
	"fmt"
	"log"
	productservice "route256/checkout/internal/other_server_connection/product_service"
	"route256/checkout/internal/protoc/checkout"

	
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) AddToCart(ctx context.Context, req *checkout.AddToCartRequest) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction ListOrder: %v", err)
	}
	defer tx.Rollback(ctx)

	query := "INSERT INTO CartItem (user_id, sku, count) VALUES ($1, $2, $3)"
	_, err = tx.Exec(ctx, query, int64(req.User), int32(req.Sku), int32(req.Count))
	if err != nil {
		return fmt.Errorf("failed to insert into CartItem: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	return nil

}

func (r *Repository) DeleteFromCart(ctx context.Context, req *checkout.DeleteFromCartRequest) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction ListOrder: %v", err)
	}
	defer tx.Rollback(ctx)

	query := "DELETE FROM CartItem WHERE user_id = $1"

	_, err = tx.Exec(ctx, query, int64(req.User))
	if err != nil {
		return fmt.Errorf("failed to insert into CartItem: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	return nil
}

func (r *Repository) ListCart(ctx context.Context, req *checkout.ListCartRequest) (*checkout.ListCartResponse, error) {

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction ListCart: %v", err)
	}
	defer tx.Rollback(ctx)

	query := "SELECT user_id, sku, count FROM CartItem WHERE user_id = $1"
	rows, err := tx.Query(ctx, query, req.User)
	if err != nil {
		return nil, fmt.Errorf("failed to query CartItem: %v", err)
	}
	defer rows.Close()

	response := &checkout.ListCartResponse{}
	var TotalPrice int
	for rows.Next() {
		var user_id int64
		var sku int
		var count int16

		err := rows.Scan(&user_id, &sku, &count)
		if err != nil {
			return nil, fmt.Errorf("failed to scan CartItem row: %v", err)
		}

		// Создание объекта CartItem и добавление его в response
		item := &checkout.Items{
			Sku:   uint32(sku),
			Count: uint32(count),
		}
		itemsNameAndPrice, err := productservice.GetProductServer("Nagnm1rZz685OJgCYYHFIQyz", uint32(sku))
		if err != nil {
			log.Println("failed to get items name and price by sku: %w", sku)
		}
		item.Name = itemsNameAndPrice.Name
		item.Price = itemsNameAndPrice.Price
		TotalPrice += int(item.Price) * int(item.Count)
		response.Items = append(response.Items, item)
	}
	response.TotalPrice = uint32(TotalPrice)

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during CartItem query iteration: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return response, nil
}

func (r *Repository) Purchase(ctx context.Context, req *checkout.PurchaseRequest) (*checkout.PurchaseResponse, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction Purchase: %v", err)
	}
	defer tx.Rollback(ctx)

	query := "SELECT orderID FROM Checkout WHERE user_id = $1"
	var orderID int64
	err = tx.QueryRow(ctx, query, req.User).Scan(&orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to query Checkout: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return &checkout.PurchaseResponse{
		OrderID: orderID,
	}, nil
}
