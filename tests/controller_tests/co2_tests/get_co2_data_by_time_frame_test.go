package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/fminister/co2monitor.api/controllers"
	"github.com/fminister/co2monitor.api/models"
	"github.com/fminister/co2monitor.api/tests"
	"github.com/stretchr/testify/assert"
)

func TestGetCo2DataByTimeFrame_ShouldReturnListOfCo2Data(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	locationId := "1"
	searchQuery := "?period=24h"
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
	f.Db.Create(&newCo2Data)
	req, writer := tests.SetupRouter(f.Db, http.MethodGet, "/:id/search", fmt.Sprintf(`/%s/search%s`, locationId, searchQuery), api.GetCo2DataByTimeFrame, nil)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	responseData := []models.Co2Data{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, writer.Code, "HTTP request status code error")
	assert.Equal(t, len(newCo2Data), len(responseData))
	assert.Equal(t, newCo2Data[0].LocationID, responseData[0].LocationID)
	assert.Equal(t, newCo2Data[0].CO2, responseData[0].CO2)
	assert.Equal(t, newCo2Data[0].Temp, responseData[0].Temp)
	assert.Equal(t, newCo2Data[1].LocationID, responseData[1].LocationID)
	assert.Equal(t, newCo2Data[1].CO2, responseData[1].CO2)
	assert.Equal(t, newCo2Data[1].Temp, responseData[1].Temp)
}

func TestGetCo2DataByTimeFrame_ShouldReturnEmptyList(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	locationId := "1"
	searchQuery := "?period=1m"
	req, writer := tests.SetupRouter(f.Db, http.MethodGet, "/:id/search", fmt.Sprintf(`/%s/search%s`, locationId, searchQuery), api.GetCo2DataByTimeFrame, nil)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	responseData := []models.Co2Data{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, writer.Code, "HTTP request status code error")
	assert.Equal(t, 0, len(responseData))
}

func TestGetCo2DataByTimeFrame_ShouldReturnDefaultTimeFrameCo2DataList(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	locationId := "1"
	searchQuery := "?period=asdf"
	newCo2Data := []models.Co2Data{
		{
			LocationID: 1,
			CO2:        666,
			Temp:       11.1,
		},
	}
	f.Db.Create(&newCo2Data)
	req, writer := tests.SetupRouter(f.Db, http.MethodGet, "/:id/search", fmt.Sprintf(`/%s/search%s`, locationId, searchQuery), api.GetCo2DataByTimeFrame, nil)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	responseData := []models.Co2Data{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, writer.Code, "HTTP request status code error")
	assert.Equal(t, len(newCo2Data), len(responseData))
	assert.Equal(t, newCo2Data[0].LocationID, responseData[0].LocationID)
	assert.Equal(t, newCo2Data[0].CO2, responseData[0].CO2)
	assert.Equal(t, newCo2Data[0].Temp, responseData[0].Temp)
}

func TestGetCo2DataByTimeFrame_ShouldReturnErrorLocationIdUnknown(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	locationId := "99"
	searchQuery := "?period=24h"
	req, writer := tests.SetupRouter(f.Db, http.MethodGet, "/:id/search", fmt.Sprintf(`/%s/search%s`, locationId, searchQuery), api.GetCo2DataByTimeFrame, nil)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	errorMessage := ""
	if err := json.Unmarshal(body, &errorMessage); err != nil {
		assert.Error(t, err)
	}
	expectedErrorMessage := fmt.Sprintf(`Could not find any location with this id: "%s".`, locationId)

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusNotFound, writer.Code, "HTTP request status code error")
	assert.Equal(t, expectedErrorMessage, errorMessage)
}
