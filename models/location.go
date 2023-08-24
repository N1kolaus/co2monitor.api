package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Name string `json:"name"`
}
