package ports

import (
	"context"
	"time"
)

type Flight struct {
	FlightNumber       string    `json:"flight_number"`
	OriginAirport      string    `json:"origin_airport"`
	DestinationAirport string    `json:"destination_airport"`
	DepartureTime      time.Time `json:"departure_time"`
	Duration           int       `json:"duration"`
	Airline            string    `json:"airline"`
}

type FlightService interface {
	SearchByOrigin(ctx context.Context, origin string) ([]Flight, error)
}
