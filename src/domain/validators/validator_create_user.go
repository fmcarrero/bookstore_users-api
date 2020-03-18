package validators

import (
	"errors"
	"regexp"
	"strings"
)

func ValidateRequired(field string, message string) error {
	if strings.TrimSpace(field) == "" {
		return errors.New(message)
	}
	return nil
}
func ValidateEmail(email string, message string) error {
	expression := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	valid := expression.MatchString(email)
	if !valid {
		return errors.New(message)
	}
	return nil
}
