package models

import "gorm.io/gorm"

type Co2Data struct {
	gorm.Model
	CO2        int     `json:"co2"`
	Temp       float32 `json:"temp"`
	LocationID int     `json:"location_id"`
	Location   Location
}
