package domain

import (
	"errors"
	"unicode"
)

var (
	ErrInvalidLength       = errors.New("airport code must be exactly 3 characters")
	ErrNotUppercaseLetters = errors.New("airport code must be uppercase letters only")
)

func ValidateAirportCode(code string) error {
	if len(code) != 3 {
		return ErrInvalidLength
	}

	for _, char := range code {
		if !unicode.IsLetter(char) || !unicode.IsUpper(char) {
			return ErrNotUppercaseLetters
		}
	}

	return nil
}
