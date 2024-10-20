package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func StreamHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		filename := c.Param("filename")
		fmt.Println(filename)
		file, err := os.Open("./tempVideos/" + filename + ".mp4")

		if err != nil {
			c.String(http.StatusNotFound, "Video not found.")
			return
		}
		defer file.Close()

		c.Header("Content-Type", "video/mp4")
		c.Writer.Header().Set("Content-Disposition", "inline; filename="+filename+".mp4")
		buffer := make([]byte, 64*1024)
		io.CopyBuffer(c.Writer, file, buffer)
	}
}
