package models

import (
	"go-training/pb"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Flight struct {
	Id            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey`
	From          string    `gorm:"column:flight_from"`
	To            string    `gorm:"column:flight_to"`
	DepartDate    time.Time
	DepartTime    time.Time
	Status        string `gorm:"column:flight_status"`
	Slot          string
	FlightPlane   string
	AvailableSlot int
}

func (m *Flight) ToPBModel() *pb.Flight {
	return &pb.Flight{
		Id:   m.Id.String(),
		From: m.From,
		To:   m.To,
		DepartDate: &pb.Date{
			Year:  int32(m.DepartDate.UTC().Year()),
			Month: int32(m.DepartDate.UTC().Month()),
			Day:   int32(m.DepartDate.UTC().Day()),
		},
		DepartTime:    timestamppb.New(m.DepartTime),
		Slot:          m.Slot,
		FlightPlane:   m.FlightPlane,
		AvailableSlot: int32(m.AvailableSlot),
	}
}
