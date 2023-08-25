package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RequireAuth(expectedAPIKeyEnvName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		expectedAPIKey := os.Getenv(expectedAPIKeyEnvName)
		APIKey := c.Request.Header.Get("X-API-KEY")

		if APIKey == "" || APIKey != expectedAPIKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			log.Printf("APIKey was empty or invalid: [%s]", APIKey)
			return
		}

		c.Next()
	}
}
