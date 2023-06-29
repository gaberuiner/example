package methods

import (
	"context"
	check "route256/checkout/internal/protoc/checkout"
)

func (m Method) DeleteFromCartMethod(client check.CheckoutClient, user uint64, sku uint32, count uint32) (check.DeleteFromCartResponse, error) {
	deleteFromCartReq := &check.DeleteFromCartRequest{
		User:  user,
		Sku:   sku,
		Count: count,
	}
	_, err := client.DeleteFromCart(context.Background(), deleteFromCartReq)
	if err != nil {
		return check.DeleteFromCartResponse{}, err
	}
	return check.DeleteFromCartResponse{}, nil
}
