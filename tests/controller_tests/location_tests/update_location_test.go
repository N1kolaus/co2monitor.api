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

func TestUpdateLocation_ShouldUpdateLocation(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	updatedLocation := models.Location{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "updated location",
	}
	requestBody, _ := json.Marshal(updatedLocation)
	req, writer := tests.SetupRouter(f.Db, http.MethodPatch, "/:id", "/1", api.UpdateLocation, requestBody)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	responseData := models.Location{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		assert.Error(t, err)
	}
	expectedInDb := []models.Location{}
	f.Db.Find(&expectedInDb)

	assert.Equal(t, http.MethodPatch, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, writer.Code, "HTTP request status code error")
	assert.Equal(t, responseData.Name, expectedInDb[0].Name)
	assert.Equal(t, responseData.Name, updatedLocation.Name)
	assert.Equal(t, responseData.ID, updatedLocation.ID)
	assert.NotEqual(t, responseData.Name, expectedInDb[1].Name)
}

func TestUpdateLocation_ShouldReturnErrorUnknownId(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	requestBody, _ := json.Marshal(models.Location{
		Model: gorm.Model{
			ID: 99,
		},
		Name: "updated location",
	})
	req, writer := tests.SetupRouter(f.Db, http.MethodPatch, "/:id", "/99", api.UpdateLocation, requestBody)
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

	assert.Equal(t, http.MethodPatch, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusNotFound, writer.Code, "HTTP request status code error")
	assert.Equal(t, expectedErrorMessage, errorMessage)
}

func TestUpdateLocation_ShouldReturnErrorMissingNameInJSON(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	requestBody, _ := json.Marshal(models.Location{
		Name: "",
	})
	req, writer := tests.SetupRouter(f.Db, http.MethodPatch, "/:id", "/1", api.UpdateLocation, requestBody)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	errorMessage := map[string][]string{}
	if err := json.Unmarshal(body, &errorMessage); err != nil {
		assert.Error(t, err)
	}
	expectedErrorMessage := map[string][]string{
		"name": {
			"required",
			"name's length must be higher equal to 3",
		},
	}

	assert.Equal(t, http.MethodPatch, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusBadRequest, writer.Code, "HTTP request status code error")
	assert.Equal(t, expectedErrorMessage, errorMessage)
}

func TestUpdateLocation_ShouldReturnErrorNameToShortInJSON(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	requestBody, _ := json.Marshal(models.Location{
		Name: "a",
	})
	req, writer := tests.SetupRouter(f.Db, http.MethodPatch, "/:id", "/1", api.UpdateLocation, requestBody)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	errorMessage := map[string][]string{}
	if err := json.Unmarshal(body, &errorMessage); err != nil {
		assert.Error(t, err)
	}
	expectedErrorMessage := map[string][]string{
		"name": {
			"name's length must be higher equal to 3",
		},
	}

	assert.Equal(t, http.MethodPatch, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusBadRequest, writer.Code, "HTTP request status code error")
	assert.Equal(t, expectedErrorMessage, errorMessage)
}

func TestUpdateLocation_ShouldReturnErrorWrongBinding(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	requestBody, _ := json.Marshal([]models.Location{
		{
			Name: "a",
		},
	})
	req, writer := tests.SetupRouter(f.Db, http.MethodPatch, "/:id", "/1", api.UpdateLocation, requestBody)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	errorMessage := ""
	if err := json.Unmarshal(body, &errorMessage); err != nil {
		assert.Error(t, err)
	}
	expectedErrorMessage := "Could not parse location details from body."

	assert.Equal(t, http.MethodPatch, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusBadRequest, writer.Code, "HTTP request status code error")
	assert.Equal(t, expectedErrorMessage, errorMessage)
}
