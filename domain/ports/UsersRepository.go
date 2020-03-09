package ports

import "github.com/fmcarrero/bookstore_users-api/domain/model"

type UsersRepository interface {
	Save(user *model.User) error
	Get(userId int64) (model.User, error)
	Update(userId int64, user model.User) (*model.User, error)
	Delete(userId int64) error
	FindByStatus(status string) ([]model.User, error)
}
