package tests

import (
	"testing"

	"github.com/fminister/co2monitor.api/db/db_calls"
	"github.com/fminister/co2monitor.api/models"
	"github.com/fminister/co2monitor.api/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestGetLocationBySearch(t *testing.T) {
	t.Run("search by id and find one", func(t *testing.T) {
		f := tests.BaseFixture{}
		f.Setup(t)
		f.AddDummyData(t)
		defer f.Teardown(t)

		result, err := db_calls.GetLocationBySearch(f.Db, "2", "")

		require.NoError(t, err)
		assert.Equal(t, tests.Locations[1].Name, result[0].Name)
		assert.Equal(t, 1, len(result))
	})

	t.Run("search by name and find one", func(t *testing.T) {
		f := tests.BaseFixture{}
		f.Setup(t)
		f.AddDummyData(t)
		defer f.Teardown(t)

		result, err := db_calls.GetLocationBySearch(f.Db, "", tests.Locations[1].Name)

		require.NoError(t, err)
		assert.Equal(t, tests.Locations[1].Name, result[0].Name)
		assert.Equal(t, 1, len(result))
	})

	t.Run("search by id and name and find both", func(t *testing.T) {
		f := tests.BaseFixture{}
		f.Setup(t)
		f.AddDummyData(t)
		defer f.Teardown(t)

		result, err := db_calls.GetLocationBySearch(f.Db, "1", tests.Locations[1].Name)

		require.NoError(t, err)
		assert.Equal(t, tests.Locations[0].Name, result[0].Name)
		assert.Equal(t, tests.Locations[1].Name, result[1].Name)
		assert.Equal(t, 2, len(result))
	})

	t.Run("search by id and name and find none", func(t *testing.T) {
		f := tests.BaseFixture{}
		f.Setup(t)
		f.AddDummyData(t)
		defer f.Teardown(t)

		result, err := db_calls.GetLocationBySearch(f.Db, "3", "not in db")

		require.NoError(t, err)
		assert.Equal(t, 0, len(result))
	})
}

func TestCreateLocation(t *testing.T) {
	t.Run("create single location", func(t *testing.T) {
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
	})

	t.Run("create multiple locations", func(t *testing.T) {
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
	})

	t.Run("empty location does not create new location", func(t *testing.T) {
		f := tests.BaseFixture{}
		f.Setup(t)
		defer f.Teardown(t)

		locationToInsert := []models.Location{}
		beforeInsert := []models.Location{}
		f.Db.Find(&beforeInsert)
		result, err := db_calls.CreateLocation(f.Db, locationToInsert)

		require.NoError(t, err)
		assert.Equal(t, 0, len(beforeInsert))
		assert.Equal(t, len(locationToInsert), len(result))
	})
}