package models

import "gorm.io/gorm"

type Co2Data struct {
	gorm.Model
	CO2        int     `g:"required" gorm:"not null;" json:"co2"`
	Temp       float32 `g:"required" gorm:"not null;" json:"temp"`
	LocationID int     `g:"required" gorm:"not null;" json:"location_id"`
	Location   Location
}
