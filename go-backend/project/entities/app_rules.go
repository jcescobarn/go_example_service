package entities

import (
	"gorm.io/gorm"
)

type AppRules struct {
	gorm.Model
	ID      uint
	Rule_id string
	App_id  int
}

func (AppRules) TableName() string {
	return "app_rules"
}
