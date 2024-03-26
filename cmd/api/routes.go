package main

import (
	"expvar"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	if app.config.env != "production" {
		router.Handler(http.MethodGet, "/v2/metrics", expvar.Handler())
	}

	router.HandlerFunc(http.MethodGet, "/v2/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v2/co2data/:id", app.co2DataByTimeFrameHandler)

	return router
}
