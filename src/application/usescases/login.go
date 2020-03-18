package usescases

import (
	"github.com/fmcarrero/bookstore_users-api/src/application/commands"
	"github.com/fmcarrero/bookstore_users-api/src/application/factory"
	"github.com/fmcarrero/bookstore_users-api/src/domain/model"
	"github.com/fmcarrero/bookstore_users-api/src/domain/ports"
)

type LoginUseCase interface {
	Handler(loginCommand commands.LoginCommand) (*model.User, error)
}

type UseCaseLogin struct {
	UserRepository ports.UsersRepository
}

func (useCaseGetUser *UseCaseLogin) Handler(loginCommand commands.LoginCommand) (*model.User, error) {
	user, err := factory.CreateUserLogin(loginCommand)
	if err != nil {
		return nil, err
	}
	return useCaseGetUser.UserRepository.Login(user)

}
