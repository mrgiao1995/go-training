package models

import (
	"go-training/pb"
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	Id          uuid.UUID
	BookingCode string
	BookingDate time.Time
	Status      string
	CustomerId  uuid.UUID
	FlightId    uuid.UUID
}

func (m *Booking) ToPBModel() *pb.Booking {
	return &pb.Booking{
		Id:          m.Id.String(),
		BookingCode: m.BookingCode,
		BookingDate: &pb.Date{
			Year:  int32(m.BookingDate.UTC().Year()),
			Month: int32(m.BookingDate.UTC().Month()),
			Day:   int32(m.BookingDate.UTC().Day()),
		},
		Status:     m.Status,
		CustomerId: m.CustomerId.String(),
		FlightId:   m.FlightId.String(),
	}
}
