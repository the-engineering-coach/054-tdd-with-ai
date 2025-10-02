package domain

import (
	"errors"
	"strings"
	"unicode"
)

func ValidateAirportCode(code string) error {
	if len(code) != 3 {
		return errors.New("airport code must be exactly 3 characters")
	}

	if code != strings.ToUpper(code) {
		return errors.New("airport code must be uppercase letters only")
	}

	for _, char := range code {
		if !unicode.IsLetter(char) {
			return errors.New("airport code must be uppercase letters only")
		}
	}

	return nil
}
