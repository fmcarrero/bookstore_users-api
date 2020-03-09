package model

import (
	"github.com/fmcarrero/bookstore_users-api/domain/utils/date_utils"
	"github.com/fmcarrero/bookstore_users-api/domain/validators"
)

type User struct {
	Id          int64
	FirstName   string
	LastName    string
	Email       string
	DateCreated string
}

func (user *User) CreateUser(firstName string, lastName string, email string) (User, error) {
	if err := validators.ValidateRequired(firstName, "FirstName should have some value"); err != nil {
		return User{}, err
	}
	if err := validators.ValidateRequired(lastName, "lastName should have some value"); err != nil {
		return User{}, err
	}
	if err := validators.ValidateRequired(email, "email should have some value"); err != nil {
		return User{}, err
	}
	if err := validators.ValidateEmail(email, "invalid email"); err != nil {
		return User{}, err
	}
	return User{
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		DateCreated: date_utils.GetNowString(),
	}, nil
}
