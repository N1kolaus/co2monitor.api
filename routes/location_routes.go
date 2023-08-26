package routes

import (
	"github.com/fminister/co2monitor.api/controllers"
	"github.com/fminister/co2monitor.api/db"
	"github.com/fminister/co2monitor.api/middleware"
	"github.com/gin-gonic/gin"
)

func locationRoutes(superRoute *gin.RouterGroup) {
	controllers := &controllers.APIEnv{
		DB: db.GetDB(),
	}

	locationRouter := superRoute.Group("/location")

	// Middleware für POST-Routen, die X_API_KEY_POST benötigen
	postMiddleware := locationRouter.Group("/")
	postMiddleware.Use(middleware.RequireAuth("X_API_KEY_POST"))
	{
		postMiddleware.POST("/new", controllers.CreateLocation)
	}

	// Middleware für andere Routen, die X_API_KEY benötigen
	getMiddleware := locationRouter.Group("/")
	getMiddleware.Use(middleware.RequireAuth("X_API_KEY"))
	{
		getMiddleware.GET("/", controllers.GetLocations)
		getMiddleware.GET("/search", controllers.GetLocationBySearch)
	}
}
