package users_mapper

import (
	"github.com/fmcarrero/bookstore_users-api/src/domain/model"
	"github.com/fmcarrero/bookstore_users-api/src/infrastructure/adapters/repository/models"
	"github.com/fmcarrero/bookstore_utils-go/date"
	"time"
)

func UserToUserDb(user model.User) models.UserDb {

	now, _ := time.Parse(date.ApiDbLayout, date.GetNowDBFormatNow())
	var userDb models.UserDb
	userDb.FirstName = user.FirstName
	userDb.LastName = user.LastName
	userDb.Email = user.Email
	userDb.DateCreated = now
	userDb.Status = user.Status
	userDb.Password = user.Password
	return userDb
}

func UserDbToUser(userDb models.UserDb) model.User {
	var user model.User
	user.Id = userDb.ID
	user.FirstName = userDb.FirstName
	user.LastName = userDb.LastName
	user.Email = userDb.Email
	user.DateCreated = date.GetDateString(userDb.DateCreated)
	user.Status = userDb.Status
	return user
}

func UsersDbToUsers(usersDb []models.UserDb) []model.User {
	var users []model.User
	for _, userDb := range usersDb {
		user := UserDbToUser(userDb)
		users = append(users, user)
	}
	return users
}
