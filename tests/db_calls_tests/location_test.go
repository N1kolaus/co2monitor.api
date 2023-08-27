package tests

import (
	"testing"

	"github.com/fminister/co2monitor.api/db/db_calls"
	"github.com/fminister/co2monitor.api/models"
	"github.com/fminister/co2monitor.api/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestGetLocation_ShouldReturnAllLocations(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	defer f.Teardown(t)

	result, err := db_calls.GetLocation(f.Db)

	require.NoError(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, tests.Locations[0].Name, result[0].Name)
	assert.Equal(t, tests.Locations[1].Name, result[1].Name)
}

func TestGetLocationBySearch_ShouldFindOneById(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	defer f.Teardown(t)

	result, err := db_calls.GetLocationBySearch(f.Db, "2", "")

	require.NoError(t, err)
	assert.Equal(t, tests.Locations[1].Name, result[0].Name)
	assert.Equal(t, 1, len(result))
}

func TestGetLocationBySearch_ShouldFindOneByName(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	defer f.Teardown(t)

	result, err := db_calls.GetLocationBySearch(f.Db, "", tests.Locations[1].Name)

	require.NoError(t, err)
	assert.Equal(t, tests.Locations[1].Name, result[0].Name)
	assert.Equal(t, 1, len(result))
}

func TestGetLocationBySearch_ShouldReturnBothByIdAndName(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	defer f.Teardown(t)

	result, err := db_calls.GetLocationBySearch(f.Db, "1", tests.Locations[1].Name)

	require.NoError(t, err)
	assert.Equal(t, tests.Locations[0].Name, result[0].Name)
	assert.Equal(t, tests.Locations[1].Name, result[1].Name)
	assert.Equal(t, 2, len(result))
}

func TestGetLocationBySearch_ShouldReturnNoneByIdAndName(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	defer f.Teardown(t)

	result, err := db_calls.GetLocationBySearch(f.Db, "3", "not in db")

	require.NoError(t, err)
	assert.Equal(t, 0, len(result))
}

func TestCreateLocation_ShouldCreateSingleLocation(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	defer f.Teardown(t)

	locationToInsert := []models.Location{tests.Locations[1]}
	beforeInsert := []models.Location{}
	f.Db.Find(&beforeInsert)
	result, err := db_calls.CreateLocation(f.Db, locationToInsert)

	require.NoError(t, err)
	assert.Equal(t, 0, len(beforeInsert))
	assert.Equal(t, len(locationToInsert), len(result))
	assert.Equal(t, tests.Locations[1].Name, result[0].Name)
}

func TestCreateLocation_ShouldCreateMultipleLocations(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	defer f.Teardown(t)

	locationToInsert := []models.Location{tests.Locations[0], tests.Locations[1]}
	beforeInsert := []models.Location{}
	f.Db.Find(&beforeInsert)
	result, err := db_calls.CreateLocation(f.Db, locationToInsert)

	require.NoError(t, err)
	assert.Equal(t, 0, len(beforeInsert))
	assert.Equal(t, len(locationToInsert), len(result))
	assert.Equal(t, tests.Locations[0].Name, result[0].Name)
	assert.Equal(t, tests.Locations[1].Name, result[1].Name)
}

func TestCreateLocation_ShouldNotCreateLocation(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	defer f.Teardown(t)

	locationToInsert := []models.Location{}
	beforeInsert := []models.Location{}
	errorMessage := "Empty list of locations to insert"
	f.Db.Find(&beforeInsert)
	result, err := db_calls.CreateLocation(f.Db, locationToInsert)

	assert.Equal(t, errorMessage, err.Error())
	assert.Equal(t, 0, len(beforeInsert))
	assert.Equal(t, len(locationToInsert), len(result))
}

func TestUpdateLocation_ShouldUpdateLocation(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	defer f.Teardown(t)

	updatedLocation := models.Location{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "updated location",
	}

	currentLocation := models.Location{}
	f.Db.First(&currentLocation, updatedLocation.ID)

	result, err := db_calls.UpdateLocation(f.Db, updatedLocation)

	require.NoError(t, err)
	assert.Equal(t, currentLocation.Name, tests.Locations[0].Name)
	assert.Equal(t, updatedLocation.Name, result.Name)
	assert.NotEqual(t, currentLocation.Name, result.Name)
	assert.NotEqual(t, currentLocation.UpdatedAt, result.UpdatedAt)
	assert.Greater(t, result.UpdatedAt.Unix(), tests.Locations[0].UpdatedAt.Unix())
}

func TestDeleteLocation_ShouldDeleteLocation(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	defer f.Teardown(t)

	locationToDelete := models.Location{
		Model: gorm.Model{
			ID: 1,
		},
	}

	oldLocations := []models.Location{}
	f.Db.Find(&oldLocations)

	err := db_calls.DeleteLocation(f.Db, locationToDelete)

	newLocations := []models.Location{}
	f.Db.Find(&newLocations)

	require.NoError(t, err)
	assert.NotEqual(t, len(oldLocations), len(newLocations))
	assert.Equal(t, len(oldLocations), len(newLocations)+1)
}
