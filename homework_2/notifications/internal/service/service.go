package service

import (
	"context"
	"route256/ProductService/internal/protoc/ProductService"
)

type Server struct {
	ProductService.UnimplementedProductServiceServer
}

// GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error)
//
//	ListSkus(context.Context, *ListSkusRequest) (*ListSkusResponse, error)
func (s *Server) GetProduct(context.Context, *ProductService.GetProductRequest) (*ProductService.GetProductResponse, error) {
	return &ProductService.GetProductResponse{
		Name:  "Name",
		Price: 1,
	}, nil
}
func (s *Server) ListSkus(context.Context, *ProductService.ListSkusRequest) (*ProductService.ListSkusResponse, error) {
	return &ProductService.ListSkusResponse{
		Skus: []uint32{1, 2, 3},
	}, nil

}

func NewServer() *Server {
	return &Server{}
}
