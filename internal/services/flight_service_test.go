package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"the-engineering-coach/tdd-with-ai/internal/ports"
)

type MockFlightRepository struct {
	mock.Mock
}

func (m *MockFlightRepository) FindByOrigin(ctx context.Context, origin string) ([]ports.Flight, error) {
	args := m.Called(ctx, origin)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]ports.Flight), args.Error(1)
}

func TestFlightService_SearchByOrigin_ValidCode(t *testing.T) {
	mockRepo := new(MockFlightRepository)
	expectedFlights := []ports.Flight{
		{FlightNumber: "AA100", OriginAirport: "JFK"},
	}
	mockRepo.On("FindByOrigin", mock.Anything, "JFK").Return(expectedFlights, nil)

	service := NewFlightService(mockRepo)
	ctx := context.Background()

	flights, err := service.SearchByOrigin(ctx, "JFK")
	require.NoError(t, err)

	assert.Len(t, flights, 1)
	assert.Equal(t, "AA100", flights[0].FlightNumber)
	mockRepo.AssertExpectations(t)
}
