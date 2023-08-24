package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fminister/co2monitor.api/db"
	"github.com/fminister/co2monitor.api/initializers"
	"github.com/fminister/co2monitor.api/routes"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	db.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	f, _ := os.Create("logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	log.SetOutput(io.MultiWriter(f, os.Stdout))

	app := gin.New()

	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] - %s \"%s %s %s %d %s %s\"\n",
			param.TimeStamp.Format(time.RFC1123),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.ErrorMessage,
		)
	}))
	app.Use(gin.Recovery())
	app.Use(gzip.Gzip(gzip.DefaultCompression))

	router := app.Group("/api")
	routes.AddRoutes(router)

	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	app.Run()
}

// CompileDaemon -command="./co2monitor.api
