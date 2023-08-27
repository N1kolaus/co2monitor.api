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

	if (c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPatch || c.Request.Method == http.MethodDelete) && APIKey == adminAPIKey {
		c.Next()
		return
	} else if APIKey == normalAPIKey {
		c.Next()
		return
	}

	log.Infof(`Unauthorized API-Key. API-Key: "%s"; URL: "%s"; Method: "%s"; Path: "%s"`, APIKey, c.Request.URL, c.Request.Method, c.FullPath())
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
}
