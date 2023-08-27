package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *APIEnv) GetCo2DataBySearch(c *gin.Context) {

	// duration, _ := time.ParseDuration(hours)

	c.JSON(http.StatusOK, gin.H{
		"data": "co2 data",
	})
}

func (a *APIEnv) GetLatestCo2Data(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "co2 data",
	})
}

func (a *APIEnv) CreateCo2Data(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "co2 data",
	})
}
