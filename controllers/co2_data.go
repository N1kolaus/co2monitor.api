package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCo2Data(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "co2 data",
	})
}
