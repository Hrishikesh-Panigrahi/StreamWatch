package main

import (
	"net/http"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	router.GET("/videos", controller.AllVideos())

	router.GET("/videos/:id", controller.GetVideo())

	router.POST("/videos", controller.AddVideo())

	router.GET("/stream/:filename", controller.StreamHandler())

	router.GET("/video/:filename", controller.VideoPageHandler())

	router.Run(":8080")

}
