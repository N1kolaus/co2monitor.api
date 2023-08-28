package controllers

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/dranikpg/dto-mapper"
	"github.com/fminister/co2monitor.api/db/db_calls"
	ex "github.com/fminister/co2monitor.api/extensions"
	"github.com/fminister/co2monitor.api/models"
	"github.com/gin-gonic/gin"
)

func (a *APIEnv) GetCo2DataByTimeFrame(c *gin.Context) {
	locationId := c.Param("id")
	period := c.Query("period")

	if _, err := db_calls.GetLocationById(a.DB, locationId); err != nil {
		log.Errorf(`Could not find any location with this id: <%s>. Error: <%s>`, locationId, err)
		c.JSON(http.StatusNotFound, fmt.Sprintf(`Could not find any location with this id: <%s>.`, locationId))
		return
	}

	duration := ex.ValidateTimeDuration(period)

	var co2Data []models.Co2Data
	co2Data, err := db_calls.GetCo2DataByTimeFrame(a.DB, locationId, duration)
	if err != nil {
		log.Errorf(`Could not find any co2 data with this locationId: <%s>. Error: <%s>`, locationId, err)
		c.JSON(http.StatusNotFound, fmt.Sprintf(`Could not find any co2 data with this locationId: <%s>.`, locationId))
		return
	}

	var co2DataDto []models.Co2DataDto
	dto.Map(&co2DataDto, co2Data)

	c.JSON(http.StatusOK, co2DataDto)
}

func (a *APIEnv) GetLatestCo2Data(c *gin.Context) {
	locationId := c.Param("id")

	co2Data, err := db_calls.GetLatestCo2Data(a.DB, locationId)
	if err != nil {
		log.Errorf(`Could not find any co2 data with this locationId: <%s>. Error: <%s>`, locationId, err)
		c.JSON(http.StatusNotFound, fmt.Sprintf(`Could not find any co2 data with this locationId: <%s>.`, locationId))
		return
	}

	var co2DataDto models.Co2DataDto
	dto.Map(&co2DataDto, co2Data)

	c.JSON(http.StatusOK, co2DataDto)
}

func (a *APIEnv) CreateCo2Data(c *gin.Context) {
	var co2Data []models.Co2Data
	if err := c.ShouldBindJSON(&co2Data); err != nil {
		log.Errorf(`Could not parse co2 data from body. Error: <%s>`, err)
		c.JSON(http.StatusBadRequest, "Could not parse co2 data from body.")
		return
	}

	if err := ex.Validator([]models.Co2Data{}).Validate(co2Data); err != nil {
		log.Errorf(`Missing values in JSON. Error: <%s>`, err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	co2Data, err := db_calls.CreateCo2Data(a.DB, co2Data)
	if err != nil {
		log.Errorf(`Could not create co2 data in db. Co2Data: <%#v> Error: <%s>`, co2Data, err)
		c.JSON(http.StatusBadRequest, "Could not create co2 data.")
		return
	}

	var co2DataDto []models.Co2DataDto
	dto.Map(&co2DataDto, co2Data)

	c.JSON(http.StatusCreated, co2DataDto)
}
