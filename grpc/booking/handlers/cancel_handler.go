package handlers

import (
	"context"
	"go-training/pb"

	"github.com/google/uuid"
)

func (h *BookingHandler) CancelBooking(ctx context.Context, in *pb.CancelBookingRequest) (*pb.Empty, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, err
	}
	err = h.bookingRepository.CancelBooking(ctx, id)

	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
