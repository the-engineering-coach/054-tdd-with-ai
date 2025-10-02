//go:build e2e
// +build e2e

package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Flight struct {
	FlightNumber       string    `json:"flight_number"`
	OriginAirport      string    `json:"origin_airport"`
	DestinationAirport string    `json:"destination_airport"`
	DepartureTime      time.Time `json:"departure_time"`
	Duration           int       `json:"duration"` // in minutes
	Airline            string    `json:"airline"`
}

func TestSearchFlightsByOrigin(t *testing.T) {
	// Setup test database
	dbPath := "test_flights.db"
	defer os.Remove(dbPath)

	db, err := sql.Open("sqlite3", dbPath)
	require.NoError(t, err)
	defer db.Close()

	// Create flights table
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

	// Insert 3 test flights with JFK origin
	departureTime1 := time.Date(2025, 10, 2, 10, 0, 0, 0, time.UTC)
	departureTime2 := time.Date(2025, 10, 2, 14, 0, 0, 0, time.UTC)
	departureTime3 := time.Date(2025, 10, 2, 16, 0, 0, 0, time.UTC)

	flights := []struct {
		FlightNumber       string
		OriginAirport      string
		DestinationAirport string
		DepartureTime      time.Time
		Duration           int
		Airline            string
	}{
		{"AA100", "JFK", "LAX", departureTime1, 360, "American Airlines"},
		{"UA200", "JFK", "SFO", departureTime2, 380, "United Airlines"},
		{"DL300", "JFK", "ORD", departureTime3, 150, "Delta Airlines"},
	}

	for _, flight := range flights {
		_, err = db.Exec(`
			INSERT INTO flights (flight_number, origin_airport, destination_airport, departure_time, duration, airline)
			VALUES (?, ?, ?, ?, ?, ?)
		`, flight.FlightNumber, flight.OriginAirport, flight.DestinationAirport, flight.DepartureTime.Format(time.RFC3339), flight.Duration, flight.Airline)
		require.NoError(t, err)
	}

	// Start the server
	server := NewServer(db)
	ts := httptest.NewServer(server)
	defer ts.Close()

	// Make HTTP request to the running server
	res, err := http.Get(ts.URL + "/flights?origin=JFK")
	require.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	var resultFlights []Flight
	err = json.NewDecoder(res.Body).Decode(&resultFlights)
	require.NoError(t, err)

	assert.Len(t, resultFlights, 3)

	for _, flight := range resultFlights {
		assert.Equal(t, "JFK", flight.OriginAirport)
	}
}
