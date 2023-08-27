package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/fminister/co2monitor.api/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupLocationRouter(db *gorm.DB, method, route string, requestRoute string, handler gin.HandlerFunc, requestBody []byte) (*http.Request, *httptest.ResponseRecorder) {
	router := gin.Default()

	switch method {
	case http.MethodGet:
		router.GET(route, handler)
	case http.MethodPost:
		router.POST(route, handler)
	case http.MethodPatch:
		router.PATCH(route, handler)
	case http.MethodDelete:
		router.DELETE(route, handler)
	default:
		panic("Unsupported HTTP method")
	}

	writer := httptest.NewRecorder()
	req, err := http.NewRequest(method, requestRoute, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}
	router.ServeHTTP(writer, req)

	return req, writer
}

func SetupMiddlewareRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", middleware.RequireApiKey, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Authorized through middleware"})
	})
	router.POST("/", middleware.RequireApiKey, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Authorized through middleware"})
	})
	router.PATCH("/", middleware.RequireApiKey, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Authorized through middleware"})
	})
	router.DELETE("/", middleware.RequireApiKey, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Authorized through middleware"})
	})

	return router
}
