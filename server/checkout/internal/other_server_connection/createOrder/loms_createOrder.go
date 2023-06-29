package loms_createOrder

import (
	"context"
	"errors"
	"log"
	loms "route256/checkout/internal/loms_client"

	"google.golang.org/grpc"
)

const grpcPort = ":50052"

var ErrorOrderId = errors.New("Create order failed")

func LomsCreateOrder(user uint64, items []*loms.Items) error {
	conn, err := grpc.Dial(grpcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := loms.NewLomsClient(conn)
	request := loms.CreateOrderRequest{
		User: user,
		Item: items,
	}
	getCreateOrderResponse, err := client.CreateOrder(context.Background(), &request)
	if err != nil {
		return err
	}
	switch {
	case getCreateOrderResponse.OrderID > 0:
		return nil
	default:
		return ErrorOrderId
	}

}
