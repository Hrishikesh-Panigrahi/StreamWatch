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

func CreateVideoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		UUIDid := uuid.New()

		cookieuser, exists := c.Get("user")

		if !exists {
			render.RenderError(c, http.StatusUnauthorized, "User not logged in. Please login to upload video.")
			return
		}

		userID := cookieuser.(models.User).ID

		fmt.Println(userID)

		var user models.User
		if err := dbConnector.DB.First(&user, userID).Error; err != nil {
			fmt.Printf("Error retrieving User: %v\n", err)

			render.RenderError(c, http.StatusInternalServerError, "An error occurred while fetching the user. Please try again later.")
			return
		}

		file, fileHeader, err := c.Request.FormFile("videoFile")
		if err != nil {
			render.RenderError(c, http.StatusBadRequest, "No video uploaded. Please upload a video and submit the form.")
			return
		}
		defer file.Close()

		if file == nil || fileHeader == nil {
			render.RenderError(c, http.StatusBadRequest, "No file uploaded.")
			return
		}

		fileType := fileHeader.Header.Get("Content-Type")
		if fileType != "video/mp4" && fileType != "video/x-matroska" {
			render.RenderError(c, http.StatusUnsupportedMediaType, "Unsupported file type. Only MP4 and MKV are allowed.")
			return
		}

		name := c.PostForm("videoTitle")
		tags := c.PostForm("tags")
		description := c.PostForm("description")

		fmt.Print(fileHeader.Filename, name, tags, description)

		filename := name
		if len(filename) > 50 {
			filename = filename[:50]
		}

		// Construct folder path (optionally, add UUID or timestamp)
		folderName := fmt.Sprintf("%s_%s_%s", name, user.Name, UUIDid.String())
		folderPath := "./tempVideos/" + folderName
		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			render.RenderError(c, http.StatusInternalServerError, "Failed to create video")
			return
		}

		originalVideoPath := folderPath + "/" + name + ".mp4"
		if err := c.SaveUploadedFile(fileHeader, originalVideoPath); err != nil {
			render.RenderError(c, http.StatusInternalServerError, "Failed to save video")
			return
		}

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
			fmt.Println("Error running FFmpeg command: ", err)
			fmt.Println("FFmpeg stderr: ", stderrOutput.String())
			render.RenderError(c, http.StatusInternalServerError, "Failed to create video")
			return
		}

		video := models.Video{
			UserID:            uint(userID),
			UUID:              UUIDid.String(),
			Name:              filename,
			Tags:              tags,
			Description:       description,
			Path:              masterPlaylist,
			OriginalVideoPath: originalVideoPath,
		}

		if err := dbConnector.DB.Create(&video).Error; err != nil {
			fmt.Printf("Error creating video: %v\n", err)
			render.RenderError(c, http.StatusInternalServerError, "Failed to create or save the video. Check your internet connection and try again.")
			return
		}

		render.Redirect(c, "/", http.StatusFound)

	}
}
