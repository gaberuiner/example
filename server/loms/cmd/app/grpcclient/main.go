package main

import (
	"context"
	"log"
	"route256/loms/internal/methods"
	"route256/loms/internal/protoc/loms"

	"google.golang.org/grpc"
)

const grpcPort = ":50052"

func main() {

	conn, err := grpc.Dial(grpcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := loms.NewLomsClient(conn)

	ctx := context.Background()

	method := methods.New()
	createOrderResp, err := method.CreateOrder(client, ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Create Order response: %v", createOrderResp)
	listOrderResp, err := method.ListOrder(client, ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("List Order response: %v", listOrderResp)
	cancelOrderResp, err := method.CancelOrder(client, ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Cancel Order response: %v", cancelOrderResp)
	orderPayedResp, err := method.OrderPayed(client, ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Order Payed response: %v", orderPayedResp)
	stocksResp, err := method.Stocks(client, ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Stocks response: %v", stocksResp)
}
