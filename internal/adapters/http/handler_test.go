package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"the-engineering-coach/tdd-with-ai/internal/ports"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type MockFlightService struct {
	SearchByOriginFunc func(ctx context.Context, origin string) ([]ports.Flight, error)
}

func (m *MockFlightService) SearchByOrigin(ctx context.Context, origin string) ([]ports.Flight, error) {
	return m.SearchByOriginFunc(ctx, origin)
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
	mockService := &MockFlightService{
		SearchByOriginFunc: func(ctx context.Context, origin string) ([]ports.Flight, error) {
			return nil, errors.New("invalid airport code")
		},
	}

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
}
