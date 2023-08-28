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

func TestGetLatestCo2Data_ShouldReturnSingleCo2Data(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	req, writer := tests.SetupRouter(f.Db, http.MethodGet, "/:id/latest", "/1/latest", api.GetLatestCo2Data, nil)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	responseData := models.Co2Data{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusOK, writer.Code, "HTTP request status code error")
	assert.Equal(t, tests.CO2[0].CO2, responseData.CO2)
	assert.Equal(t, tests.CO2[0].Temp, responseData.Temp)
	assert.Equal(t, tests.CO2[0].LocationID, responseData.LocationID)
}

func TestGetLatestCo2Data_ShouldReturnErrorLocationIdUnknown(t *testing.T) {
	f := tests.BaseFixture{}
	f.Setup(t)
	f.AddDummyData(t)
	api := &controllers.APIEnv{DB: f.Db}
	locationId := "99"
	req, writer := tests.SetupRouter(f.Db, http.MethodGet, "/:id/latest", fmt.Sprintf(`/%s/latest`, locationId), api.GetLatestCo2Data, nil)
	defer f.Teardown(t)

	body, err := io.ReadAll(writer.Body)
	if err != nil {
		assert.Error(t, err)
	}
	errorMessage := ""
	if err := json.Unmarshal(body, &errorMessage); err != nil {
		assert.Error(t, err)
	}
	expectedErrorMessage := fmt.Sprintf(`Could not find any co2 data with this locationId: <%s>.`, locationId)

	assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
	assert.Equal(t, http.StatusNotFound, writer.Code, "HTTP request status code error")
	assert.Equal(t, expectedErrorMessage, errorMessage)
}
