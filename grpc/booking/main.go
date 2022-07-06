package main

import (
	"fmt"
	"go-training/grpc/booking/handlers"
	"go-training/grpc/booking/repository"
	"go-training/pb"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Starting booking service...")
	listen, err := net.Listen("tcp", ":3002")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer()),
	)

	bookingRepository, err := repository.NewDBManager()

	if err != nil {
		panic(err)
	}

	handler, err := handlers.NewBookingHandler(bookingRepository)

	if err != nil {
		panic(err)
	}

	reflection.Register(server)
	pb.RegisterMyBookingServer(server, handler)

	server.Serve(listen)
}
