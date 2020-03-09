package users_mapper

import (
	"github.com/fmcarrero/bookstore_users-api/domain/model"
	"github.com/fmcarrero/bookstore_users-api/domain/utils/date_utils"
	"github.com/fmcarrero/bookstore_users-api/infrastructure/adapters/repository/models"
)

func UserToUserDb(user model.User) models.UserDb {

	var userDb models.UserDb
	userDb.FirstName = user.FirstName
	userDb.LastName = user.LastName
	userDb.Email = user.Email
	userDb.DateCreated = date_utils.GetNowString()
	return userDb
}
