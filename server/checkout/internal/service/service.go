package service

import (
	"context"
	"route256/checkout/internal/protoc/checkout"
	"route256/checkout/repository"

	"google.golang.org/grpc"
)

type Service struct {
	checkout.UnimplementedCheckoutServer
	CheckoutClient
	TransactionManager
	repository *repository.Repository
}

type TransactionManager interface {
	// Определите методы управления транзакциями, если требуется
}

type CheckoutClient interface {
	AddToCart(ctx context.Context, in *checkout.AddToCartRequest, opts ...grpc.CallOption) (*checkout.AddToCartResponse, error)
	DeleteFromCart(ctx context.Context, in *checkout.DeleteFromCartRequest, opts ...grpc.CallOption) (*checkout.DeleteFromCartResponse, error)
	ListCart(ctx context.Context, in *checkout.ListCartRequest, opts ...grpc.CallOption) (*checkout.ListCartResponse, error)
	Purchase(ctx context.Context, in *checkout.PurchaseRequest, opts ...grpc.CallOption) (*checkout.PurchaseResponse, error)
}

func (s *Service) AddToCart(ctx context.Context, req *checkout.AddToCartRequest) (*checkout.AddToCartResponse, error) {
	err := s.repository.AddToCart(ctx, req)
	if err != nil {
		return nil, err
	}
	return &checkout.AddToCartResponse{}, nil
}

func (s *Service) DeleteFromCart(ctx context.Context, req *checkout.DeleteFromCartRequest) (*checkout.DeleteFromCartResponse, error) {
	err := s.repository.DeleteFromCart(ctx, req)
	if err != nil {
		return nil, err
	}
	return &checkout.DeleteFromCartResponse{}, nil
}

func (s *Service) ListCart(ctx context.Context, req *checkout.ListCartRequest) (*checkout.ListCartResponse, error) {
	items, err := s.repository.ListCart(ctx, req)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Service) Purchase(ctx context.Context, req *checkout.PurchaseRequest) (*checkout.PurchaseResponse, error) {
	resp, err := s.repository.Purchase(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Service) mustEmbedUnimplementedCheckoutServer() {}

func NewServer(repository *repository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}
