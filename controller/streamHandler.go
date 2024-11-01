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

		filefolder := video.Name
		filename := "test.mp4_480p.mp4"
		path := fmt.Sprintf("./tempVideos/%s_%s_%s/%s", filefolder, user.Name, fileUUID, filename)

		file, err := os.Open(path)
		if err != nil {
			c.String(http.StatusNotFound, "Video not found.")
			return
		}
		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			c.String(http.StatusInternalServerError, "Unable to get file info.")
			return
		}
		fileSize := fileInfo.Size()

		// Set headers
		c.Header("Content-Type", "video/mp4")
		c.Header("Accept-Ranges", "bytes")

		// Handle range requests
		rangeHeader := c.GetHeader("Range")
		if rangeHeader == "" {
			c.Header("Content-Length", strconv.FormatInt(fileSize, 10))
			http.ServeContent(c.Writer, c.Request, filename, fileInfo.ModTime(), file)
			return
		}

		// Parse range header
		var start, end int64
		fmt.Sscanf(rangeHeader, "bytes=%d-%d", &start, &end)
		if end == 0 {
			end = fileSize - 1
		}

		// Validate range
		if start < 0 || end >= fileSize || start > end {
			c.String(http.StatusRequestedRangeNotSatisfiable, "Invalid range")
			return
		}

		// Set partial content headers
		c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, fileSize))
		c.Header("Content-Length", strconv.FormatInt(end-start+1, 10))
		c.Status(http.StatusPartialContent)

		// Seek to the start position and copy the data to the response
		file.Seek(start, 0)
		bufferSize := 64 * 1024
		buffer := make([]byte, bufferSize)
		for {
			if remaining := end - start + 1; remaining < int64(bufferSize) {
				buffer = buffer[:remaining]
			}
			n, err := file.Read(buffer)
			if err != nil {
				break
			}
			c.Writer.Write(buffer[:n])
			start += int64(n)
			if start > end {
				break
			}
		}
	}
}
