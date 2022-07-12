package handlers

import (
	"context"
	"errors"
	"go-training/grpc/booking/models"
	"go-training/pb"
	"time"

	"github.com/google/uuid"
)

func (h *BookingHandler) CreateBooking(ctx context.Context, m *pb.Booking) (*pb.Booking, error) {
	flight, err := getFlightInfomation(m.FlightId)

	if err != nil {
		return nil, err
	}

	bookingDate := time.Date(int(m.BookingDate.Year), time.Month(m.BookingDate.Month), int(m.BookingDate.Day), 0, 0, 0, 0, time.UTC)
	flightDate := time.Date(int(flight.DepartDate.Year), time.Month(flight.DepartDate.Month), int(flight.DepartDate.Day), flight.DepartTime.AsTime().Hour(), flight.DepartTime.AsTime().Minute(), 0, 0, time.UTC)

	if bookingDate.Sub(flightDate).Hours() >= 12 {
		return nil, errors.New("MUST BE BOOKING 12HOURS EARLY")
	}

	if flight.AvailableSlot <= 0 {
		return nil, errors.New("NO AVAILABLE SLOT")
	}

	_, err = getCustomerInfomation(m.CustomerId)

	if err != nil {
		return nil, err
	}

	cid, err := uuid.Parse(m.CustomerId)

	if err != nil {
		return nil, err
	}

	fid, err := uuid.Parse(m.FlightId)

	if err != nil {
		return nil, err
	}
	b, err := h.bookingRepository.CreateBooking(ctx, &models.Booking{
		BookingDate: bookingDate,
		Status:      "Created",
		CustomerId:  cid,
		FlightId:    fid,
		BookingCode: m.BookingCode,
	})

	if err != nil {
		return nil, err
	}

	flight.AvailableSlot = flight.AvailableSlot - 1
	err = updateFlightAvailableSlot(flight)
	if err != nil {
		return nil, err
	}

	return b.ToPBModel(), nil
}
