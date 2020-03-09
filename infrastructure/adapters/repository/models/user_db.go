package models

type UserDb struct {
	ID          int64
	FirstName   string
	LastName    string
	Email       string
	DateCreated string `gorm:"type:varchar(45)"`
}

func (UserDb) TableName() string {
	return "user"
}
