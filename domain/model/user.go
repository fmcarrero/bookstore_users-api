package model

import (
	"github.com/fmcarrero/bookstore_users-api/domain/utils/crypto_utils"
	"github.com/fmcarrero/bookstore_users-api/domain/utils/date_utils"
	"github.com/fmcarrero/bookstore_users-api/domain/validators"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64
	FirstName   string
	LastName    string
	Email       string
	DateCreated string
	Status      string
	Password    string
}

func (user *User) CreateUser(firstName string, lastName string, email string, password string) (User, error) {
	if err := validators.ValidateRequired(password, "Password should have some value"); err != nil {
		return User{}, err
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
		return User{}, err
	}

	return User{
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		DateCreated: date_utils.GetNowString(),
		Status:      StatusActive,
		Password:    crypto_utils.GetMd5(password),
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
		Password: crypto_utils.GetMd5(password),
	}, nil
}
