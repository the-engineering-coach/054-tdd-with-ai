package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateAirportCode_Valid(t *testing.T) {
	err := ValidateAirportCode("JFK")
	assert.NoError(t, err)
}

func TestValidateAirportCode_Invalid_TooShort(t *testing.T) {
	err := ValidateAirportCode("JF")
	assert.Error(t, err)
	assert.Equal(t, "airport code must be exactly 3 characters", err.Error())
}
