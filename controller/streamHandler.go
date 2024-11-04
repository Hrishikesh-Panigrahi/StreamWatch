package controller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/dbConnector"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/gin-gonic/gin"
)

func StreamHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		userID := 1

		dbConnector.DB.First(&user, userID)

		fileUUID := c.Param("UUID")

		var video models.Video
		dbConnector.DB.Where("uuid = ?", fileUUID).First(&video)

		// Define the folder where the HLS files are stored
		filefolder := video.Name
		folderPath := fmt.Sprintf("./tempVideos/%s_%s_%s", filefolder, user.Name, fileUUID)

		// Check if the requested file is a playlist or segment
		fileType := c.Query("type")
		var filePath string

		if fileType == "playlist" {
			// Serve the master playlist file
			filePath = fmt.Sprintf("%s/master.m3u8", folderPath)
			c.Header("Content-Type", "application/x-mpegURL")
		} else {
			// Serve the requested segment file (e.g., segment_001.ts)
			segment := c.Query("segment")
			filePath = fmt.Sprintf("%s/%s", folderPath, segment)
			c.Header("Content-Type", "video/MP2T")
		}

		// Open and serve the requested file
		file, err := os.Open(filePath)
		if err != nil {
			c.String(http.StatusNotFound, "File not found.")
			return
		}
		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			c.String(http.StatusInternalServerError, "Unable to get file info.")
			return
		}

		c.Header("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))
		http.ServeContent(c.Writer, c.Request, fileInfo.Name(), fileInfo.ModTime(), file)
	}
}
