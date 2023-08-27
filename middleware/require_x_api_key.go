package middleware

import (
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func RequireApiKey(c *gin.Context) {
	APIKey := c.Request.Header.Get("X-API-KEY")
	adminAPIKey := os.Getenv("X_API_KEY_ADMIN")
	normalAPIKey := os.Getenv("X_API_KEY")

	if APIKey == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		log.Infof(`API-Key was not found or wrong. API-Key: "%s"; URL: "%s"`, APIKey, c.Request.URL)
		return
	}

	switch c.Request.Method {
	case http.MethodGet:
		if APIKey == adminAPIKey || APIKey == normalAPIKey {
			c.Next()
			return
		}
	case http.MethodPost, http.MethodPatch, http.MethodDelete:
		if APIKey == adminAPIKey {
			c.Next()
			return
		}
	}

	log.Infof(`Unauthorized API-Key. API-Key: "%s"; URL: "%s"; Method: "%s"; Path: "%s"`, APIKey, c.Request.URL, c.Request.Method, c.FullPath())
	c.AbortWithStatus(http.StatusUnauthorized)
}
