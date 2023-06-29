package methods

import (
	"context"
	"fmt"
	"route256/loms/internal/protoc/loms"
)

func (m Method) Stocks(client loms.LomsClient, ctx context.Context) (*loms.StocksResponse, error) {
	stocksRequest := &loms.StocksRequest{Sku: 1}
	stocksResponse, err := client.Stocks(ctx, stocksRequest)
	if err != nil {
		return &loms.StocksResponse{}, fmt.Errorf("Stocks failed: %v", err)

	}
	return stocksResponse, nil

}
