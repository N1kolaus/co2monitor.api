package routes

import (
	"github.com/fminister/co2monitor.api/controllers"
	"github.com/fminister/co2monitor.api/middleware"
	"github.com/gin-gonic/gin"
)

func co2DataRoutes(superRoute *gin.RouterGroup) {
	co2DataRouter := superRoute.Group("/co2data")
	{
		co2DataRouter.Use(middleware.RequireAuth("X_API_KEY"))
		{
			co2DataRouter.GET("/", controllers.GetCo2Data)
		}
		co2DataRouter.Use(middleware.RequireAuth("X_API_KEY_POST"))
		{
			co2DataRouter.POST("/", controllers.GetCo2Data)
		}
	}
}
