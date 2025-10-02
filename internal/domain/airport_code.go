package domain

import "errors"

func ValidateAirportCode(code string) error {
	if len(code) != 3 {
		return errors.New("airport code must be exactly 3 characters")
	}
	return nil
}
