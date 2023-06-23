package methods

import (
	"context"
	"fmt"
	"route256/loms/internal/protoc/loms"
)

func (m Method) ListOrder(client loms.LomsClient, ctx context.Context) (*loms.ListOrderResponse, error) {
	listOrderRequest := &loms.ListOrderRequest{OrderID: 645}
	listOrderResponse, err := client.ListOrder(ctx, listOrderRequest)
	if err != nil {
		return &loms.ListOrderResponse{}, fmt.Errorf("ListOrder failed: %v", err)

	}
	return listOrderResponse, nil
}
