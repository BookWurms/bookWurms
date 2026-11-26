package auth

import (
	"errors"
	"unicode/utf8"
)

const (
	MinPasswordLength = 8
	MaxPasswordLength = 40
)

var ErrPasswordTooShort = errors.New("password is too short")

/* ValidatePassword enforces predefined password policies during user registration */
func ValidatePassword(password string) error {
	length := utf8.RuneCountInString(password)

	if length < MinPasswordLength {
		return ErrPasswordTooShort
	}
	if length > MaxPasswordLength {
		return ErrPasswordTooLong
	}

	return nil
}
