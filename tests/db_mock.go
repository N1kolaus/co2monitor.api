package tests

import (
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/fminister/co2monitor.api/models"
)

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
