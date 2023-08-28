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
	co2DataRouter.Use(middleware.RequireApiKey)
	{
		co2DataRouter.GET("/:id/search", controllers.GetCo2DataByTimeFrame)
		co2DataRouter.GET("/:id/latest", controllers.GetLatestCo2Data)
		co2DataRouter.POST("/new", controllers.CreateCo2Data)
	}

}
