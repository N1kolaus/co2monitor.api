package controllers

import (
	"net/http"

	"github.com/charmbracelet/log"

	"github.com/fminister/co2monitor.api/db/db_calls"
	ex "github.com/fminister/co2monitor.api/extensions"
	"github.com/fminister/co2monitor.api/models"
	"github.com/gin-gonic/gin"
)

// type APIEnv struct {
// 	DB *gorm.DB
// }

func (a *APIEnv) GetLocations(c *gin.Context) {
	locations, err := db_calls.GetLocation(a.DB)
	if err != nil {
		log.Errorf(`Could not find any locations. Error: "%s"`, err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Could not find any locations.",
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
		log.Errorf(`Could not find any locations by id or name. id: "%s"; name: "%s"; Error: "%s"`, id, name, err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Could not find any locations.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"locations": locations,
	})
}

func (a *APIEnv) CreateLocation(c *gin.Context) {
	var locations []models.Location
	if err := c.ShouldBindJSON(&locations); err != nil {
		log.Errorf(`Could not parse location from body. Error: "%s"`, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not parse location from body.",
		})
		return
	}

	if err := ex.Validator([]models.Location{}).Validate(locations); err != nil {
		log.Errorf(`Could parse location from body. Error: "%s"`, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	locations, err := db_calls.CreateLocation(a.DB, locations)
	if err != nil {
		log.Errorf(`Could not create location in db. Locations: "%#v" Error: "%s"`, locations, err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Could not create location.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"locations": locations,
	})
}

func (a *APIEnv) UpdateLocation(c *gin.Context) {
	locationId := c.Param("id")

	if _, err := db_calls.GetLocationById(a.DB, locationId); err != nil {
		log.Errorf(`Could not find location by id. id: "%s"; Error: "%s"`, locationId, err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Could not find location by id.",
		})
		return
	}

	var location models.Location
	if err := c.ShouldBindJSON(&location); err != nil {
		log.Errorf(`Could not parse location details from body. Error: "%s"`, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not parse location details from body.",
		})
		return
	}

	if err := ex.Validator(models.Location{}).Validate(location); err != nil {
		log.Errorf(`Could parse location from body. Error: "%s"`, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	location, err := db_calls.UpdateLocation(a.DB, location)
	if err != nil {
		log.Errorf(`Could not update location in db. Location: "%#v" Error: "%s"`, location, err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Could not update location.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"location": location,
	})
}
