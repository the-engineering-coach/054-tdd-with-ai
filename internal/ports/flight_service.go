package ports

import "context"

type Flight struct {
	FlightNumber       string
	OriginAirport      string
	DestinationAirport string
	DateTime           string
	Duration           int
	Airline            string
}

type FlightService interface {
	SearchByOrigin(ctx context.Context, origin string) ([]Flight, error)
}
