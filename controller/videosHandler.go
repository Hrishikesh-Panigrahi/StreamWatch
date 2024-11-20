package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/dbConnector"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/utils"
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

func CreateVideoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		form, _ := c.MultipartForm()
		files := form.File["videoFile"]

		if len(files) > 1 {
			fmt.Println("Multiple files uploaded")
			// for _, file := range files {
			// 	log.Println(file.Filename)

			// 	// Upload the file to specific dst.
			// 	c.SaveUploadedFile(file, dst)
			// }
			return
		}

		UUIDid := uuid.New()

		user := utils.GetUserFromCache(c)

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
		GenerateMasterPlaylistErr := utils.GenerateMasterPlaylist(c, folderPath, originalVideoPath)

		if GenerateMasterPlaylistErr != nil {
			render.RenderError(c, http.StatusInternalServerError, "Failed to create video")
			return
		}

		video := models.Video{
			UserID:            user.ID,
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
