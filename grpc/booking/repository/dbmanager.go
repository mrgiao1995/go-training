package repository

import (
	"go-training/database"
	"go-training/grpc/booking/models"

	"gorm.io/gorm"
)

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (BookingRepository, error) {
	db, err := database.OpenDBConnection("booking")

	if err != nil {
		return nil, err
	}

	db = db.Debug()

	err = db.AutoMigrate(
		&models.Booking{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db}, nil
}
