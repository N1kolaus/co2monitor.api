package controllers

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/dranikpg/dto-mapper"

	"github.com/fminister/co2monitor.api/db/db_calls"
	ex "github.com/fminister/co2monitor.api/extensions"
	"github.com/fminister/co2monitor.api/models"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// GetLocations godoc
//
//	@Summary		Get all locations
//	@Description	Get all locations.
//	@Tags			Locations
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	[]models.LocationDto
//	@Failure		404	{object} string	"Something went wrong, please refer to the error message."
//	@Router			/location [get]
//
// @Security ApiKeyAuth
func (a *APIEnv) GetLocations(c *gin.Context) {
	locations, err := db_calls.GetLocation(a.DB)
	if err != nil {
		log.Errorf(`Could not find any locations. Error: <%s>`, err)
		c.JSON(http.StatusNotFound, "Could not find any locations.")
		return
	}

	var locationDto []models.LocationDto
	dto.Map(&locationDto, locations)

	c.JSON(http.StatusOK, locationDto)
}

// GetLocationBySearch godoc
//
//	@Summary		Get one or more locations with search parameters
//	@Description	Get one or more locations by passing a location id and/or name as parameter.
//	@Tags			Locations
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	[]models.LocationDto
//	@Failure		404	{object} string	"Something went wrong, please refer to the error message."
//	@Router			/location/search [get]
//	@Param			id	query		string	 	false	"LocationId" example(1)
//	@Param			name	query		string	 	false	"Name of location" example(Office)
//
// @Security ApiKeyAuth
func (a *APIEnv) GetLocationBySearch(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")

	locations, err := db_calls.GetLocationBySearch(a.DB, id, name)
	if err != nil {
		log.Errorf(`Could not find any locations by id or name. id: <%s>; name: <%s>; Error: <%s>`, id, name, err)
		c.JSON(http.StatusNotFound, "Could not find any locations.")
		return
	}

	var locationDto []models.LocationDto
	dto.Map(&locationDto, locations)

	c.JSON(http.StatusOK, locationDto)
}

// CreateLocation godoc
//
//	@Summary		Create a new location
//	@Description	Create a new location by posting a list of location objects.
//	@Tags			Locations
//	@Accept			json
//	@Produce		json
//	@Success		201		{object}	[]models.LocationDto
//	@Failure		400	{object} string	"Something went wrong, please refer to the error message."
//	@Router			/location/new [post]
//	@Param			location	body		[]models.LocationPostDto	 true	"New Location"
//
// @Security ApiKeyAuth
func (a *APIEnv) CreateLocation(c *gin.Context) {
	var locations []models.Location
	if err := c.ShouldBindJSON(&locations); err != nil {
		log.Errorf(`Could not parse location from body. Error: <%s>`, err)
		c.JSON(http.StatusBadRequest, "Could not parse location from body.")
		return
	}

	if err := ex.Validator([]models.Location{}).Validate(locations); err != nil {
		log.Errorf(`Missing values in JSON. Error: <%s>`, err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	locations, err := db_calls.CreateLocation(a.DB, locations)
	if err != nil {
		log.Errorf(`Could not create location in db. Locations: <%#v> Error: <%s>`, locations, err)
		c.JSON(http.StatusBadRequest, "Could not create location. Name already exists.")
		return
	}

	var locationDto []models.LocationDto
	dto.Map(&locationDto, locations)

	c.JSON(http.StatusCreated, locationDto)
}

// UpdateLocation godoc
//
//	@Summary		Update a location
//	@Description	Update a location by posting a location object.
//	@Tags			Locations
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	models.LocationDto
//	@Failure		400	{object} string	"Something went wrong, please refer to the error message."
//	@Failure		404	{object} string	"Something went wrong, please refer to the error message."
//	@Router			/location/{id} [patch]
//	@Param			id	path		int	 	true	"LocationId"
//	@Param			location	body		models.LocationPostDto	 true	"Update Location"
//
// @Security ApiKeyAuth
func (a *APIEnv) UpdateLocation(c *gin.Context) {
	locationId := c.Param("id")

	if _, err := db_calls.GetLocationById(a.DB, locationId); err != nil {
		log.Errorf(`Could not find location by id. id: <%s>; Error: <%s>`, locationId, err)
		c.JSON(http.StatusNotFound, "Could not find location by id.")
		return
	}

	var location models.Location
	if err := c.ShouldBindJSON(&location); err != nil {
		log.Errorf(`Could not parse location details from body. Error: <%s>`, err)
		c.JSON(http.StatusBadRequest, "Could not parse location details from body.")
		return
	}

	if err := ex.Validator(models.Location{}).Validate(location); err != nil {
		log.Errorf(`Missing values in JSON. Error: <%s>`, err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	location, err := db_calls.UpdateLocation(a.DB, location)
	if err != nil {
		log.Errorf(`Could not update location in db. Location: <%#v> Error: <%s>`, location, err)
		c.JSON(http.StatusNotFound, "Could not update location.")
		return
	}

	var locationDto models.LocationDto
	dto.Map(&locationDto, location)

	c.JSON(http.StatusOK, locationDto)
}

// DeleteLocation godoc
//
//	@Summary		Delete a location
//	@Description	Delete a location by passing the location id as parameter.
//	@Tags			Locations
//	@Accept			json
//	@Produce		json
//	@Success		204 "Deleted successfully"
//	@Failure		404	{object} string	"Something went wrong, please refer to the error message."
//	@Router			/location/{id} [delete]
//	@Param			id	path		int	 	true	"LocationId"
//
// @Security ApiKeyAuth
func (a *APIEnv) DeleteLocation(c *gin.Context) {
	locationId := c.Param("id")

	location, err := db_calls.GetLocationById(a.DB, locationId)
	if err != nil {
		log.Errorf(`Could not find location by id. id: <%s>; Error: <%s>`, locationId, err)
		c.JSON(http.StatusNotFound, "Could not find location by id.")
		return
	}

	err = db_calls.DeleteLocation(a.DB, location)
	if err != nil {
		log.Errorf(`Could not delete location in db. Location: <%#v> Error: <%s>`, location, err)
		c.JSON(http.StatusNotFound, "Could not delete location.")
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
