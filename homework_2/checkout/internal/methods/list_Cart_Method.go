package methods

import (
	"context"
	check "route256/checkout/internal/protoc/checkout"
)

func (m Method) ListCartMethod(client check.CheckoutClient, user uint64) (check.ListCartResponse, error) {
	listCartReq := &check.ListCartRequest{
		User: user,
	}
	listCartRes, err := client.ListCart(context.Background(), listCartReq)
	if err != nil {
		return check.ListCartResponse{}, err
	}
	return check.ListCartResponse{Items: listCartRes.Items, TotalPrice: listCartRes.TotalPrice}, nil
}
