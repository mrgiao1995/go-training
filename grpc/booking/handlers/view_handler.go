package handlers

import (
	"context"
	"go-training/pb"

	"github.com/google/uuid"
)

func (h *BookingHandler) ViewBooking(ctx context.Context, in *pb.ViewBookingRequest) (*pb.ViewBookingResponse, error) {
	cid, err := uuid.Parse(in.CustomerId)
	if err != nil {
		return nil, err
	}
	b, err := h.bookingRepository.ViewBooking(ctx, cid, in.BookingCode)
	if err != nil {
		return nil, err
	}
	c, err := getCustomerInfomation(in.CustomerId)
	if err != nil {
		return nil, err
	}
	f, err := getFlightInfomation(b.FlightId.String())
	if err != nil {
		return nil, err
	}

	return &pb.ViewBookingResponse{
		BookingCode: in.BookingCode,
		BookingDate: &pb.Date{
			Year:  int32(b.BookingDate.UTC().Year()),
			Month: int32(b.BookingDate.UTC().Month()),
			Day:   int32(b.BookingDate.UTC().Day()),
		},
		Customer: c,
		Flight:   f,
	}, nil
}
