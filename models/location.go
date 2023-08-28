package models

import (
	"time"

	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Name string `g:"required,min=3" gorm:"unique;not null;" json:"name"`
}

type LocationDto struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}
