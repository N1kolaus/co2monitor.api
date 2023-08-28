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

func TestGetLocation_ShouldReturnEmptyList(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	api := &controllers.APIEnv{DB: f.Db}
	req, writer := tests.SetupRouter(f.Db, http.MethodGet, "/", "/", api.GetLocations, nil)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	actual := []models.Location{}
	if err := json.Unmarshal(body, &actual); err != nil {
		assert.Error(t, err)
	}
	expected := []models.Location{}

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, writer.Code, "HTTP request status code error")
	assert.Equal(t, len(expected), len(actual))
	assert.Equal(t, len(actual), 0)
}

func TestGetLocation_ShouldReturnAllLocations(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	req, writer := tests.SetupRouter(f.Db, http.MethodGet, "/", "/", api.GetLocations, nil)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	responseData := []models.Location{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		assert.Error(t, err)
	}
	expected := tests.Locations

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, writer.Code, "HTTP request status code error")
	assert.Equal(t, len(expected), len(responseData))
}
