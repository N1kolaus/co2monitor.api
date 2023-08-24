package routes

import "github.com/gin-gonic/gin"

func AddRoutes(superRoute *gin.RouterGroup) {
	co2DataRoutes(superRoute)
	locationRoutes(superRoute)
}
