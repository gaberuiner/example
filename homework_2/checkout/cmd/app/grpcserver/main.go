package main

import (
	"context"
	"log"
	"net"
	"route256/checkout/internal/mw/panic"
	"route256/checkout/internal/protoc/checkout"
	"route256/checkout/internal/service"
	"route256/checkout/repository"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/jackc/pgx/v4/pgxpool"

	"net/http"
	_ "net/http/pprof"
)

const grpcPort = ":50051"

func main() {

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("can't create listener: %s", err)
	}
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	repository.BD_Data()
	pool, err := pgxpool.Connect(context.Background(), "postgresql://postgres:123@localhost:5433/checkout")
	if err != nil {
		log.Fatalf("не удалось подключиться к базе данных: %s", err)
	}
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(panic.Interceptor))
	reflection.Register(s)
	checkout.RegisterCheckoutServer(s, service.NewServer(repository.New(pool)))
	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %w", err)
	}

}
