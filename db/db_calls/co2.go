package db_calls

import (
	"errors"
	"time"

	"github.com/fminister/co2monitor.api/models"
	"gorm.io/gorm"
)

func GetCo2DataByTimeFrame(db *gorm.DB, locationId int, hours time.Duration) ([]models.Co2Data, error) {
	var co2Data []models.Co2Data

	err := db.Where("location_id = ? AND created_at > ?", locationId, time.Now().Add(-hours)).Find(&co2Data).Error

	return co2Data, err
}

func GetLatestCo2Data(db *gorm.DB, locationId string) (models.Co2Data, error) {
	var co2Data models.Co2Data

	err := db.Order("created_at desc").Where("location_id = ?", locationId).First(&co2Data).Error

	return co2Data, err
}

func CreateCo2Data(db *gorm.DB, co2Data []models.Co2Data) ([]models.Co2Data, error) {
	if len(co2Data) == 0 {
		return co2Data, errors.New("Empty list of co2 data to insert")
	}

	err := db.Create(&co2Data).Error

	return co2Data, err
}
