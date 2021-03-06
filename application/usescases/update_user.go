package usescases

import (
	"github.com/fmcarrero/bookstore_users-api/application/commands"
	"github.com/fmcarrero/bookstore_users-api/application/factory"
	"github.com/fmcarrero/bookstore_users-api/domain/model"
	"github.com/fmcarrero/bookstore_users-api/domain/ports"
)

type UpdateUserUseCase interface {
	Handler(userId int64, userCommand commands.UserCommand) (*model.User, error)
}

type UseCaseUpdateUser struct {
	UserRepository ports.UsersRepository
}

func (useCaseUpdateUser *UseCaseUpdateUser) Handler(userId int64, userCommand commands.UserCommand) (*model.User, error) {
	user, err := factory.CreateUser(userCommand)
	if err != nil {
		return nil, err
	}
	userUpdated, err := useCaseUpdateUser.UserRepository.Update(userId, user)
	if err != nil {
		return userUpdated, err
	}
	return userUpdated, nil
}
