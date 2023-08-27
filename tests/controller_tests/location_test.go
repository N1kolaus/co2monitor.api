package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/fminister/co2monitor.api/models"
	"github.com/fminister/co2monitor.api/tests"
	"github.com/stretchr/testify/assert"
)

func TestGetLocation(t *testing.T) {
	t.Run("should return empty list", func(t *testing.T) {
		f := tests.BaseFixture{}
		f.Setup(t)
		req, writer := tests.SetupLocationRouter(f.Db)
		defer f.Teardown(t)

		assert.Equal(t, http.MethodGet, req.Method, "HTTP request method error")
		assert.Equal(t, http.StatusOK, writer.Code, "HTTP request status code error")

		body, err := io.ReadAll(writer.Body)
		if err != nil {
			assert.Error(t, err)
		}

		actual := models.Location{}
		if err := json.Unmarshal(body, &actual); err != nil {
			assert.Error(t, err)
		}

		expected := models.Location{}
		assert.Equal(t, expected, actual)
	})

	t.Run("should return all locations", func(t *testing.T) {
		f := tests.BaseFixture{}
		f.Setup(t)
		f.AddDummyData(t)
		req, writer := tests.SetupLocationRouter(f.Db)
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
	})
}
