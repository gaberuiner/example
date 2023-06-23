package loms_stocks

import (
	"context"
	"fmt"
	"log"
	loms "route256/checkout/internal/loms_client"

	"google.golang.org/grpc"
)

const grpcPort = ":50052"

func LomsStocks(sku uint32) (loms.StocksResponse, error) {

	conn, err := grpc.Dial(grpcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := loms.NewLomsClient(conn)
	request := loms.StocksRequest{
		Sku: sku,
	}

	getProductResponse, err := client.Stocks(context.Background(), &request)
	if err != nil {
		fmt.Println("ERROR HERE")
		return loms.StocksResponse{}, err
	}

	return loms.StocksResponse{Stocks: getProductResponse.Stocks}, nil

}
