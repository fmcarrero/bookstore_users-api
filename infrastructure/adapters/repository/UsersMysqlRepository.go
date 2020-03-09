package repository

import (
	"errors"
	"fmt"
	"github.com/fmcarrero/bookstore_users-api/domain/model"
	"github.com/fmcarrero/bookstore_users-api/infrastructure/adapters/repository/models"
	"github.com/fmcarrero/bookstore_users-api/infrastructure/mappers/users_mapper"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserMysqlRepository struct {
	Db *gorm.DB
}

func (userMysqlRepository *UserMysqlRepository) Save(user *model.User) error {

	var userDb models.UserDb
	userDb = users_mapper.UserToUserDb(*user)
	if userMysqlRepository.Db.Create(&userDb).Error != nil {
		return errors.New(fmt.Sprintf("can't work with %s", userDb.FirstName))
	}
	user.Id = userDb.ID

	return nil
}
