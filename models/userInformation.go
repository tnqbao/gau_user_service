package models

import (
	"gorm.io/gorm"
)

type UserInformaiton struct {
	gorm.Model
	UserId      uint   `gorm:"primaryKey" json:"user_id"`
	Name        string `json:"name"`
	Phone       string `gorm:"unique" json:"phone"`
	Email       string `gorm:"unique" json:"email"`
	DateOfBirth string `gorm:"unique" json:"date_of_birth"`
}
