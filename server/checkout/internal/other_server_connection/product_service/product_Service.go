package productservice

import (
	"context"
	"errors"
	"log"
	productService_v1 "route256/checkout/internal/other_server_connection/product_service/ProductService"
	"route256/checkout/internal/protoc/checkout"

	"google.golang.org/grpc"
)

var serverAddress = "route256.pavl.uk:8082"

func GetProductServer(token string, sku uint32) (productService_v1.GetProductResponse, error) {

	// Connect to the gRPC server
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	client := productService_v1.NewProductServiceClient(conn)
	request := productService_v1.GetProductRequest{
		Token: token,
		Sku:   sku,
	}
	getCreateOrderResponse, err := client.GetProduct(context.Background(), &request)
	if err != nil {
		return productService_v1.GetProductResponse{}, errors.New("ERROR HERE HELLO")
	}
	return productService_v1.GetProductResponse{Name: getCreateOrderResponse.Name, Price: getCreateOrderResponse.Price}, nil
}

func ListSkusServer(token string, startAfterSku uint32, count uint32) (productService_v1.ListSkusResponse, error) {

	// Connect to the gRPC server
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := productService_v1.NewProductServiceClient(conn)
	request := productService_v1.ListSkusRequest{
		Token:         token,
		StartAfterSku: startAfterSku,
		Count:         count,
	}
	getListSkusResponse, err := client.ListSkus(context.Background(), &request)
	if err != nil {
		return productService_v1.ListSkusResponse{}, err
	}
	return productService_v1.ListSkusResponse{Skus: getListSkusResponse.Skus}, nil
}

func GetFromProductService(token string) (checkout.ListCartResponse, error) {
	listskusresp, err := ListSkusServer(token, 0, 100)
	if err != nil {
		return checkout.ListCartResponse{}, err
	}

	var total uint32
	var newitems []*checkout.Items
	for _, sku := range listskusresp.Skus {
		resp, err := GetProductServer(token, sku)
		if err != nil {
			return checkout.ListCartResponse{}, err
		}

		newitem := &checkout.Items{ // Инициализация newitem перед заполнением значений
			Sku:   sku,
			Count: 1,
			Name:  resp.Name,
			Price: resp.Price,
		}

		total += resp.Price
		newitems = append(newitems, newitem)
	}

	return checkout.ListCartResponse{Items: newitems, TotalPrice: total}, nil
}
