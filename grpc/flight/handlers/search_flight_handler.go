package handlers

import (
	"context"
	"go-training/grpc/flight/requests"
	"go-training/pb"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *FlightHandler) SearchFlight(ctx context.Context, m *pb.SearchFlightRequest) (*pb.SearchFlightResponse, error) {
	req := &requests.SearchFlightRequest{
		From: m.From,
		To:   m.From,
		Date: time.Date(int(m.DepartDate.Year), time.Month(m.DepartDate.Month), int(m.DepartDate.Day), 0, 0, 0, 0, time.UTC),
	}

	flights, err := h.flightRepository.SearchFlight(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err != nil {
		return nil, err
	}

	res := &pb.SearchFlightResponse{
		Flights: []*pb.Flight{},
	}

	for _, flight := range flights {
		res.Flights = append(res.Flights, flight.ToPBModel())
	}

	return res, nil
}
