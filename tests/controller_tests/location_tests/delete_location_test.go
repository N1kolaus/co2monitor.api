package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/fminister/co2monitor.api/controllers"
	"github.com/fminister/co2monitor.api/models"
	"github.com/fminister/co2monitor.api/tests"
	"github.com/stretchr/testify/assert"
)

func TestDeleteLocation_ShouldDeleteLocation(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	req, writer := tests.SetupRouter(f.Db, http.MethodDelete, "/:id", "/1", api.DeleteLocation, nil)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	responseData := ""
	if err := json.Unmarshal(body, &responseData); err != nil {
		assert.Error(t, err)
	}
	expectedInDb := []models.Location{}
	f.Db.Find(&expectedInDb)

	assert.Equal(t, http.MethodDelete, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusNoContent, writer.Code, "HTTP request status code error")
	assert.Equal(t, len(expectedInDb), 1)
	assert.Equal(t, expectedInDb[0].Name, tests.Locations[1].Name)
}

func TestDeleteLocation_ShouldReturnErrorUnknownId(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	req, writer := tests.SetupRouter(f.Db, http.MethodDelete, "/:id", "/99", api.DeleteLocation, nil)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	errorMessage := ""
	if err := json.Unmarshal(body, &errorMessage); err != nil {
		assert.Error(t, err)
	}
	expectedErrorMessage := "Could not find location by id."
	expectedInDb := []models.Location{}
	f.Db.Find(&expectedInDb)

	assert.Equal(t, http.MethodDelete, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusNotFound, writer.Code, "HTTP request status code error")
	assert.Equal(t, expectedErrorMessage, errorMessage)
	assert.Equal(t, len(expectedInDb), len(tests.Locations))
}
