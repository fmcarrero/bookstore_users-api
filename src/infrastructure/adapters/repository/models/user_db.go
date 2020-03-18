package models

import "time"

type UserDb struct {
	ID          int64
	FirstName   string
	LastName    string
	Email       string
	DateCreated time.Time
	Status      string `gorm:"type:varchar(45)"`
	Password    string `gorm:"type:varchar(32)"`
}

func (UserDb) TableName() string {
	return "user"
}
