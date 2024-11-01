package controller

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

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
		// temp user id for testing
		userID := 1

		dbConnector.DB.First(&user, userID)

		file, err := c.FormFile("video")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Video file is required"})
			return
		}

		// Construct folder path (optionally, add UUID or timestamp)
		folderName := fmt.Sprintf("%s_%s_%s", file.Filename, user.Name, UUIDid.String())
		folderPath := "./tempVideos/" + folderName
		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create folder for video"})
			return
		}

		originalVideoPath := folderPath + "/" + file.Filename
		if err := c.SaveUploadedFile(file, originalVideoPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save video file"})
			return
		}

		resolutions := []string{"480p", "720p", "1080p"}

		for _, res := range resolutions {
			outputPath := fmt.Sprintf("%s/%s_%s.mp4", folderPath, file.Filename, res)
			err := createResolution(originalVideoPath, outputPath, res)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create %s resolution", res)})
				return
			}
		}

		// Truncate filename if it's too long
		filename := file.Filename
		if len(filename) > 50 {
			filename = filename[:50]
		}

		video := models.Video{
			UserID: uint(userID),
			UUID:   UUIDid.String(),
			Name:   filename,
			Path:   originalVideoPath,
		}

		dbConnector.DB.Create(&video)

		c.JSON(http.StatusOK, gin.H{"message": "Video is created", "videoPath": originalVideoPath})

	}
}

// Helper function to create video resolution using FFmpeg
func createResolution(inputPath, outputPath, resolution string) error {
	var ffmpegCmd *exec.Cmd
	switch resolution {
	case "480p":
		ffmpegCmd = exec.Command("ffmpeg", "-i", inputPath, "-vf", "scale=854:480", outputPath)
	case "720p":
		ffmpegCmd = exec.Command("ffmpeg", "-i", inputPath, "-vf", "scale=1280:720", outputPath)
	case "1080p":
		ffmpegCmd = exec.Command("ffmpeg", "-i", inputPath, "-vf", "scale=1920:1080", outputPath)
	default:
		return fmt.Errorf("unsupported resolution")
	}

	if err := ffmpegCmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg error: %v", err)
	}
	return nil
}
