package models

import "gorm.io/gorm"

type Co2Data struct {
	gorm.Model
	CO2      int      `json:"co2"`
	Temp     float32  `json:"temp"`
	Location Location `json:"location"`
}
