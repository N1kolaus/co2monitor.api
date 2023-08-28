package models

import (
	"time"

	"gorm.io/gorm"
)

type Co2Data struct {
	gorm.Model
	CO2        int     `g:"required" gorm:"not null;" json:"co2"`
	Temp       float32 `g:"required" gorm:"not null;" json:"temp"`
	LocationID int     `g:"required" gorm:"not null;" json:"location_id"`
	Location   Location
}

type Co2DataDto struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CO2        int       `json:"co2"`
	Temp       float32   `json:"temp"`
	LocationID int       `json:"location_id"`
}
