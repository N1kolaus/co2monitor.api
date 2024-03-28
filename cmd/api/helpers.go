package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/FMinister/co2monitor-api/internal/validator"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]any

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	// Use http.MaxBytesReader() to limit the size of the request body to 1MB.
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		var maxBytesError http.MaxBytesError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)

		case errors.Is(err, &maxBytesError):
			return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)

		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}
	}

	err = decoder.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

func (app *application) writeJSON(w http.ResponseWriter, statusCode int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(js)

	return nil
}

func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

func (app *application) validateTimeFrame(digits int, unit string) (time.Duration, error) {
	maxTimeFrame := time.Duration(app.config.maxTimeFrameDays) * 24 * time.Hour

	if unit != "m" && unit != "h" && unit != "d" {
		app.logger.Info(fmt.Sprintf("Invalid unit for time frame: %s", unit))
		return 0, errors.New("invalid unit for time frame, valid units are: m, h, d")
	}

	if unit == "m" && digits > 60*24*app.config.maxTimeFrameDays {
		app.logger.Info(fmt.Sprintf("Digits to large for unit: %d%s; returning max time frame %s", digits, unit, maxTimeFrame))
		return maxTimeFrame, nil
	}
	if unit == "h" && digits > 24*app.config.maxTimeFrameDays {
		app.logger.Info(fmt.Sprintf("Digits to large for unit: %d%s; returning max time frame %s", digits, unit, maxTimeFrame))
		return maxTimeFrame, nil
	}
	if unit == "d" && digits > app.config.maxTimeFrameDays {
		app.logger.Info(fmt.Sprintf("Digits to large for unit: %d%s; returning max time frame %s", digits, unit, maxTimeFrame))
		return maxTimeFrame, nil
	}

	switch unit {
	case "m":
		return time.Duration(digits) * time.Minute, nil
	case "h":
		return time.Duration(digits) * time.Hour, nil
	case "d":
		return time.Duration(digits) * time.Hour * 24, nil
	default:
		return 6 * time.Hour, nil
	}
}

func (app *application) readString(qs url.Values, key string, defaultValue string) string {
	value := qs.Get(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func (app *application) readCSV(qs url.Values, key string, defaultValue []string) []string {
	value := qs.Get(key)
	if value == "" {
		return defaultValue
	}

	return strings.Split(value, ",")
}

func (app *application) readInt(qs url.Values, key string, defaultValue int, v *validator.Validator) int {
	value := qs.Get(key)
	if value == "" {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		v.AddError(key, "must be an integer value")
		return defaultValue
	}

	return intValue
}

// recovery method to use in go routines
func (app *application) background(fn func()) {
	app.wg.Add(1)

	go func() {
		defer app.wg.Done()

		defer func() {
			if err := recover(); err != nil {
				app.logger.Error(fmt.Sprintf("%v", err))
			}
		}()

		fn()
	}()
}
