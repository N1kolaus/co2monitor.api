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
	"gorm.io/gorm"
)

func TestCreateLocation_ShouldCreateSingleLocation(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	api := &controllers.APIEnv{DB: f.Db}
	req, writer := tests.SetupRouter(f.Db, http.MethodPost, "/new", "/new", api.CreateLocation, tests.LocationsToJSON([]models.Location{tests.Locations[0]}))
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	responseData := []models.Location{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		assert.Error(t, err)
	}
	expected := []models.Location{tests.Locations[0]}
	expectedInDb := []models.Location{}
	f.Db.Find(&expectedInDb)

	assert.Equal(t, http.MethodPost, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusCreated, writer.Code, "HTTP request status code error")
	assert.Equal(t, len(expected), len(responseData))
	assert.Equal(t, len(expectedInDb), len(expected))
	assert.Equal(t, expected[0].Name, responseData[0].Name)
	assert.Equal(t, expectedInDb[0].Name, expected[0].Name)
}

func TestCreateLocation_ShouldCreateMultipleLocations(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	api := &controllers.APIEnv{DB: f.Db}
	req, writer := tests.SetupRouter(f.Db, http.MethodPost, "/new", "/new", api.CreateLocation, tests.LocationsToJSON([]models.Location{tests.Locations[0], tests.Locations[1]}))
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
	expectedInDb := []models.Location{}
	f.Db.Find(&expectedInDb)

	assert.Equal(t, http.MethodPost, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusCreated, writer.Code, "HTTP request status code error")
	assert.Equal(t, len(expected), len(responseData))
	assert.Equal(t, len(expectedInDb), len(expected))
	assert.Equal(t, expected[0].Name, responseData[0].Name)
	assert.Equal(t, expectedInDb[0].Name, expected[0].Name)
	assert.Equal(t, expected[1].Name, responseData[1].Name)
	assert.Equal(t, expectedInDb[1].Name, expected[1].Name)
}

func TestCreateLocation_ShouldReturnErrorMissingNameInJSON(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	api := &controllers.APIEnv{DB: f.Db}
	requestBody, _ := json.Marshal([]models.Location{
		{
			Name: "",
		},
	})
	req, writer := tests.SetupRouter(f.Db, http.MethodPost, "/new", "/new", api.CreateLocation, requestBody)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	errorMessage := map[string]map[string][]string{}
	if err := json.Unmarshal(body, &errorMessage); err != nil {
		assert.Error(t, err)
	}
	expectedErrorMessage := map[string]map[string][]string{
		"0": {
			"name": {
				"required",
				"name's length must be higher equal to 3",
			},
		},
	}

	assert.Equal(t, http.MethodPost, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusBadRequest, writer.Code, "HTTP request status code error")
	assert.Equal(t, expectedErrorMessage, errorMessage)
}

func TestCreateLocation_ShouldReturnErrorNameToShortInJSON(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	api := &controllers.APIEnv{DB: f.Db}
	requestBody, _ := json.Marshal([]models.Location{
		{
			Name: "a",
		},
	})
	req, writer := tests.SetupRouter(f.Db, http.MethodPost, "/new", "/new", api.CreateLocation, requestBody)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	errorMessage := map[string]map[string][]string{}
	if err := json.Unmarshal(body, &errorMessage); err != nil {
		assert.Error(t, err)
	}
	expectedErrorMessage := map[string]map[string][]string{
		"0": {
			"name": {
				"name's length must be higher equal to 3",
			},
		},
	}

	assert.Equal(t, http.MethodPost, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusBadRequest, writer.Code, "HTTP request status code error")
	assert.Equal(t, expectedErrorMessage, errorMessage)
}

func TestCreateLocation_ShouldReturnErrorNameAlreadyExistsInDb(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	req, writer := tests.SetupRouter(f.Db, http.MethodPost, "/new", "/new", api.CreateLocation, tests.LocationsToJSON([]models.Location{tests.Locations[0]}))
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	errorMessage := ""
	if err := json.Unmarshal(body, &errorMessage); err != nil {
		assert.Error(t, err)
	}
	expectedErrorMessage := "Could not create location. Name already exists."

	assert.Equal(t, http.MethodPost, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusBadRequest, writer.Code, "HTTP request status code error")
	assert.Equal(t, expectedErrorMessage, errorMessage)
}

func TestCreateLocation_ShouldReturnErrorWrongBinding(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	api := &controllers.APIEnv{DB: f.Db}
	requestBody, _ := json.Marshal(models.Location{
		Model: gorm.Model{
			ID: 99,
		},
		Name: "updated location",
	})
	req, writer := tests.SetupRouter(f.Db, http.MethodPost, "/new", "/new", api.CreateLocation, requestBody)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	errorMessage := ""
	if err := json.Unmarshal(body, &errorMessage); err != nil {
		assert.Error(t, err)
	}
	expectedErrorMessage := "Could not parse location from body."

	assert.Equal(t, http.MethodPost, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusBadRequest, writer.Code, "HTTP request status code error")
	assert.Equal(t, expectedErrorMessage, errorMessage)
}
