package controllers

import (
	"net/http"

	"github.com/fminister/co2monitor.api/db/db_calls"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type APIEnv struct {
	DB *gorm.DB
}

func (a *APIEnv) GetLocations(c *gin.Context) {
	locations, err := db_calls.GetLocation(a.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"location": locations,
	})
}

func (a *APIEnv) GetLocationBySearch(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")

	locations, err := db_calls.GetLocationBySearch(a.DB, id, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"locations": locations,
	})
}
