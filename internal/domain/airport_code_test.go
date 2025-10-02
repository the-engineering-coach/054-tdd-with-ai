package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateAirportCode_Valid(t *testing.T) {
	err := ValidateAirportCode("JFK")
	assert.NoError(t, err)
}
