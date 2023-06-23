package main

import (
	"context"
	"fmt"
	"log"
	"route256/ProductService/internal/protoc/ProductService"

	"google.golang.org/grpc"
)

const grpcPort = ":50053"

func main() {

	conn, err := grpc.Dial(grpcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := ProductService.NewProductServiceClient(conn)

	getProductRequest := &ProductService.GetProductRequest{}

	getProductResponse, err := client.GetProduct(context.Background(), getProductRequest)
	if err != nil {
		log.Fatalf("Error while calling GetProduct RPC: %v", err)
	}

	fmt.Println("GetProduct Response:", getProductResponse)

	listSkusRequest := &ProductService.ListSkusRequest{}

	listSkusResponse, err := client.ListSkus(context.Background(), listSkusRequest)
	if err != nil {
		log.Fatalf("Error while calling ListSkus RPC: %v", err)
	}

	fmt.Println("ListSkus Response:", listSkusResponse)
}
