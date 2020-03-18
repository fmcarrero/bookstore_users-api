package marshallers

import (
	"encoding/json"
	"fmt"
	"github.com/fmcarrero/bookstore_users-api/src/domain/model"
)

type PublicUser struct {
	Id          int64  `json:"user_id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}
type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

func Marshall(isPublic bool, user model.User) interface{} {
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	userJson, errUn := json.Marshal(user)
	fmt.Println(errUn)
	fmt.Println(user)
	var privateUser PrivateUser
	_ = json.Unmarshal(userJson, &privateUser)
	return privateUser
}

func MarshallArray(isPublic bool, users []model.User) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = Marshall(isPublic, user)
	}
	return result
}
