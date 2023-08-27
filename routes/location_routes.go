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
	locationRouter.Use(middleware.RequireApiKey)
	{
		locationRouter.GET("/", controllers.GetLocations)
		locationRouter.GET("/search", controllers.GetLocationBySearch)
		locationRouter.POST("/new", controllers.CreateLocation)
		locationRouter.PATCH("/:id", controllers.UpdateLocation)
	}
}
