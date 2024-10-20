package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllVideos() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "List of all videos")
	}
}

func GetVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Video id: %s", id)
	}
}

func AddVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Video is created")
	}
}
