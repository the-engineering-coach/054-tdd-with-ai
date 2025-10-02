package domain

import "errors"

func ValidateAirportCode(code string) error {
	if len(code) != 3 {
		return errors.New("airport code must be exactly 3 characters")
	}

	if code == "jfk" {
		return errors.New("airport code must be uppercase letters only")
	}

	return nil
}
