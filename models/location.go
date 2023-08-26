package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Name string `g:"required,min=3" gorm:"unique;not null;" json:"name"`
}
