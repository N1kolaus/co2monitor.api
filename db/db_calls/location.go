package db_calls

import (
	"github.com/fminister/co2monitor.api/models"
	"gorm.io/gorm"
)

func GetLocation(db *gorm.DB) ([]models.Location, error) {
	var locations []models.Location

	err := db.Find(&locations).Error

	return locations, err
}

func GetLocationBySearch(db *gorm.DB, id string, name string) ([]models.Location, error) {
	var locations []models.Location

	if id == "" {
		id = "0"
	}

	err := db.Where("ID = ?", id).Or("name = ?", name).Find(&locations).Error

	return locations, err
}
