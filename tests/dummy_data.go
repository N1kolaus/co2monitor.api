package tests

import (
	"encoding/json"
	"time"

	"github.com/fminister/co2monitor.api/models"
	"gorm.io/gorm"
)

var Locations = []models.Location{
	{
		Model: gorm.Model{
			CreatedAt: time.Date(2023, 8, 1, 12, 30, 0, 100, time.Local),
			UpdatedAt: time.Date(2023, 8, 1, 12, 30, 0, 100, time.Local),
		},
		Name: "test location 1",
	},
	{
		Model: gorm.Model{
			CreatedAt: time.Date(2023, 8, 1, 12, 30, 0, 100, time.Local).Add(-time.Minute * 1),
			UpdatedAt: time.Date(2023, 8, 1, 12, 30, 0, 100, time.Local).Add(-time.Minute * 1),
		},
		Name: "test location 2",
	},
}

var CO2 = []models.Co2Data{
	{
		Model: gorm.Model{
			CreatedAt: time.Date(2023, 8, 1, 12, 30, 0, 100, time.Local),
			UpdatedAt: time.Date(2023, 8, 1, 12, 30, 0, 100, time.Local),
		},
		LocationID: 1,
		CO2:        1001,
		Temp:       20.1,
	},
	{
		Model: gorm.Model{
			CreatedAt: time.Date(2023, 8, 1, 12, 30, 0, 100, time.Local).Add(-time.Minute * 1),
			UpdatedAt: time.Date(2023, 8, 1, 12, 30, 0, 100, time.Local).Add(-time.Minute * 1),
		},
		LocationID: 1,
		CO2:        2001,
		Temp:       22.1,
	},
	{
		Model: gorm.Model{
			CreatedAt: time.Date(2023, 8, 1, 12, 30, 0, 100, time.Local),
			UpdatedAt: time.Date(2023, 8, 1, 12, 30, 0, 100, time.Local),
		},
		LocationID: 2,
		CO2:        1002,
		Temp:       20.2,
	},
	{
		Model: gorm.Model{
			CreatedAt: time.Date(2023, 8, 1, 12, 30, 0, 100, time.Local).Add(-time.Minute * 1),
			UpdatedAt: time.Date(2023, 8, 1, 12, 30, 0, 100, time.Local).Add(-time.Minute * 1),
		},
		LocationID: 2,
		CO2:        2002,
		Temp:       22.2,
	},
}

func LocationsToJSON(locations []models.Location) []byte {
	locationsJSON, err := json.Marshal(locations)
	if err != nil {
		panic(err)
	}

	return locationsJSON
}

func CO2ToJSON(co2 []models.Co2Data) []byte {
	co2JSON, err := json.Marshal(co2)
	if err != nil {
		panic(err)
	}

	return co2JSON
}
