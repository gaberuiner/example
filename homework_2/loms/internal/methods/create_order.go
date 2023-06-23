package methods

import (
	"context"
	"fmt"
	"route256/loms/internal/protoc/loms"
)

func (m Method) CreateOrder(client loms.LomsClient, ctx context.Context) (*loms.CreateOrderResponse, error) {
	createOrderRequest := &loms.CreateOrderRequest{User: 1, Item: []*loms.Items{{Sku: 1, Count: 2}}}
	createOrderResponse, err := client.CreateOrder(ctx, createOrderRequest)
	if err != nil {
		return &loms.CreateOrderResponse{}, fmt.Errorf("CreateOrder failed: %v", err)

	}
	return createOrderResponse, nil
}
