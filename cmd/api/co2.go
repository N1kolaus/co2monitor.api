package main

import (
	"fmt"
	"net/http"

	"github.com/FMinister/co2monitor-api/internal/validator"
)

func (app *application) co2DataByTimeFrameHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Digits int
		Unit   string
	}

	v := validator.New()
	qs := r.URL.Query()

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

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

	// co2Data, err := app.models.Co2Data.GetByTimeFrame(timeFrame)
	// if err != nil {
	// 	app.serverErrorResponse(w, r, err)
	// 	return
	// }

	err = app.writeJSON(w, http.StatusOK, envelope{"time_frame": fmt.Sprintf("id: %d; time_frame: %v", id, timeFrame)}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
