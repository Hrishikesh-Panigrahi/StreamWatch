package controller

import (
	"fmt"
	"net/http"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/dbConnector"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
	"github.com/gin-gonic/gin"
)

func HomePageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var videos []models.Video

		dbConnector.DB.Preload("User").Find(&videos)

		type Data struct {
			Title              string
			Message            string
			Videos             []models.Video
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
			render.RenderHtml(c, http.StatusBadRequest, "error.html", data)
			return
		}

		err := dbConnector.DB.Preload("User").Where("uuid = ?", VideoUUID).First(&video).Error

		if err != nil {
			data := Data{Title: "Error", Message: "Video not found."}
			render.RenderHtml(c, http.StatusNotFound, "error.html", data)
			return
		}
		fmt.Println(video.Name)
		data := Data{Title: "Video", Message: "this is index", Video: video}

		render.RenderHtml(c, http.StatusOK, "base.html", data)
	}
}
