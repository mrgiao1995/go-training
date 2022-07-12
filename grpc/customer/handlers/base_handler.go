package handlers

import (
	"context"
	"go-training/grpc/customer/repository"
	"go-training/pb"
	"time"

	"google.golang.org/grpc"
)

type CustomerHandler struct {
	pb.UnimplementedMyCustomerServer
	customerRepository repository.CustomerRepository
}

func NewCustomerHandler(customerRepository repository.CustomerRepository) (*CustomerHandler, error) {
	return &CustomerHandler{customerRepository: customerRepository}, nil
}

func getBookings(id string) (*pb.Bookings, error) {
	bookingConn, err := grpc.Dial(":3002", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	bookingClient := pb.NewMyBookingClient(bookingConn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	f, err := bookingClient.GetBookings(ctx, &pb.GetBookingsRequest{CustomerId: id})
	if err != nil {
		return nil, err
	}

	return f, nil
}
