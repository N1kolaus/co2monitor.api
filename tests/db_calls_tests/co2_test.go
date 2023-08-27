package tests

import (
	"testing"
	"time"

	"github.com/fminister/co2monitor.api/db/db_calls"
	"github.com/fminister/co2monitor.api/models"
	"github.com/fminister/co2monitor.api/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestGetCo2DataByTimeFrame_ShouldReturnLastTwoValues(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	// add dummy data or creating newData will fail because of foreign key constraint
	f.AddDummyData(t)
	defer f.Teardown(t)

	newData := []models.Co2Data{
		{
			Model: gorm.Model{
				CreatedAt: time.Now().Add(-1 * time.Hour),
				UpdatedAt: time.Now().Add(-1 * time.Hour),
			},
			CO2:        500,
			Temp:       24.5,
			LocationID: 1,
		},
		{
			Model: gorm.Model{
				CreatedAt: time.Now().Add(-2 * time.Hour),
				UpdatedAt: time.Now().Add(-2 * time.Hour),
			},
			CO2:        505,
			Temp:       25.5,
			LocationID: 1,
		},
	}
	f.Db.Create(&newData)
	result, err := db_calls.GetCo2DataByTimeFrame(f.Db, 1, 24*time.Hour)

	require.NoError(t, err)
	assert.Equal(t, 2, len(newData))
	assert.Equal(t, 2, len(result))
	assert.Equal(t, result[0].CO2, newData[0].CO2)
	assert.Equal(t, result[0].Temp, newData[0].Temp)
	assert.Equal(t, result[1].CO2, newData[1].CO2)
	assert.Equal(t, result[1].Temp, newData[1].Temp)
}

func TestGetCo2DataByTimeFrame_ShouldReturnEmptyList(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	defer f.Teardown(t)

	result, err := db_calls.GetCo2DataByTimeFrame(f.Db, 1, 24*time.Hour)

	require.NoError(t, err)
	assert.Equal(t, 0, len(result))
}

func TestGetLatestCo2Data_ShouldReturnLastValue(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	defer f.Teardown(t)

	result, err := db_calls.GetLatestCo2Data(f.Db, "1")

	require.NoError(t, err)
	assert.Equal(t, result.CO2, tests.CO2[0].CO2)
	assert.Equal(t, result.Temp, tests.CO2[0].Temp)
}

func TestCreateCo2Data_ShouldCreateSingleValue(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	defer f.Teardown(t)

	newData := []models.Co2Data{
		{
			Model: gorm.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			CO2:        500,
			Temp:       24.5,
			LocationID: 1,
		},
	}
	beforeInsert := []models.Co2Data{}
	f.Db.Find(&beforeInsert)
	result, err := db_calls.CreateCo2Data(f.Db, newData)
	afterInsert := []models.Co2Data{}
	f.Db.Find(&afterInsert)

	require.NoError(t, err)
	assert.Equal(t, 4, len(beforeInsert))
	assert.Equal(t, 1, len(result))
	assert.Equal(t, 5, len(afterInsert))
}

func TestCreateCo2Data_ShouldCreateMultipleValues(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	defer f.Teardown(t)

	newData := []models.Co2Data{
		{
			Model: gorm.Model{
				CreatedAt: time.Now().Add(-1 * time.Hour),
				UpdatedAt: time.Now().Add(-1 * time.Hour),
			},
			CO2:        500,
			Temp:       24.5,
			LocationID: 1,
		},
		{
			Model: gorm.Model{
				CreatedAt: time.Now().Add(-2 * time.Hour),
				UpdatedAt: time.Now().Add(-2 * time.Hour),
			},
			CO2:        505,
			Temp:       25.5,
			LocationID: 1,
		},
	}
	beforeInsert := []models.Co2Data{}
	f.Db.Find(&beforeInsert)
	result, err := db_calls.CreateCo2Data(f.Db, newData)
	afterInsert := []models.Co2Data{}
	f.Db.Find(&afterInsert)

	require.NoError(t, err)
	assert.Equal(t, 4, len(beforeInsert))
	assert.Equal(t, 2, len(result))
	assert.Equal(t, 6, len(afterInsert))

}

func TestCreateCo2Data_ShouldNotCreateValue(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	defer f.Teardown(t)

	dataToInsert := []models.Co2Data{}
	beforeInsert := []models.Co2Data{}
	f.Db.Find(&beforeInsert)
	errorMessage := "Empty list of co2 data to insert"
	result, err := db_calls.CreateCo2Data(f.Db, dataToInsert)

	assert.Equal(t, errorMessage, err.Error())
	assert.Equal(t, 0, len(beforeInsert))
	assert.Equal(t, 0, len(result))

}
