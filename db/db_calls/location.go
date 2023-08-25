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

func GetLocationById(db *gorm.DB, locationId string) (models.Location, error) {
	var location models.Location

	err := db.First(&location, locationId).Error

	return location, err
}
