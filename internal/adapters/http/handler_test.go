package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"the-engineering-coach/tdd-with-ai/internal/ports"
)

type FlightResponse struct {
	FlightNumber       string    `json:"flight_number"`
	OriginAirport      string    `json:"origin_airport"`
	DestinationAirport string    `json:"destination_airport"`
	DepartureTime      time.Time `json:"departure_time"`
	Duration           int       `json:"duration"`
	Airline            string    `json:"airline"`
}

type MockFlightService struct {
	mock.Mock
}

func (m *MockFlightService) SearchByOrigin(ctx context.Context, origin string) ([]ports.Flight, error) {
	args := m.Called(ctx, origin)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]ports.Flight), args.Error(1)
}

func TestFlightHandler_NoOriginProvided(t *testing.T) {
	handler := NewFlightHandler(nil)

	req := httptest.NewRequest(http.MethodGet, "/flights", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var errResp ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&errResp)
	require.NoError(t, err)

	assert.Equal(t, "origin parameter is required", errResp.Error)
}

func TestFlightHandler_InvalidOriginAirportCode(t *testing.T) {
	mockService := new(MockFlightService)
	mockService.On("SearchByOrigin", mock.Anything, "INVALID").Return(nil, errors.New("invalid airport code"))

	handler := NewFlightHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/flights?origin=INVALID", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var errResp ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&errResp)
	require.NoError(t, err)

	assert.Equal(t, "invalid airport code", errResp.Error)
	mockService.AssertExpectations(t)
}

func TestFlightHandler_SearchByOrigin_Success(t *testing.T) {
	departureTime1, _ := time.Parse(time.RFC3339, "2025-10-02T10:00:00Z")
	departureTime2, _ := time.Parse(time.RFC3339, "2025-10-02T14:00:00Z")

	expectedFlights := []ports.Flight{
		{
			FlightNumber:       "AA100",
			OriginAirport:      "JFK",
			DestinationAirport: "LAX",
			DepartureTime:      departureTime1,
			Duration:           360,
			Airline:            "American Airlines",
		},
		{
			FlightNumber:       "UA200",
			OriginAirport:      "JFK",
			DestinationAirport: "SFO",
			DepartureTime:      departureTime2,
			Duration:           380,
			Airline:            "United Airlines",
		},
	}

	mockService := new(MockFlightService)
	mockService.On("SearchByOrigin", mock.Anything, "JFK").Return(expectedFlights, nil)

	handler := NewFlightHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/flights?origin=JFK", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var flights []FlightResponse
	err := json.NewDecoder(w.Body).Decode(&flights)
	require.NoError(t, err)

	assert.Len(t, flights, 2)

	assert.Equal(t, "AA100", flights[0].FlightNumber)
	assert.Equal(t, "JFK", flights[0].OriginAirport)
	assert.Equal(t, "LAX", flights[0].DestinationAirport)
	assert.Equal(t, departureTime1, flights[0].DepartureTime)
	assert.Equal(t, 360, flights[0].Duration)
	assert.Equal(t, "American Airlines", flights[0].Airline)

	assert.Equal(t, "UA200", flights[1].FlightNumber)
	assert.Equal(t, "JFK", flights[1].OriginAirport)
	assert.Equal(t, "SFO", flights[1].DestinationAirport)
	assert.Equal(t, departureTime2, flights[1].DepartureTime)
	assert.Equal(t, 380, flights[1].Duration)
	assert.Equal(t, "United Airlines", flights[1].Airline)

	mockService.AssertExpectations(t)
}
