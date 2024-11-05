package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/dbConnector"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
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

		cookieuser, _ := c.Get("user")
		userID := cookieuser.(models.User).ID

		var user models.User
		err := dbConnector.DB.First(&user, userID).Error

		if err != nil {
			fmt.Printf("Error retrieving User: %v\n", err)

			type ErrorData struct {
				Title   string
				Message string
			}

			data := ErrorData{
				Title:   "Error",
				Message: "An error occurred while retrieving User. Please try after logging in again.",
			}

			render.RenderHtml(c, http.StatusInternalServerError, "base.html", data)
			return
		}

		file, err := c.FormFile("video")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Video file is required"})
			return
		}

		filename := file.Filename
		if len(filename) > 50 {
			filename = filename[:50]
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

		// resolutions := []string{"480p", "720p", "1080p"}

		// for _, res := range resolutions {
		// 	// tod0: filename:= timestamp ... for sorting
		// 	outputPath := fmt.Sprintf("%s/%s_%s.mp4", folderPath, file.Filename, res)
		// 	err := utils.CreateResolution(originalVideoPath, outputPath, res)
		// 	if err != nil {
		// 		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create %s resolution", res)})
		// 		return
		// 	}
		// }

		masterPlaylist := folderPath + "/master.m3u8"
		cmd := exec.Command("ffmpeg", "-i", originalVideoPath,
			// HLS output options
			"-f", "hls", "-hls_time", "4", "-hls_playlist_type", "vod",
			"-hls_segment_filename", folderPath+"/segment_%03d.ts",
			folderPath+"/master.m3u8",
		)

		var stderrOutput bytes.Buffer
		cmd.Stderr = &stderrOutput

		// Run FFmpeg command
		if err := cmd.Run(); err != nil {
			// Log the FFmpeg error output for debugging
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to transcode video",
				"details": stderrOutput.String(),
			})
			return
		}

		// Truncate filename if it's too long

		video := models.Video{
			UserID:            uint(userID),
			UUID:              UUIDid.String(),
			Name:              filename,
			Path:              masterPlaylist,
			OriginalVideoPath: originalVideoPath,
		}

		dbConnector.DB.Create(&video)

		c.JSON(http.StatusOK, gin.H{"message": "Video is created", "videoPath": originalVideoPath})

	}
}
