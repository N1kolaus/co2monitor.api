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

func TestGetLocationBySearch_ShouldReturnEmptyList(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	req, writer := tests.SetupLocationRouter(f.Db, http.MethodGet, "/search", "/search?id=999&name=not in database", api.GetLocationBySearch, nil)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	responseData := []models.Location{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		assert.Error(t, err)
	}
	expected := []models.Location{}

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, writer.Code, "HTTP request status code error")
	assert.Equal(t, len(expected), len(responseData))
	assert.Equal(t, len(responseData), 0)
}

func TestGetLocationBySearch_ShouldReturnLocationById(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	req, writer := tests.SetupLocationRouter(f.Db, http.MethodGet, "/search", "/search?id=2", api.GetLocationBySearch, nil)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	responseData := []models.Location{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		assert.Error(t, err)
	}
	expected := []models.Location{tests.Locations[1]}

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, writer.Code, "HTTP request status code error")
	assert.Equal(t, len(expected), len(responseData))
	assert.Equal(t, expected[0].Name, responseData[0].Name)
}

func TestGetLocationBySearch_ShouldReturnLocationByName(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	req, writer := tests.SetupLocationRouter(f.Db, http.MethodGet, "/search", "/search?name=test location 2", api.GetLocationBySearch, nil)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	responseData := []models.Location{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		assert.Error(t, err)
	}
	expected := []models.Location{tests.Locations[1]}

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, writer.Code, "HTTP request status code error")
	assert.Equal(t, len(expected), len(responseData))
	assert.Equal(t, expected[0].Name, responseData[0].Name)
}
func TestGetLocationBySearch_ShouldReturnLocationByIdAndName(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	req, writer := tests.SetupLocationRouter(f.Db, http.MethodGet, "/search", "/search?id=1&name=test location 2", api.GetLocationBySearch, nil)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	responseData := []models.Location{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		assert.Error(t, err)
	}
	expected := []models.Location{tests.Locations[0], tests.Locations[1]}

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, writer.Code, "HTTP request status code error")
	assert.Equal(t, len(expected), len(responseData))
	assert.Equal(t, expected[0].Name, responseData[0].Name)
	assert.Equal(t, expected[1].Name, responseData[1].Name)
}
