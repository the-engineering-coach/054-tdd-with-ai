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

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Flight struct {
	FlightNumber       string `json:"flight_number"`
	OriginAirport      string `json:"origin_airport"`
	DestinationAirport string `json:"destination_airport"`
	DateTime           string `json:"date_time"`
	Duration           int    `json:"duration"` // in minutes
	Airline            string `json:"airline"`
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
			date_time TEXT,
			duration INTEGER,
			airline TEXT
		)
	`)
	require.NoError(t, err)

	// Insert 3 test flights with JFK origin
	flights := []Flight{
		{FlightNumber: "AA100", OriginAirport: "JFK", DestinationAirport: "LAX", DateTime: "2025-10-02T10:00:00Z", Duration: 360, Airline: "American Airlines"},
		{FlightNumber: "UA200", OriginAirport: "JFK", DestinationAirport: "SFO", DateTime: "2025-10-02T14:00:00Z", Duration: 380, Airline: "United Airlines"},
		{FlightNumber: "DL300", OriginAirport: "JFK", DestinationAirport: "ORD", DateTime: "2025-10-02T16:00:00Z", Duration: 150, Airline: "Delta Airlines"},
	}

	for _, flight := range flights {
		_, err = db.Exec(`
			INSERT INTO flights (flight_number, origin_airport, destination_airport, date_time, duration, airline)
			VALUES (?, ?, ?, ?, ?, ?)
		`, flight.FlightNumber, flight.OriginAirport, flight.DestinationAirport, flight.DateTime, flight.Duration, flight.Airline)
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
