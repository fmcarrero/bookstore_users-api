package factory

import (
	"github.com/fmcarrero/bookstore_users-api/src/application/commands"
	"github.com/fmcarrero/bookstore_users-api/src/domain/model"
)

func CreateUser(userCommand commands.UserCommand) (model.User, error) {
	var user model.User
	user, err := user.CreateUser(userCommand.FirstName, userCommand.LastName, userCommand.Email, userCommand.Password)
	return user, err
}

func CreateUserLogin(command commands.LoginCommand) (model.User, error) {
	var user model.User
	user, err := user.CreateUserLogin(command.Email, command.Password)
	return user, err

}
