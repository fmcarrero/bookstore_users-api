package usescases

import (
	"github.com/fmcarrero/bookstore_users-api/src/domain/ports"
)

type DeleteUserUseCase interface {
	Handler(userId int64) error
}

type UseCaseDeleteUser struct {
	UserRepository ports.UsersRepository
}

func (useCaseDeleteUser *UseCaseDeleteUser) Handler(userId int64) error {
	user, getUserError := useCaseDeleteUser.UserRepository.Get(userId)
	if getUserError != nil {
		return getUserError
	}
	err := useCaseDeleteUser.UserRepository.Delete(user.Id)
	return err
}
