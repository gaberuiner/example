package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func BD_Data() {
	db, err := sql.Open("postgres", "host=localhost port=5433 user=postgres password=123 dbname=checkout sslmode=disable")
	if err != nil {
		fmt.Println("failed to connect to bd: ", err)
		return
	}

	queryPrepareAddToCart := "INSERT INTO Checkout (user_id, totalPrice, orderID) VALUES ($1, $2, $3)"
	_, err = db.Exec(queryPrepareAddToCart, 1, 10, 2)
	if err != nil {
		fmt.Println("ERR STOCS CREATING: ", err)
	}
	_, err = db.Exec(queryPrepareAddToCart, 3, 11, 3)
	if err != nil {
		fmt.Println("ERR ListCart CREATING: ", err)
	}
	queryPrepareListCart := "INSERT INTO CartItem (user_id, sku, count) VALUES ($1, $2, $3)"
	_, err = db.Exec(queryPrepareListCart, 3, 773297411, 2)
	if err != nil {
		fmt.Println("ERR ListCart CREATING: ", err)
	}

}
