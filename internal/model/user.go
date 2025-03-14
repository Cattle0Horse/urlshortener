package model

import (
	"github.com/Cattle0Horse/url-shortener/pkg/tools"
	"gorm.io/gorm"
)

type User struct {
	Model
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(100);not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Password = tools.PasswordEncrypt(u.Password)
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.Password = tools.PasswordEncrypt(u.Password)
	return nil
}
