package main

import (
	"fmt"
	"go-training/grpc/flight/handlers"
	"go-training/grpc/flight/repository"
	"go-training/pb"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Starting flight service...")
	listen, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer()),
	)

	flightRepository, err := repository.NewDBManager()

	if err != nil {
		panic(err)
	}

	handler, err := handlers.NewFlightHandler(flightRepository)

	if err != nil {
		panic(err)
	}

	reflection.Register(server)
	pb.RegisterMyFlightServer(server, handler)

	server.Serve(listen)
}
