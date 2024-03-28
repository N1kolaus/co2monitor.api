package main

import (
	"net/http"

	"github.com/FMinister/co2monitor-api/internal/data"
	"github.com/FMinister/co2monitor-api/internal/validator"
)

func (app *application) co2DataLatestHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	co2Data, err := app.models.Co2.GetLatest(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"latest": co2Data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listCo2DataByTimeFrameHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var input struct {
		Digits int
		Unit   string
	}

	v := validator.New()
	qs := r.URL.Query()

	input.Digits = app.readInt(qs, "digits", 6, v)
	input.Unit = app.readString(qs, "unit", "m")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	timeFrame, err := app.validateTimeFrame(input.Digits, input.Unit)
	if err != nil {
		app.failedValidationResponse(w, r, map[string]string{"unit": err.Error()})
		return
	}

	co2Data, err := app.models.Co2.GetByTimeFrame(id, timeFrame)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"data": co2Data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createCo2DataHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var input struct {
		Co2      int     `json:"co2"`
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	co2Data := &data.Co2Data{
		LocationID: id,
		Co2:        input.Co2,
		Temp:       input.Temp,
		Humidity:   input.Humidity,
	}

	v := validator.New()

	if data.ValidateCo2Data(v, co2Data); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Co2.Insert(co2Data)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"latest": co2Data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
