package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/log"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/fminister/co2monitor.api/db"
	"github.com/fminister/co2monitor.api/docs"
	"github.com/fminister/co2monitor.api/initializers"
	"github.com/fminister/co2monitor.api/routes"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	db.ConnectToDb()
	initializers.SyncDatabase()
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-KEY
// @description Paste in the api key
func main() {
	f, _ := os.Create("logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	log.SetTimeFormat(fmt.Sprintf("[%s]", time.RFC1123))
	log.SetReportCaller(true)
	log.SetReportTimestamp(true)
	log.SetOutput(io.MultiWriter(f, os.Stdout))

	app := gin.New()

	log.Fatal(autotls.Run(app, os.Getenv("APP_HOST")))

	docs.SwaggerInfo.Title = "CO2 Monitor API"
	docs.SwaggerInfo.Description = "CO2 Monitor API for the CO2 Monitor project."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("APP_HOST")
	docs.SwaggerInfo.BasePath = "/api"
	if os.Getenv("APP_ENV") != "development" {
		docs.SwaggerInfo.Schemes = []string{"https"}
	} else {
		docs.SwaggerInfo.Schemes = []string{"http"}
	}

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

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "The requested route does not exist."})
	})

	app.Run()
}

// CompileDaemon -command="./co2monitor.api
