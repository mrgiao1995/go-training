package handlers

import (
	"context"
	"go-training/grpc/flight/models"
	"go-training/pb"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *FlightHandler) UpdateFlight(ctx context.Context, m *pb.Flight) (*pb.Flight, error) {
	id, err := uuid.Parse(m.Id)

	if err != nil {
		panic(err)
	}
	req := &models.Flight{
		Id:            id,
		From:          m.From,
		To:            m.From,
		DepartDate:    time.Date(int(m.DepartDate.Year), time.Month(m.DepartDate.Month), int(m.DepartDate.Day), 0, 0, 0, 0, time.UTC),
		DepartTime:    m.DepartTime.AsTime(),
		Status:        m.Status,
		Slot:          m.Slot,
		FlightPlane:   m.FlightPlane,
		AvailableSlot: int(m.AvailableSlot),
	}

	flight, err := h.flightRepository.UpdateFlight(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err != nil {
		return nil, err
	}

	return flight.ToPBModel(), nil
}
