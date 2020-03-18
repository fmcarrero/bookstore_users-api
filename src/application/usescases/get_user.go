package usescases

import (
	"github.com/fmcarrero/bookstore_users-api/src/domain/model"
	"github.com/fmcarrero/bookstore_users-api/src/domain/ports"
)

type GetUserUseCase interface {
	Handler(userId int64) (model.User, error)
}

type UseCaseGetUser struct {
	UserRepository ports.UsersRepository
}

func (useCaseGetUser *UseCaseGetUser) Handler(userId int64) (model.User, error) {

	user, err := useCaseGetUser.UserRepository.Get(userId)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
