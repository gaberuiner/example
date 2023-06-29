package client

import (
	"context"
	"route256/loms/internal/protoc/loms"

	"google.golang.org/grpc"
)

type Client struct {
	LomsClient
}

type LomsClient interface {
	CreateOrder(ctx context.Context, in *loms.CreateOrderRequest, opts ...grpc.CallOption) (*loms.CreateOrderResponse, error)
	ListOrder(ctx context.Context, in *loms.ListOrderRequest, opts ...grpc.CallOption) (*loms.ListOrderResponse, error)
	CancelOrder(ctx context.Context, in *loms.CancelOrderRequest, opts ...grpc.CallOption) (*loms.CancelOrderResponse, error)
	OrderPayed(ctx context.Context, in *loms.OrderPayedRequest, opts ...grpc.CallOption) (*loms.OrderPayedResponse, error)
	Stocks(ctx context.Context, in *loms.StocksRequest, opts ...grpc.CallOption) (*loms.StocksResponse, error)
}

func (c *Client) CreateOrder(ctx context.Context, in *loms.CreateOrderRequest, opts ...grpc.CallOption) (*loms.CreateOrderResponse, error) {
	return &loms.CreateOrderResponse{}, nil

}
func (c *Client) ListOrder(ctx context.Context, in *loms.ListOrderRequest, opts ...grpc.CallOption) (*loms.ListOrderResponse, error) {
	return &loms.ListOrderResponse{}, nil

}
func (c *Client) CancelOrder(ctx context.Context, in *loms.CancelOrderRequest, opts ...grpc.CallOption) (*loms.CancelOrderResponse, error) {
	return &loms.CancelOrderResponse{}, nil

}
func (c *Client) OrderPayed(ctx context.Context, in *loms.OrderPayedRequest, opts ...grpc.CallOption) (*loms.OrderPayedResponse, error) {
	return &loms.OrderPayedResponse{}, nil

}
func (c *Client) Stocks(ctx context.Context, in *loms.StocksRequest, opts ...grpc.CallOption) (*loms.StocksResponse, error) {
	return &loms.StocksResponse{}, nil
}
