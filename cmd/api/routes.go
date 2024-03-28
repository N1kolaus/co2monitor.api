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

	router.HandlerFunc(http.MethodGet, "/v2/co2data/:id/latest", app.co2DataLatestHandler)
	router.HandlerFunc(http.MethodGet, "/v2/co2data/:id", app.listCo2DataByTimeFrameHandler)
	router.HandlerFunc(http.MethodPost, "/v2/co2data/:id", app.createCo2DataHandler)

	router.HandlerFunc(http.MethodGet, "/v2/location", app.listLocationsHandler)
	router.HandlerFunc(http.MethodPost, "/v2/location", app.createLocationHandler)
	router.HandlerFunc(http.MethodPatch, "/v2/location/:id", app.updateLocationHandler)
	router.HandlerFunc(http.MethodDelete, "/v2/location/:id", app.deleteLocationHandler)

	router.HandlerFunc(http.MethodGet, "/v2/user", app.getUserHandler)
	router.HandlerFunc(http.MethodPost, "/v2/user", app.createUserHandler)
	router.HandlerFunc(http.MethodPatch, "/v2/user", app.updateUserHandler)
	router.HandlerFunc(http.MethodDelete, "/v2/user/:id", app.deleteUserHandler)

	return router
}
