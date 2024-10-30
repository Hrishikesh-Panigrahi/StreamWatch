package controller

import (
	"net/http"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/dbConnector"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
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

func CreateVideo() gin.HandlerFunc {
	return func(c *gin.Context) {

		UUIDid := uuid.New()

		var user models.User
		// userID := c.PostForm("userID")

		userID := 1

		dbConnector.DB.First(&user, userID)

		file, err := c.FormFile("video")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Video file is required"})
			return
		}

		videoPath := "./tempVideos/" + file.Filename
		if err := c.SaveUploadedFile(file, videoPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save video file"})
			return
		}

		video := models.Video{
			UserID: uint(userID),
			UUID:   UUIDid.String(),
			Name:   file.Filename,
			Path:   videoPath,
		}

		dbConnector.DB.Create(&video)

		c.JSON(http.StatusOK, gin.H{"message": "Video is created", "videoPath": videoPath})

	}
}
