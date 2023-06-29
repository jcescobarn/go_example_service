package entities

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       uint
	Username string
	Name     string
	Password string
	Email    string
}

func (Users) TableName() string {
	return "users"
}
