package repository

import (
	"context"
	"errors"
	"go-training/grpc/flight/models"
	"go-training/grpc/flight/requests"

	"github.com/google/uuid"
)

type FlightRepository interface {
	CreateFlight(context context.Context, model *models.Flight) (*models.Flight, error)
	UpdateFlight(context context.Context, model *models.Flight) (*models.Flight, error)
	SearchFlight(context context.Context, request *requests.SearchFlightRequest) ([]*models.Flight, error)
	FlightDetails(context context.Context, id uuid.UUID) (*models.Flight, error)
}

func (conn *dbmanager) CreateFlight(context context.Context, model *models.Flight) (*models.Flight, error) {
	model.Id = uuid.New()
	if err := conn.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (conn *dbmanager) UpdateFlight(context context.Context, model *models.Flight) (*models.Flight, error) {
	flights := []*models.Flight{}
	err := conn.Where(&models.Flight{Id: model.Id}).Find(&flights).Updates(model).Error
	if err != nil {
		return nil, err
	}

	if len(flights) == 0 {
		return nil, errors.New("RECORD NOT FOUND")
	}

	return model, nil
}

func (conn *dbmanager) SearchFlight(context context.Context, request *requests.SearchFlightRequest) ([]*models.Flight, error) {
	flights := []*models.Flight{}

	if err := conn.Where("flight_from = ? or flight_to = ? or depart_date = ?",
		request.From, request.To, request.Date).Find(&flights).Error; err != nil {
		return nil, err
	}

	return flights, nil
}

func (conn *dbmanager) FlightDetails(context context.Context, id uuid.UUID) (*models.Flight, error) {
	cs := &models.Flight{}
	err := conn.First(&models.Flight{Id: id}).Find(&cs).Error

	if err != nil {
		return nil, err
	}

	return cs, nil
}
