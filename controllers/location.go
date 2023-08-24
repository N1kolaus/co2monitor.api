package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLocation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"location": "location",
	})
}
