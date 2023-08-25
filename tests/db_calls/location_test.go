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

	expectedSQL := "SELECT (.+) FROM \"locations\""
	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, expected[0].Name).AddRow(2, expected[1].Name)
	mock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	result, err := db_calls.GetLocation(db)

	assert.NoError(t, err)
	assert.Equal(t, len(expected), len(result))
	assert.Equal(t, expected[0].Name, result[0].Name)
	assert.Equal(t, expected[1].Name, result[1].Name)
}

func TestDbGetLocationById(t *testing.T) {
	sqlDB, db, mock := tests.DbMock(t)
	defer sqlDB.Close()

	expected := models.Location{
		Name: "test location",
	}

	expectedSQL := "SELECT (.+) FROM \"locations\" WHERE \"locations\".\"id\" = (.+) AND \"locations\".\"deleted_at\" IS NULL ORDER BY \"locations\".\"id\" LIMIT 1"
	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, expected.Name)
	mock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	result, err := db_calls.GetLocationById(db, "1")

	assert.NoError(t, err)
	assert.Equal(t, expected.Name, result.Name)
}
