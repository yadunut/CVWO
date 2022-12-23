package utils

import (
	"errors"
	"fmt"
	"regexp"
)

func IsAlphaNumeric(s string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9]*$")
	return re.MatchString(s)
}

func ValidateUsername(s string) error {
	if len(s) < 3 {
		return errors.New("username must be greater than 3 characters")
	}
	if len(s) > 10 {
		return errors.New("username must be shorter than 10 characters")
	}
	if !IsAlphaNumeric(s) {
		return errors.New(fmt.Sprintf("%s invalid, only alphabets and numbers allowed", s))
	}
	return nil
}

func ValidatePassword(s string) error {
	if len(s) < 8 {
		return errors.New("password must be longer than 8 characters")
	}
	if len(s) > 30 {
		return errors.New("password must be shorter than 30 characters")
	}

	return nil
}
