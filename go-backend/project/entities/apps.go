package entities

import (
	"gorm.io/gorm"
)

type Apps struct {
	gorm.Model
	ID          uint
	Name        string
	Description string
}

func (Apps) TableName() string {
	return "apps"
}
