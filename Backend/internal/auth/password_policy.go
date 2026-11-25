package auth

import (
	"errors"
	"unicode/utf8"
)

const MinPasswordLength = 8

var ErrPasswordTooShort = errors.New("password is too short")

/* ValidatePassword enforces predefined password policies during user registration */
func ValidatePassword(password string) error {
	if utf8.RuneCountInString(password) < MinPasswordLength {
		return ErrPasswordTooShort
	}

	return nil
}
