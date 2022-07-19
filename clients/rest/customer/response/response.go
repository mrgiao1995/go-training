package response

import (
	"time"

	"github.com/gofrs/uuid"
)

type Customer struct {
	Id      uuid.UUID
	Email   string
	Name    string
	Address string
	DoB     time.Time
}

type CustomerBookingHistories struct {
	Bookings []ViewBooking
}
type Flight struct {
	Id            string
	From          string
	To            string
	DepartDate    string
	DepartTime    string
	Status        string
	Slot          string
	FlightPlane   string
	AvailableSlot int
}
type ViewBooking struct {
	BookingCode string
	Flight      Flight
}
