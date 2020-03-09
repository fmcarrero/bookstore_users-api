package ports

import "github.com/fmcarrero/bookstore_users-api/domain/model"

type UsersRepository interface {
	Save(user *model.User) error
}
