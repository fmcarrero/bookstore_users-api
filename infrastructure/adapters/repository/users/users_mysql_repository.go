package users

import (
	"errors"
	"fmt"
	"github.com/fmcarrero/bookstore_users-api/domain/model"
	"github.com/fmcarrero/bookstore_users-api/infrastructure/adapters/repository/models"
	"github.com/fmcarrero/bookstore_users-api/infrastructure/mappers/users_mapper"
	"github.com/fmcarrero/bookstore_utils-go/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserMysqlRepository struct {
	Db *gorm.DB
}

func (userMysqlRepository *UserMysqlRepository) Save(user *model.User) error {

	var userDb models.UserDb
	userDb = users_mapper.UserToUserDb(*user)
	if err := userMysqlRepository.Db.Create(&userDb).Error; err != nil {
		logger.Error(fmt.Sprintf("can't work with %s", userDb.FirstName), err)
		return errors.New(fmt.Sprintf("can't work with %s", userDb.FirstName))
	}
	user.Id = userDb.ID
	user.Password = ""
	return nil
}

func (userMysqlRepository *UserMysqlRepository) Get(userId int64) (model.User, error) {
	var userDb models.UserDb
	if userMysqlRepository.Db.First(&userDb, userId).Error != nil {
		return model.User{}, errors.New(fmt.Sprintf("user not found %v", userId))
	}
	fmt.Println(userDb)
	user := users_mapper.UserDbToUser(userDb)
	return user, nil
}

func (userMysqlRepository *UserMysqlRepository) FindByStatus(status string) ([]model.User, error) {
	var usersDb []models.UserDb
	if userMysqlRepository.Db.Where(&models.UserDb{Status: status}).Find(&usersDb).Error != nil {
		return nil, errors.New(fmt.Sprintf("no users matching status"))
	}
	if len(usersDb) <= 0 {
		return nil, errors.New(fmt.Sprintf("no users matching status"))
	}
	users := users_mapper.UsersDbToUsers(usersDb)
	return users, nil
}

func (userMysqlRepository *UserMysqlRepository) Update(userId int64, user model.User) (*model.User, error) {
	var current models.UserDb
	if userMysqlRepository.Db.First(&current, userId).RecordNotFound() {
		return nil, errors.New(fmt.Sprintf("user not found %v", userId))
	}
	if userMysqlRepository.Db.Model(&current).Update(models.UserDb{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email}).Error != nil {
		return nil, errors.New(fmt.Sprintf("error when updated user %v", userId))
	}
	userUpdated := users_mapper.UserDbToUser(current)
	return &userUpdated, nil
}

func (userMysqlRepository *UserMysqlRepository) Delete(userId int64) error {

	var current models.UserDb
	current.ID = userId
	if userMysqlRepository.Db.Delete(current).Error != nil {
		return errors.New(fmt.Sprintf("cannot delete user  %v", userId))
	}
	return nil
}

func (userMysqlRepository *UserMysqlRepository) Login(user model.User) (*model.User, error) {
	var userDb models.UserDb
	if userMysqlRepository.Db.Where(&models.UserDb{Email: user.Email, Password: user.Password, Status: model.StatusActive}).First(&userDb).Error != nil {
		return nil, errors.New(fmt.Sprintf("no users matching "))
	}
	userFound := users_mapper.UserDbToUser(userDb)
	return &userFound, nil
}
