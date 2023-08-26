package middleware

import (
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func RequireAuth(expectedAPIKeyEnvName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		expectedAPIKey := os.Getenv(expectedAPIKeyEnvName)
		APIKey := c.Request.Header.Get("X-API-KEY")

		if APIKey == "" || APIKey != expectedAPIKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			log.Infof(`API-Key was not found or not valid. API-Key: "%s"; URL: "%s"`, APIKey, c.Request.URL)
			return
		}

		c.Next()
	}
}
