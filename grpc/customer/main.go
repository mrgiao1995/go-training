package main

import (
	"fmt"
	"go-training/grpc/customer/handlers"
	"go-training/grpc/customer/repository"
	"go-training/pb"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Starting customer service...")
	listen, err := net.Listen("tcp", ":3001")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer()),
	)

	customerRepository, err := repository.NewDBManager()

	if err != nil {
		panic(err)
	}

	handler, err := handlers.NewCustomerHandler(customerRepository)

	if err != nil {
		panic(err)
	}

	reflection.Register(server)
	pb.RegisterMyCustomerServer(server, handler)

	server.Serve(listen)
}
