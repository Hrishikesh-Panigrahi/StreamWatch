package controller

import (
	"net/http"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
	"github.com/gin-gonic/gin"
)

func HomePageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: get all the videos.

		type Data struct {
			Title   string
			Message string
		}

		data := Data{Title: "Index", Message: "this is index"}

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
