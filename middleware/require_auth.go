package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	APIKey := c.Request.Header.Get("X-API-KEY")

	if APIKey == "" || APIKey != os.Getenv("X_API_KEY") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		log.Printf("APIKey was empty or invalid: [%s]", APIKey)
		return
	}

	c.Next()
}
