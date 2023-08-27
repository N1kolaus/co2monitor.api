package tests

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupLocationRouter(db *gorm.DB, method, route string, handler gin.HandlerFunc) (*http.Request, *httptest.ResponseRecorder) {
	router := gin.New()

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

	req, err := http.NewRequest(method, route, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, req)

	return req, writer
}
