package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	APIKey := c.Request.Header.Get("X-API-KEY")

	if APIKey == "" || APIKey != os.Getenv("X_API_KEY") {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}
