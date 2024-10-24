package main

import (
	"github.com/Hrishikesh-Panigrahi/StreamWatch/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./web/static")

	router.LoadHTMLGlob("web/templates/**/*.html")

	router.GET("/", controller.HomePageHandler())

	// API for video streaming
	router.GET("/stream/:filename", controller.StreamHandler())

	// frontend url -- The request responds to a url matching: /video?filename=example.mp4
	router.GET("/video", controller.VideoPageHandler())

	router.Run(":8080")
}
