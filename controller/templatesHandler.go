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

			render.RenderError(c, http.StatusInternalServerError, "An error occurred while fetching the videos. Please try again later.")
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
			Title     string
			Message   string
			Video     models.Video
			LikeCount int64
			WatchLog  models.WatchLog
		}

		if VideoUUID == "" {
			fmt.Println("Video UUID not provided")
			render.RenderError(c, http.StatusBadRequest, "Video UUID not provided.")
			return
		}

		err := dbConnector.DB.Preload("User").Where("uuid = ?", VideoUUID).First(&video).Error

		if err != nil {
			fmt.Printf("Error retrieving video: %v\n", err)
			if err == gorm.ErrRecordNotFound {
				fmt.Println("Video not found")
				render.RenderError(c, http.StatusNotFound, "Video not found.")
			} else {
				fmt.Printf("Error retrieving video: %v\n", err)
				render.RenderError(c, http.StatusInternalServerError, "An error occurred while fetching the video. Please try again later.")
			}
			return
		}

		var likeCount int64
		if err := dbConnector.DB.Model(&models.Likes{}).Where("video_id = ?", video.ID).Count(&likeCount).Error; err != nil {
			fmt.Printf("Error retrieving like count: %v\n", err)
			render.RenderError(c, http.StatusInternalServerError, "An error occurred while fetching the like count. Please try again later.")
			return
		}

		var watchlog models.WatchLog
		if err := dbConnector.DB.Where("video_id = ? AND user_id = ?", video.ID, 1).First(&watchlog).Error; err != nil {
			fmt.Printf("Error retrieving watch log: %v\n", err)
			render.RenderError(c, http.StatusInternalServerError, "An error occurred while fetching the watch log. Please try again later.")
			return
		}

		data := Data{Title: "Video", Message: "this is index", Video: video, LikeCount: likeCount, WatchLog: watchlog}

		render.RenderHtml(c, http.StatusOK, "base.html", data)
	}
}

func LoginPageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Data struct {
			Title   string
			Message string
		}

		data := Data{Title: "Login", Message: "this is Login page"}
		render.RenderHtml(c, http.StatusOK, "base.html", data)
	}
}

func RegisterPageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Data struct {
			Title   string
			Message string
		}

		data := Data{Title: "register", Message: "this is registration page"}
		render.RenderHtml(c, http.StatusOK, "base.html", data)
	}
}

func CreateVideoPageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Data struct {
			Title   string
			Message string
		}

		data := Data{Title: "Create Video", Message: "this is create video page"}

		render.RenderHtml(c, http.StatusOK, "base.html", data)
	}
}
