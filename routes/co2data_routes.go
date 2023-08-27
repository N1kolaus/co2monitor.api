package routes

import (
	"github.com/fminister/co2monitor.api/controllers"
	"github.com/fminister/co2monitor.api/db"
	"github.com/fminister/co2monitor.api/middleware"
	"github.com/gin-gonic/gin"
)

func co2DataRoutes(superRoute *gin.RouterGroup) {
	controllers := &controllers.APIEnv{
		DB: db.GetDB(),
	}
	co2DataRouter := superRoute.Group("/co2data")

	// Middleware for routes that require X_API_KEY_ADMIN, i.e. elevated privileges
	adminRouter := co2DataRouter.Group("/")
	adminRouter.Use(middleware.RequireAuth("X_API_KEY_ADMIN"))
	{
		adminRouter.POST("/new", controllers.CreateCo2Data)
	}

	// Middleware for other routes that require X_API_KEY, i.e. no elevated privileges
	normalRouter := co2DataRouter.Group("/")
	normalRouter.Use(middleware.RequireAuth("X_API_KEY"))
	{
		normalRouter.GET("/search", controllers.GetCo2DataBySearch)
	}
}
