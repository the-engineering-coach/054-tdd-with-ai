package sqlite

import (
	"context"
	"database/sql"
	"time"

	"the-engineering-coach/tdd-with-ai/internal/ports"
)

type FlightRepository struct {
	db *sql.DB
}

func NewFlightRepository(db *sql.DB) *FlightRepository {
	return &FlightRepository{db: db}
}

func (r *FlightRepository) FindByOrigin(ctx context.Context, origin string) ([]ports.Flight, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT flight_number, origin_airport, destination_airport, departure_time, duration, airline FROM flights WHERE origin_airport = ?", origin)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flights []ports.Flight
	for rows.Next() {
		var flight ports.Flight
		var departureTimeStr string
		err := rows.Scan(&flight.FlightNumber, &flight.OriginAirport, &flight.DestinationAirport, &departureTimeStr, &flight.Duration, &flight.Airline)
		if err != nil {
			return nil, err
		}
		flight.DepartureTime, err = time.Parse(time.RFC3339, departureTimeStr)
		if err != nil {
			return nil, err
		}
		flights = append(flights, flight)
	}

	return flights, nil
}
