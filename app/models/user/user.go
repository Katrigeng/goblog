package user

import "goblog/app/models"

type User struct {
	models.BaseModel

	Name            string `gorm:"type:varchar(255);not null;unique"`
	Email           string `gorm:"type:varchar(255);unique;"`
	Password        string `gorm:"type:varchar(255)"`
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}
