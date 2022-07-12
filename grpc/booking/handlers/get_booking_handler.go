package handlers

import (
	"context"
	"go-training/pb"

	"github.com/google/uuid"
)

func (h *BookingHandler) GetBookings(ctx context.Context, in *pb.GetBookingsRequest) (*pb.Bookings, error) {
	cid, err := uuid.Parse(in.CustomerId)
	if err != nil {
		return nil, err
	}

	bs, err := h.bookingRepository.GetBookings(ctx, cid)

	if err != nil {
		return nil, err
	}

	res := &pb.Bookings{
		Bookings: []*pb.Booking{},
	}

	for _, b := range bs {
		res.Bookings = append(res.Bookings, b.ToPBModel())
	}

	return res, nil
}
