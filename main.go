package main

import (
	"github.com/fminister/co2monitor.api/db"
	"github.com/fminister/co2monitor.api/initializers"
	"github.com/gin-gonic/gin"
)

func main() {
	initializers.LoadEnvVariables()
	db.ConnectToDb()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

// CompileDaemon -command="./co2monitor.api
