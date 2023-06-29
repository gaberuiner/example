package client

import (
	"context"
	"route256/ProductService/internal/protoc/ProductService"

	"google.golang.org/grpc"
)

type Client struct {
	ProductService.ProductServiceClient
}

func (c *Client) GetProduct(ctx context.Context, in *ProductService.GetProductRequest, opts ...grpc.CallOption) (*ProductService.GetProductResponse, error) {
	return &ProductService.GetProductResponse{}, nil

}

func (c *Client) ListSkus(ctx context.Context, in *ProductService.ListSkusRequest, opts ...grpc.CallOption) (*ProductService.ListSkusResponse, error) {
	return &ProductService.ListSkusResponse{}, nil
}

// GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error)
// 	ListSkus(ctx context.Context, in *ListSkusRequest, opts ...grpc.CallOption) (*ListSkusResponse, error)
