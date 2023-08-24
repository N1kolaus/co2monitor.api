package routes

import (
	"github.com/fminister/co2monitor.api/middleware"
	"github.com/gin-gonic/gin"
)

func locationRoutes(superRoute *gin.RouterGroup) {
	locationRouter := superRoute.Group("/location")
	{
		locationRouter.Use(middleware.RequireAuth)
		{
		}
	}
}
