package controller

import (
	"net/http"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
	"github.com/gin-gonic/gin"
)

func HomePageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: get all the videos.


		render.RenderHtml(c, http.StatusOK, "index.html", nil)
	}
}

func VideoPageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		filename := c.Query("filename")

		if filename == "" {
			render.RenderHtml(c, http.StatusBadRequest, "error.html", "No video file specified.")
			return
		}

		render.RenderHtml(c, http.StatusOK, "videoTemplate.html", filename)
	}
}
