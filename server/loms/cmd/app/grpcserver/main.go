package main

import (
	"context"
	"log"
	"net"
	"route256/loms/internal/mw/panic"
	"route256/loms/internal/protoc/loms"
	"route256/loms/internal/service"
	"route256/loms/repository"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = ":50052"

func main() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("can't create listener: %s", err)
	}
	repository.Random_BD()

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(panic.Interceptor))
	reflection.Register(s)
	pool, err := pgxpool.Connect(context.Background(), "postgresql://postgres_u:123@localhost:5434/loms")
	if err != nil {
		log.Fatalf("не удалось подключиться к базе данных: %s", err)
	}
	loms.RegisterLomsServer(s, service.NewServer(repository.NewRepository(pool)))
	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %w", err)
	}

}
