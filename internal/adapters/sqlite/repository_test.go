package sqlite

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFlightRepository_FindByOrigin(t *testing.T) {
	dbPath := "test_flights.db"
	defer os.Remove(dbPath)

	db, err := sql.Open("sqlite3", dbPath)
	require.NoError(t, err)
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE flights (
			flight_number TEXT,
			origin_airport TEXT,
			destination_airport TEXT,
			departure_time TEXT,
			duration INTEGER,
			airline TEXT
		)
	`)
	require.NoError(t, err)

	departureTime1 := time.Date(2025, 10, 2, 10, 0, 0, 0, time.UTC)
	departureTime2 := time.Date(2025, 10, 2, 14, 0, 0, 0, time.UTC)

	_, err = db.Exec(`
		INSERT INTO flights (flight_number, origin_airport, destination_airport, departure_time, duration, airline)
		VALUES (?, ?, ?, ?, ?, ?)
	`, "AA100", "JFK", "LAX", departureTime1.Format(time.RFC3339), 360, "American Airlines")
	require.NoError(t, err)

	_, err = db.Exec(`
		INSERT INTO flights (flight_number, origin_airport, destination_airport, departure_time, duration, airline)
		VALUES (?, ?, ?, ?, ?, ?)
	`, "UA200", "JFK", "SFO", departureTime2.Format(time.RFC3339), 380, "United Airlines")
	require.NoError(t, err)

	repo := NewFlightRepository(db)
	ctx := context.Background()

	flights, err := repo.FindByOrigin(ctx, "JFK")
	require.NoError(t, err)

	assert.Len(t, flights, 2)
	assert.Equal(t, "AA100", flights[0].FlightNumber)
	assert.Equal(t, "JFK", flights[0].OriginAirport)
	assert.Equal(t, "LAX", flights[0].DestinationAirport)
	assert.Equal(t, departureTime1, flights[0].DepartureTime)
	assert.Equal(t, 360, flights[0].Duration)
	assert.Equal(t, "American Airlines", flights[0].Airline)
}
