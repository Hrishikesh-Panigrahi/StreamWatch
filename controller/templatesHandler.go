package controller

import (
	"net/http"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
	"github.com/gin-gonic/gin"
)

func HomePageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: get all the videos.
		type Videos struct {
			Title    string
			Message  string
			videoSRC string
		}

		videos := []Videos{
			{Title: "Sample 1", Message: "This is Sample Video 1", videoSRC: "http://localhost:8080/video?filename=sample1.mp4"},
			{Title: "Sample 2", Message: "This is Sample Video 2", videoSRC: "http://localhost:8080/video?filename=sample2.mp4"},
			{Title: "Sample 3", Message: "This is Sample Video 3", videoSRC: "http://localhost:8080/video?filename=sample3.mp4"},
			{Title: "Sample 4", Message: "This is Sample Video 4", videoSRC: "http://localhost:8080/video?filename=sample4.mp4"},
		}

		type Data struct {
			Title   string
			Message string
			Videos  []Videos
		}

		data := Data{Title: "Index", Message: "this is index", Videos: videos}

		render.RenderHtml(c, http.StatusOK, "base.html", data)
	}
}

func VideoPageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		filename := c.Query("filename")
		type Data struct {
			Title    string
			Message  string
			Filename string
		}

		if filename == "" {
			data := Data{Title: "Error", Message: "No video file specified."}
			render.RenderHtml(c, http.StatusBadRequest, "error.html", data)
			return
		}

		data := Data{Title: "Video", Message: "this is index", Filename: filename}

		render.RenderHtml(c, http.StatusOK, "base.html", data)
	}
}
