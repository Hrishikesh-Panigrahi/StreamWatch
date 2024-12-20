package main

import (
	"github.com/Hrishikesh-Panigrahi/StreamWatch/dbConnector"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	dbConnector.LoadEnvVariables()
	dbConnector.Connection()
	dbConnector.SyncDB()
}

func main() {
	router := gin.Default()
	router.Static("/static", "./web/static")
	router.Static("/tempVideos", "./tempVideos")

	router.LoadHTMLGlob("web/templates/**/*.html")

	routers := router.Group("/")
	routes.Routes(routers)

	router.Run(":8080")
}