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
	req, writer := tests.SetupLocationRouter(f.Db, http.MethodGet, "/", "/", api.GetLocations)
	defer f.Teardown(t)

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, writer.Code, "HTTP request status code error")

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}

	actual := struct {
		Locations []models.Location `json:"location"`
	}{}
	if err := json.Unmarshal(body, &actual); err != nil {
		assert.Error(t, err)
	}

	expected := struct {
		Locations []models.Location `json:"location"`
	}{}
	assert.Equal(t, len(expected.Locations), len(actual.Locations))
	assert.Equal(t, len(actual.Locations), 0)
}

func TestGetLocation_ShouldReturnAllLocations(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	req, writer := tests.SetupLocationRouter(f.Db, http.MethodGet, "/", "/", api.GetLocations)
	defer f.Teardown(t)

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, writer.Code, "HTTP request status code error")

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}

	responseData := struct {
		Locations []models.Location `json:"location"`
	}{}

	if err := json.Unmarshal(body, &responseData); err != nil {
		assert.Error(t, err)
	}

	expected := tests.Locations

	assert.Equal(t, len(expected), len(responseData.Locations))
}

func TestGetLocationBySearch_ShouldReturnEmptyList(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	req, writer := tests.SetupLocationRouter(f.Db, http.MethodGet, "/search", "/search?id=999&name=not in database", api.GetLocationBySearch)
	defer f.Teardown(t)

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, writer.Code, "HTTP request status code error")

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}

	responseData := struct {
		Locations []models.Location `json:"location"`
	}{}

	if err := json.Unmarshal(body, &responseData); err != nil {
		assert.Error(t, err)
	}

	expected := []models.Location{}

	assert.Equal(t, len(expected), len(responseData.Locations))
	assert.Equal(t, len(responseData.Locations), 0)
}

func TestGetLocationBySearch_ShouldReturnLocationById(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	req, writer := tests.SetupLocationRouter(f.Db, http.MethodGet, "/search", "/search?id=2", api.GetLocationBySearch)
	defer f.Teardown(t)

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, writer.Code, "HTTP request status code error")

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}

	responseData := struct {
		Locations []models.Location `json:"location"`
	}{}

	if err := json.Unmarshal(body, &responseData); err != nil {
		assert.Error(t, err)
	}

	expected := []models.Location{tests.Locations[1]}

	assert.Equal(t, len(expected), len(responseData.Locations))
	assert.Equal(t, expected[0].Name, responseData.Locations[0].Name)
}

func TestGetLocationBySearch_ShouldReturnLocationByName(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	req, writer := tests.SetupLocationRouter(f.Db, http.MethodGet, "/search", "/search?name=test location 2", api.GetLocationBySearch)
	defer f.Teardown(t)

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, writer.Code, "HTTP request status code error")

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}

	responseData := struct {
		Locations []models.Location `json:"location"`
	}{}

	if err := json.Unmarshal(body, &responseData); err != nil {
		assert.Error(t, err)
	}

	expected := []models.Location{tests.Locations[1]}

	assert.Equal(t, len(expected), len(responseData.Locations))
	assert.Equal(t, expected[0].Name, responseData.Locations[0].Name)
}
func TestGetLocationBySearch_ShouldReturnLocationByIdAndName(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	req, writer := tests.SetupLocationRouter(f.Db, http.MethodGet, "/search", "/search?id=1&name=test location 2", api.GetLocationBySearch)
	defer f.Teardown(t)

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, writer.Code, "HTTP request status code error")

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}

	responseData := struct {
		Locations []models.Location `json:"location"`
	}{}

	if err := json.Unmarshal(body, &responseData); err != nil {
		assert.Error(t, err)
	}

	expected := []models.Location{tests.Locations[0], tests.Locations[1]}

	assert.Equal(t, len(expected), len(responseData.Locations))
	assert.Equal(t, expected[0].Name, responseData.Locations[0].Name)
	assert.Equal(t, expected[1].Name, responseData.Locations[1].Name)
}
