package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/fminister/co2monitor.api/tests"

	"github.com/stretchr/testify/assert"
)

func TestRequireApiKey_UnauthorizedWithoutApiKey(t *testing.T) {
	router := tests.SetupMiddlewareRouter()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestRequireApiKey_AuthorizedWithNormalApiKey(t *testing.T) {
	router := tests.SetupMiddlewareRouter()
	normalAPIKey := "YOUR_NORMAL_API_KEY"
	os.Setenv("X_API_KEY", normalAPIKey)

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-API-KEY", normalAPIKey)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRequireApiKey_UnauthorizedWithWrongApiKey(t *testing.T) {
	router := tests.SetupMiddlewareRouter()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-API-KEY", "WRONG_API_KEY")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestRequireApiKey_AuthorizedWithAdminApiKeyForPostMethod(t *testing.T) {
	router := tests.SetupMiddlewareRouter()
	adminAPIKey := "YOUR_ADMIN_API_KEY"
	os.Setenv("X_API_KEY_ADMIN", adminAPIKey)

	req, err := http.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set("X-API-KEY", adminAPIKey)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRequireApiKey_UnauthorizedWithNormalApiKeyForPostMethod(t *testing.T) {
	router := tests.SetupMiddlewareRouter()
	normalAPIKey := "YOUR_NORMAL_API_KEY"
	os.Setenv("X_API_KEY", normalAPIKey)

	req, err := http.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set("X-API-KEY", normalAPIKey)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestRequireApiKey_AuthorizedWithAdminApiKeyForPatchMethod(t *testing.T) {
	router := tests.SetupMiddlewareRouter()
	adminAPIKey := "YOUR_ADMIN_API_KEY"
	os.Setenv("X_API_KEY_ADMIN", adminAPIKey)

	req, err := http.NewRequest(http.MethodPatch, "/", nil)
	req.Header.Set("X-API-KEY", adminAPIKey)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRequireApiKey_UnauthorizedWithNormalApiKeyForPatchMethod(t *testing.T) {
	router := tests.SetupMiddlewareRouter()
	normalAPIKey := "YOUR_NORMAL_API_KEY"
	os.Setenv("X_API_KEY", normalAPIKey)

	req, err := http.NewRequest(http.MethodPatch, "/", nil)
	req.Header.Set("X-API-KEY", normalAPIKey)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestRequireApiKey_AuthorizedWithAdminApiKeyForDeleteMethod(t *testing.T) {
	router := tests.SetupMiddlewareRouter()
	adminAPIKey := "YOUR_ADMIN_API_KEY"
	os.Setenv("X_API_KEY_ADMIN", adminAPIKey)

	req, err := http.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set("X-API-KEY", adminAPIKey)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRequireApiKey_UnauthorizedWithNormalApiKeyForDeleteMethod(t *testing.T) {
	router := tests.SetupMiddlewareRouter()
	normalAPIKey := "YOUR_NORMAL_API_KEY"
	os.Setenv("X_API_KEY", normalAPIKey)

	req, err := http.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set("X-API-KEY", normalAPIKey)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
