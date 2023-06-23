package methods

import (
	"context"
	"fmt"
	"route256/loms/internal/protoc/loms"
)

func (m Method) OrderPayed(client loms.LomsClient, ctx context.Context) (*loms.OrderPayedResponse, error) {
	orderPayedRequest := &loms.OrderPayedRequest{OrderID: 4}
	orderPayedResponse, err := client.OrderPayed(ctx, orderPayedRequest)
	if err != nil {
		return &loms.OrderPayedResponse{}, fmt.Errorf("OrderPayed failed: %v", err)

	}
	return orderPayedResponse, nil
}
