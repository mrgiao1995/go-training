package repository

import (
	"context"
	"go-training/grpc/booking/models"

	"github.com/google/uuid"
)

type BookingRepository interface {
	CreateBooking(context context.Context, model *models.Booking) (*models.Booking, error)
	ViewBooking(context context.Context, id uuid.UUID, bookingCode string) (*models.Booking, error)
	CancelBooking(context context.Context, id uuid.UUID) error
	GetBookings(context context.Context, cid uuid.UUID) ([]*models.Booking, error)
}

func (conn *dbmanager) CreateBooking(context context.Context, model *models.Booking) (*models.Booking, error) {
	model.Id = uuid.New()
	if err := conn.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (conn *dbmanager) GetBookings(context context.Context, cid uuid.UUID) ([]*models.Booking, error) {
	b := []*models.Booking{}
	err := conn.Where(&models.Booking{CustomerId: cid}).Find(&b).Error

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (conn *dbmanager) ViewBooking(context context.Context, cid uuid.UUID, bookingCode string) (*models.Booking, error) {
	b := &models.Booking{}
	err := conn.First(&models.Booking{CustomerId: cid, BookingCode: bookingCode}).Find(&b).Error

	if err != nil {
		return nil, err
	}

	return b, nil
}
func (conn *dbmanager) CancelBooking(context context.Context, id uuid.UUID) error {
	err := conn.Where(&models.Booking{Id: id}).Select("status").Updates(&models.Booking{Id: id, Status: "Canceled"}).Error

	if err != nil {
		return err
	}
	return nil
}
