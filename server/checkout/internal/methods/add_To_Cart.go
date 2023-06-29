package methods

import (
	"context"
	"fmt"
	loms_stocks "route256/checkout/internal/other_server_connection/lomsstocks"
	check "route256/checkout/internal/protoc/checkout"
)

func (m Method) AddToCartMethod(client check.CheckoutClient, user uint64, sku uint32, count uint32) (check.AddToCartResponse, error) {
	addToCartReq := &check.AddToCartRequest{
		User:  user,
		Sku:   sku,
		Count: count,
	}

	stocks, err := loms_stocks.LomsStocks(addToCartReq.Sku)
	if len(stocks.Stocks) == 0 || err != nil {
		return check.AddToCartResponse{}, fmt.Errorf("Error checking items in stock failed %v", err)
	}
	_, err = client.AddToCart(context.Background(), addToCartReq)
	if err != nil {
		return check.AddToCartResponse{}, err
	}

	return check.AddToCartResponse{}, nil
}
