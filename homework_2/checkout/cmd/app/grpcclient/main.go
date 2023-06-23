package main

import (
	"log"
	"route256/checkout/internal/methods"
	check "route256/checkout/internal/protoc/checkout"

	"google.golang.org/grpc"
)

const grpcPort = ":50051"

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Ошибка при установлении соединения: %v", err)
	}
	defer conn.Close()

	client := check.NewCheckoutClient(conn)

	Method := methods.New()
	_, err = Method.AddToCartMethod(client, 1, 1, 3)
	if err != nil {
		log.Fatalf("failed to request AddToCart handle: %v", err)
	}
	log.Printf("AddToCart response success")
	_, err = Method.DeleteFromCartMethod(client, 1, 2, 3)
	if err != nil {
		log.Fatalf("failed to request DeleteFromCart handle: %v", err)
	}
	log.Printf("DeleteFromCart response success")

	ListCartResp, err := Method.ListCartMethod(client, 3)
	if err != nil {
		log.Fatalf("failed to request ListCart handle: %v", err)
	}
	log.Printf("List Cart Response: items: %v, totalPrice %v", ListCartResp.Items, ListCartResp.TotalPrice)
	log.Printf("ListCart response success")

	PurchaseResp, err := Method.PurchaseMethod(client, 3)
	if err != nil {
		log.Fatalf("failed to request Purchase handle: %v", err)
	}
	log.Printf("Purchase Response: %v", PurchaseResp.OrderID)

	log.Printf("Purchase response success")
}
