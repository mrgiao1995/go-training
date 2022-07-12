package responses

import "go-training/grpc/flight/models"

type SearchFlightResponse struct {
	flights []models.Flight
}
