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

func TestValidateAirportCode_Invalid_TooLong(t *testing.T) {
	err := ValidateAirportCode("JFKX")
	assert.Error(t, err)
	assert.Equal(t, "airport code must be exactly 3 characters", err.Error())
}

func TestValidateAirportCode_Invalid_Lowercase(t *testing.T) {
	err := ValidateAirportCode("jfk")
	assert.Error(t, err)
	assert.Equal(t, "airport code must be uppercase letters only", err.Error())
}

func TestValidateAirportCode_Invalid_MixedCase(t *testing.T) {
	err := ValidateAirportCode("JfK")
	assert.Error(t, err)
	assert.Equal(t, "airport code must be uppercase letters only", err.Error())
}

func TestValidateAirportCode_Invalid_WithNumbers(t *testing.T) {
	err := ValidateAirportCode("JF1")
	assert.Error(t, err)
	assert.Equal(t, "airport code must be uppercase letters only", err.Error())
}

func TestValidateAirportCode_Invalid_WithSpecialChars(t *testing.T) {
	err := ValidateAirportCode("JF-")
	assert.Error(t, err)
	assert.Equal(t, "airport code must be uppercase letters only", err.Error())
}
