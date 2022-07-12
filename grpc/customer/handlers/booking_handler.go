package handlers

import (
	"context"
	"go-training/pb"
)

func (h *CustomerHandler) ViewCustomerBookingHistories(ctx context.Context, m *pb.ViewCustomerBookingHistoriesRequest) (*pb.ViewCustomerBookingHistoriesResponse, error) {
	bs, err := getBookings(m.Id)

	if err != nil {
		return nil, err
	}

	res := &pb.ViewCustomerBookingHistoriesResponse{
		Bookings: []*pb.ViewBooking{},
	}

	for _, b := range bs.Bookings {
		res.Bookings = append(res.Bookings, &pb.ViewBooking{
			BookingCode: b.BookingCode,
			BookingDate: b.BookingDate,
		})
	}
	return res, nil
}
