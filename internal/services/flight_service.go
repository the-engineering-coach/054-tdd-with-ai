package services

import (
	"context"

	"the-engineering-coach/tdd-with-ai/internal/domain"
	"the-engineering-coach/tdd-with-ai/internal/ports"
)

type FlightRepository interface {
	FindByOrigin(ctx context.Context, origin string) ([]ports.Flight, error)
}

type FlightService struct {
	repo FlightRepository
}

func NewFlightService(repo FlightRepository) *FlightService {
	return &FlightService{repo: repo}
}

func (s *FlightService) SearchByOrigin(ctx context.Context, origin string) ([]ports.Flight, error) {
	if err := domain.ValidateAirportCode(origin); err != nil {
		return nil, err
	}

	return s.repo.FindByOrigin(ctx, origin)
}
