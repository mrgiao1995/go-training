package handlers

import (
	"context"
	"go-training/grpc/booking/repository"
	"go-training/pb"
	"time"

	"google.golang.org/grpc"
)

type BookingHandler struct {
	pb.UnimplementedMyBookingServer
	bookingRepository repository.BookingRepository
}

func NewBookingHandler(bookingRepository repository.BookingRepository) (*BookingHandler, error) {
	return &BookingHandler{bookingRepository: bookingRepository}, nil
}

func getCustomerInfomation(id string) (*pb.Customer, error) {
	customerConn, err := grpc.Dial(":3001", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	customerClient := pb.NewMyCustomerClient(customerConn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	customer, err := customerClient.CustomerDetails(ctx, &pb.FindCustomerRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func getFlightInfomation(id string) (*pb.Flight, error) {
	flightConn, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	filghtClient := pb.NewMyFlightClient(flightConn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	f, err := filghtClient.FlightDetails(ctx, &pb.FindFlightRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return f, nil
}

func updateFlightAvailableSlot(flight *pb.Flight) error {
	flightConn, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		return err
	}

	filghtClient := pb.NewMyFlightClient(flightConn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = filghtClient.UpdateFlight(ctx, flight)
	if err != nil {
		return err
	}

	return nil

}
