package tests

import (
	"net/http"
	"net/http/httptest"

	"github.com/fminister/co2monitor.api/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupLocationRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	router := gin.New()
	api := &controllers.APIEnv{DB: db}
	router.GET("/", api.GetLocations)

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, req)

	return req, writer
}
