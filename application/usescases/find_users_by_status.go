package usescases

import (
	"github.com/fmcarrero/bookstore_users-api/domain/model"
	"github.com/fmcarrero/bookstore_users-api/domain/ports"
)

type FindUsersByStatusUseCase interface {
	Handler(status string) ([]model.User, error)
}
type UseCaseFindUserByStatus struct {
	UserRepository ports.UsersRepository
}

func (useCaseFindUserByStatus *UseCaseFindUserByStatus) Handler(status string) ([]model.User, error) {
	return useCaseFindUserByStatus.UserRepository.FindByStatus(status)

}
