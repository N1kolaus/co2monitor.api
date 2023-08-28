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

func TestCreateCo2Data_ShouldCreateSingleCo2DataValue(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	// add dummy data or creating newData will fail because of foreign key constraint
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	newCo2Data := []models.Co2Data{
		{
			LocationID: 1,
			CO2:        666,
			Temp:       11.1,
		},
	}
	req, writer := tests.SetupRouter(f.Db, http.MethodPost, "/new", "/new", api.CreateCo2Data, tests.CO2ToJSON(newCo2Data))
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	responseData := []models.Co2Data{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		assert.Error(t, err)
	}
	expectedInDb := []models.Co2Data{}
	f.Db.Find(&expectedInDb)

	assert.Equal(t, http.MethodPost, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusCreated, writer.Code, "HTTP request status code error")
	assert.Equal(t, len(newCo2Data), len(responseData))
	assert.Equal(t, len(expectedInDb), len(tests.CO2)+len(newCo2Data))
	assert.Equal(t, newCo2Data[0].LocationID, responseData[0].LocationID)
	assert.Equal(t, newCo2Data[0].CO2, responseData[0].CO2)
}

func TestCreateCo2Data_ShouldCreateMultipleCo2DataValues(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	// add dummy data or creating newData will fail because of foreign key constraint
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	newCo2Data := []models.Co2Data{
		{
			LocationID: 1,
			CO2:        666,
			Temp:       11.1,
		},
		{
			LocationID: 1,
			CO2:        777,
			Temp:       22.2,
		},
	}
	req, writer := tests.SetupRouter(f.Db, http.MethodPost, "/new", "/new", api.CreateCo2Data, tests.CO2ToJSON(newCo2Data))
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	responseData := []models.Co2Data{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		assert.Error(t, err)
	}
	expectedInDb := []models.Co2Data{}
	f.Db.Find(&expectedInDb)

	assert.Equal(t, http.MethodPost, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusCreated, writer.Code, "HTTP request status code error")
	assert.Equal(t, len(newCo2Data), len(responseData))
	assert.Equal(t, len(expectedInDb), len(tests.CO2)+len(newCo2Data))
	assert.Equal(t, newCo2Data[0].LocationID, responseData[0].LocationID)
	assert.Equal(t, newCo2Data[0].CO2, responseData[0].CO2)
	assert.Equal(t, newCo2Data[1].LocationID, responseData[1].LocationID)
	assert.Equal(t, newCo2Data[1].CO2, responseData[1].CO2)
}

func TestCreateCo2Data_ShouldReturnErrorMissingValuesInJSON(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	api := &controllers.APIEnv{DB: f.Db}
	newCo2Data := []models.Co2Data{
		{
			LocationID: 1,
			CO2:        666,
		},
		{
			Temp: 15,
			CO2:  777,
		},
		{
			LocationID: 1,
			Temp:       10,
		},
	}
	req, writer := tests.SetupRouter(f.Db, http.MethodPost, "/new", "/new", api.CreateCo2Data, tests.CO2ToJSON(newCo2Data))
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
			"temp": {
				"required",
			},
		},
		"1": {
			"location_id": {
				"required",
			},
		},
		"2": {
			"co2": {
				"required",
			},
		},
	}

	assert.Equal(t, http.MethodPost, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusBadRequest, writer.Code, "HTTP request status code error")
	assert.Equal(t, expectedErrorMessage, errorMessage)
}

func TestCreateCo2Data_ShouldReturnErrorCouldNotCreateDataDbError(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	api := &controllers.APIEnv{DB: f.Db}
	newCo2Data := []models.Co2Data{
		{
			LocationID: 99,
			CO2:        666,
			Temp:       11.1,
		},
	}
	req, writer := tests.SetupRouter(f.Db, http.MethodPost, "/new", "/new", api.CreateCo2Data, tests.CO2ToJSON(newCo2Data))
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	errorMessage := ""
	if err := json.Unmarshal(body, &errorMessage); err != nil {
		assert.Error(t, err)
	}
	expectedErrorMessage := "Could not create co2 data."

	assert.Equal(t, http.MethodPost, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusBadRequest, writer.Code, "HTTP request status code error")
	assert.Equal(t, expectedErrorMessage, errorMessage)
}

func TestCreateCo2Data_ShouldReturnErrorWrongBinding(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	api := &controllers.APIEnv{DB: f.Db}
	requestBody, _ := json.Marshal(models.Co2Data{
		LocationID: 99,
		CO2:        666,
		Temp:       11.1,
	})
	req, writer := tests.SetupRouter(f.Db, http.MethodPost, "/new", "/new", api.CreateCo2Data, requestBody)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	errorMessage := ""
	if err := json.Unmarshal(body, &errorMessage); err != nil {
		assert.Error(t, err)
	}
	expectedErrorMessage := "Could not parse co2 data from body."

	assert.Equal(t, http.MethodPost, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusBadRequest, writer.Code, "HTTP request status code error")
	assert.Equal(t, expectedErrorMessage, errorMessage)
}
