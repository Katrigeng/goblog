package user

import (
	"goblog/app/models"
	"goblog/pkg/password"
)

type User struct {
	models.BaseModel

	Name            string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email           string `gorm:"type:varchar(255);unique;" valid:"email"`
	Password        string `gorm:"type:varchar(255)" valid:"password"`
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}

func (User *User) ComparePassword(_password string) bool {
	return password.CheckHash(_password, User.Password)
}

// Link 方法用来生成用户链接
func (user *User) Link() string {
	return ""
}
