package repository

import (
	"database/sql"
	"fmt"
	"route256/loms/internal/protoc/loms"

	_ "github.com/lib/pq"
)

func Random_BD() {
	db, err := sql.Open("postgres", "host=localhost port=5434 user=postgres_u password=123 dbname=loms sslmode=disable")
	if err != nil {
		fmt.Println("failed to connect to bd: ", err)
		return
	}

	queryStocks := "INSERT INTO Stocks (warehouseID, sku, count) VALUES ($1, $2, $3)"
	_, err = db.Exec(queryStocks, 12, 1, 2)
	if err != nil {
		fmt.Println("ERR STOCS CREATING: ", err)
	}
	queryListOrder := "INSERT INTO Orders (orderID, status, user_id) VALUES ($1, $2, $3)"
	_, err = db.Exec(queryListOrder, 645, "new", 5)
	if err != nil {
		fmt.Println("ERR STOCS CREATING: ", err)
	}
	createOrderRequest := &loms.CreateOrderRequest{User: 1, Item: []*loms.Items{{Sku: 1, Count: 2}}}
	orderID := 1
	_, err = db.Exec("INSERT INTO Orders (orderID, status, user_id) VALUES ($1, $2, $3)",
		orderID, "pending", createOrderRequest.User)
	if err != nil {
		fmt.Println("failed to insert data into Orders table: ", err)
	}
	_, err = db.Exec("INSERT INTO OrderItems (orderID, sku, count) VALUES ($1, $2, $3)",
		orderID, 1, 2)
	if err != nil {
		fmt.Println("failed to insert data into OrderItems table: ", err)
	}
	_, err = db.Exec("INSERT INTO stocks (warehouseID, sku, count) VALUES ($1, $2, $3)",
		1, 1, 2)

}
