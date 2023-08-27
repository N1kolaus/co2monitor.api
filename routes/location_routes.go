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

	// Middleware for routes that require X_API_KEY_ADMIN, i.e. elevated privileges
	adminRouter := locationRouter.Group("/")
	adminRouter.Use(middleware.RequireAuth("X_API_KEY_ADMIN"))
	{
		adminRouter.POST("/new", controllers.CreateLocation)
	}

	// Middleware for other routes that require X_API_KEY, i.e. no elevated privileges
	normalRouter := locationRouter.Group("/")
	normalRouter.Use(middleware.RequireAuth("X_API_KEY"))
	{
		normalRouter.GET("/", controllers.GetLocations)
		normalRouter.GET("/search", controllers.GetLocationBySearch)
	}
}
