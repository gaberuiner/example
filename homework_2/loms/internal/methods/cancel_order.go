package methods

import (
	"context"
	"route256/loms/internal/protoc/loms"
)

func (m Method) CancelOrder(client loms.LomsClient, ctx context.Context) (*loms.CancelOrderResponse, error) {
	cancelOrderRequest := &loms.CancelOrderRequest{OrderId: 645}
	cancelOrderResponse, _ := client.CancelOrder(ctx, cancelOrderRequest)
	// if err != nil {

	// 	return &loms.CancelOrderResponse{}, fmt.Errorf("CancelOrder: %v", err)

	// }
	return cancelOrderResponse, nil
}
