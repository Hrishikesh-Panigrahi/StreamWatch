package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to our video streaming platform!")
	})

	router.GET("/videos", func(c *gin.Context) {
		c.String(http.StatusOK, "List of all videos")
	})

	router.GET("/videos/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Video id: %s", id)
	})

	router.POST("/videos", func(c *gin.Context) {
		c.String(http.StatusOK, "Video is created")
	})

	router.GET("/stream/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		fmt.Println(filename)
		file, err := os.Open("./tempVideos/" + filename + ".mp4")

		if err != nil {
			c.String(http.StatusNotFound, "Video not found.")
			return
		}
		defer file.Close()

		c.Header("Content-Type", "video/mp4")
		buffer := make([]byte, 64*1024) // 64KB buffer size
		io.CopyBuffer(c.Writer, file, buffer)
	})

	router.Run(":8080")

}
