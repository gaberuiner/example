package service

import (
	"context"
	"route256/loms/internal/protoc/loms"
	"route256/loms/repository"
)

type Server struct {
	repository *repository.Repository
	loms.UnimplementedLomsServer
	LomsServer
}

type LomsServer interface {
	CreateOrder(context.Context, *loms.CreateOrderRequest) (*loms.CreateOrderResponse, error)
	ListOrder(context.Context, *loms.ListOrderRequest) (*loms.ListOrderResponse, error)
	CancelOrder(context.Context, *loms.CancelOrderRequest) (*loms.CancelOrderResponse, error)
	OrderPayed(context.Context, *loms.OrderPayedRequest) (*loms.OrderPayedResponse, error)
	Stocks(context.Context, *loms.StocksRequest) (*loms.StocksResponse, error)
	mustEmbedUnimplementedLomsServer()
}

func (s *Server) CreateOrder(ctx context.Context, req *loms.CreateOrderRequest) (*loms.CreateOrderResponse, error) {
	resp, err := s.repository.CreateOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (s *Server) ListOrder(ctx context.Context, req *loms.ListOrderRequest) (*loms.ListOrderResponse, error) {
	resp, err := s.repository.ListOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (s *Server) CancelOrder(ctx context.Context, req *loms.CancelOrderRequest) (*loms.CancelOrderResponse, error) {

	err := s.repository.CancelOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	return &loms.CancelOrderResponse{}, nil
}
func (s *Server) OrderPayed(ctx context.Context, req *loms.OrderPayedRequest) (*loms.OrderPayedResponse, error) {
	err := s.repository.OrderPaid(ctx, req)
	if err != nil {
		return nil, err
	}
	return &loms.OrderPayedResponse{}, nil

}
func (s *Server) Stocks(ctx context.Context, req *loms.StocksRequest) (*loms.StocksResponse, error) {

	resp, err := s.repository.Stocks(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (s *Server) mustEmbedUnimplementedLomsServer() {}

func NewServer(repository *repository.Repository) *Server {
	return &Server{repository: repository}
}
