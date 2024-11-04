package controller

import (
	"fmt"
	"net/http"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/dbConnector"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HomePageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var videos []models.Video

		err := dbConnector.DB.Preload("User").Find(&videos).Error

		if err != nil {
			fmt.Printf("Error retrieving videos: %v\n", err)

			type ErrorData struct {
				Title   string
				Message string
			}

			data := ErrorData{
				Title:   "Error",
				Message: "An error occurred while loading videos. Please try again later.",
			}

			render.RenderHtml(c, http.StatusInternalServerError, "error.html", data)
			return
		}

		type Data struct {
			Title   string
			Message string
			Videos  []models.Video
		}

		data := Data{Title: "Index", Message: "this is index", Videos: videos}

		render.RenderHtml(c, http.StatusOK, "base.html", data)
	}
}

func VideoPageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		VideoUUID := c.Query("videoUUID")
		var video models.Video

		type Data struct {
			Title   string
			Message string
			Video   models.Video
		}

		if VideoUUID == "" {
			data := Data{Title: "Error", Message: "No video file specified."}
			render.RenderHtml(c, http.StatusBadRequest, "base.html", data)
			return
		}

		err := dbConnector.DB.Preload("User").Where("uuid = ?", VideoUUID).First(&video).Error

		if err != nil {
			type ErrorData struct {
				Title   string
				Message string
			}

			fmt.Printf("Error retrieving video: %v\n", err)
			if err == gorm.ErrRecordNotFound {
				fmt.Println("Video not found")
				data := ErrorData{Title: "Error", Message: "Video not found with the provided UUID."}
				render.RenderHtml(c, http.StatusNotFound, "base.html", data)
			} else {
				fmt.Printf("Error retrieving video: %v\n", err)
				data := ErrorData{Title: "Error", Message: "An error occurred while fetching the video."}
				render.RenderHtml(c, http.StatusInternalServerError, "base.html", data)
			}
			return
		}

		fmt.Println(video.Name)
		data := Data{Title: "Video", Message: "this is index", Video: video}

		render.RenderHtml(c, http.StatusOK, "base.html", data)
	}
}
