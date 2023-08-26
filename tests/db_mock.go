package tests

import (
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/fminister/co2monitor.api/models"
)

var Locations = []models.Location{
	{
		Name: "test location 1",
	},
	{
		Name: "test location 2",
	},
}

var CO2 = []models.Co2Data{
	{
		LocationID: 1,
		CO2:        1001,
		Temp:       20.1,
	},
	{
		LocationID: 1,
		CO2:        2001,
		Temp:       22.1,
	},
	{
		LocationID: 2,
		CO2:        1002,
		Temp:       20.2,
	},
	{
		LocationID: 2,
		CO2:        2002,
		Temp:       22.2,
	},
}

type BaseFixture struct {
	Db *gorm.DB
}

func (f *BaseFixture) Setup(t *testing.T) {
	var err error
	f.Db, err = gorm.Open(sqlite.Open(":memory:?_pragma=foreign_keys(1)"), &gorm.Config{})
	require.NoError(t, err)
	f.Db.AutoMigrate(&models.Location{}, &models.Co2Data{})
}

func (f *BaseFixture) Teardown(t *testing.T) {
	f.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Location{}, &models.Co2Data{})
}

func (f *BaseFixture) AddDummyData(t *testing.T) {
	f.Db.Create(&Locations)
	f.Db.Create(&CO2)
}
