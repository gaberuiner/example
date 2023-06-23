package client

import (
	"context"
	"route256/checkout/internal/protoc/checkout"

	"google.golang.org/grpc"
)

type Client struct {
	checkout.CheckoutClient
}

func (c *Client) AddToCart(ctx context.Context, in *checkout.AddToCartRequest, opts ...grpc.CallOption) (*checkout.AddToCartResponse, error) {
	return &checkout.AddToCartResponse{}, nil

}

func (c *Client) DeleteFromCart(ctx context.Context, in *checkout.DeleteFromCartRequest, opts ...grpc.CallOption) (*checkout.DeleteFromCartResponse, error) {
	return &checkout.DeleteFromCartResponse{}, nil
}

func (c *Client) ListCart(ctx context.Context, in *checkout.ListCartRequest, opts ...grpc.CallOption) (*checkout.ListCartResponse, error) {
	return &checkout.ListCartResponse{}, nil
}

func (c *Client) Purchase(ctx context.Context, in *checkout.PurchaseRequest, opts ...grpc.CallOption) (*checkout.PurchaseResponse, error) {
	return &checkout.PurchaseResponse{}, nil
}

func NewClient(c checkout.CheckoutClient) *Client {
	return &Client{CheckoutClient: c}
}
