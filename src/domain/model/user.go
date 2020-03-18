package model

import (
	"github.com/fmcarrero/bookstore_users-api/src/domain/exceptions"
	"github.com/fmcarrero/bookstore_users-api/src/domain/validators"
	"github.com/fmcarrero/bookstore_utils-go/crypto"
	"github.com/fmcarrero/bookstore_utils-go/date"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

func (user *User) CreateUser(firstName string, lastName string, email string, password string) (User, error) {
	if err := validators.ValidateRequired(password, "Password should have some value"); err != nil {
		return User{}, exceptions.InvalidPassword{ErrMessage: err.Error()}
	}
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
		return User{}, exceptions.InvalidEmail{ErrMessage: err.Error()}
	}

	return User{
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		DateCreated: date.GetNowString(),
		Status:      StatusActive,
		Password:    crypto.GetMd5(password),
	}, nil
}
func (user *User) CreateUserLogin(email string, password string) (User, error) {
	if err := validators.ValidateRequired(password, "Password should have some value"); err != nil {
		return User{}, err
	}
	if err := validators.ValidateRequired(email, "email should have some value"); err != nil {
		return User{}, err
	}
	if err := validators.ValidateEmail(email, "invalid email"); err != nil {
		return User{}, err
	}
	return User{
		Email:    email,
		Password: crypto.GetMd5(password),
	}, nil
}
