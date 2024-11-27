package entities

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Username string	`gorm:"uniqueIndex" valid:"required~Username is required"`
	Password string
	Email    string	`gorm:"uniqueIndex" valid:"email~Email is invalid,required~Email is invalid"`
}