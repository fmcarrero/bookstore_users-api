package factory

import (
	"github.com/fmcarrero/bookstore_users-api/application/commands"
	"github.com/fmcarrero/bookstore_users-api/domain/model"
)

func CreateUser(userCommand commands.UserCommand) (model.User, error) {
	var user model.User
	user, err := user.CreateUser(userCommand.FirstName, userCommand.LastName, userCommand.Email, userCommand.Password)
	return user, err
}
