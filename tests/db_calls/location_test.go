package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fminister/co2monitor.api/db/db_calls"
	"github.com/fminister/co2monitor.api/models"
	"github.com/fminister/co2monitor.api/tests"
	"github.com/stretchr/testify/assert"
)

func TestDbGetLocation_ShouldGetListOfLocations(t *testing.T) {
	sqlDB, db, mock := tests.DbMock(t)
	defer sqlDB.Close()

	expected := []models.Location{
		{
			Name: "test location 1",
		},
		{
			Name: "test location 2",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, expected[0].Name).AddRow(2, expected[1].Name)
	mock.ExpectQuery(`SELECT (.+) FROM "locations"`).WillReturnRows(rows)

	result, err := db_calls.GetLocation(db)

	assert.NoError(t, err)
	assert.Equal(t, len(expected), len(result))
	assert.Equal(t, expected[0].Name, result[0].Name)
	assert.Equal(t, expected[1].Name, result[1].Name)
}

func TestDbGetLocationBySearch_ShouldFindById(t *testing.T) {
	sqlDB, db, mock := tests.DbMock(t)
	defer sqlDB.Close()

	expected := []models.Location{
		{
			Name: "test location 1",
		},
		{
			Name: "test location 2",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(2, expected[1].Name)
	mock.ExpectQuery(`SELECT (.+) FROM "locations"`).WillReturnRows(rows)

	result, err := db_calls.GetLocationBySearch(db, "2", "")

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, expected[1].Name, result[0].Name)
}

func TestDbGetLocationBySearch_ShouldFindByName(t *testing.T) {
	sqlDB, db, mock := tests.DbMock(t)
	defer sqlDB.Close()

	expected := []models.Location{
		{
			Name: "test location 1",
		},
		{
			Name: "test location 2",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(2, expected[1].Name)
	mock.ExpectQuery(`SELECT (.+) FROM "locations"`).WillReturnRows(rows)

	result, err := db_calls.GetLocationBySearch(db, "", expected[1].Name)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, expected[1].Name, result[0].Name)
}

func TestDbGetLocationBySearch_ShouldFindByIdAndName(t *testing.T) {
	sqlDB, db, mock := tests.DbMock(t)
	defer sqlDB.Close()

	expected := []models.Location{
		{
			Name: "test location 1",
		},
		{
			Name: "test location 2",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, expected[0].Name).AddRow(2, expected[1].Name)
	mock.ExpectQuery(`SELECT (.+) FROM "locations"`).WillReturnRows(rows)

	result, err := db_calls.GetLocationBySearch(db, "1", expected[1].Name)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, expected[0].Name, result[0].Name)
	assert.Equal(t, expected[1].Name, result[1].Name)
}
