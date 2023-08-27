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
		id = "-1"
	}

	err := db.Where("ID = ?", id).Or("name = ?", name).Find(&locations).Error

	return locations, err
}

func GetLocationById(db *gorm.DB, id string) (models.Location, error) {
	var location models.Location

	err := db.First(&location, id).Error

	return location, err
}

func CreateLocation(db *gorm.DB, locations []models.Location) ([]models.Location, error) {
	if len(locations) == 0 {
		return locations, nil
	}

	err := db.Create(&locations).Error

	return locations, err
}

func UpdateLocation(db *gorm.DB, location models.Location) (models.Location, error) {
	err := db.Save(&location).Error

	return location, err
}

func DeleteLocation(db *gorm.DB, location models.Location) error {
	err := db.Delete(&location).Error

	return err
}
