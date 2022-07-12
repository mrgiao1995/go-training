package handlers

import (
	"context"
	"go-training/pb"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *FlightHandler) FlightDetails(ctx context.Context, m *pb.FindFlightRequest) (*pb.Flight, error) {
	id, err := uuid.Parse(m.Id)

	if err != nil {
		return nil, err
	}

	c, err := h.flightRepository.FlightDetails(ctx, id)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err != nil {
		return nil, err
	}

	return c.ToPBModel(), nil
}
