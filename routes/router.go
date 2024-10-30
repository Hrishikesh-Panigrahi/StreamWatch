package routes

import (
	"net/http"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/controller"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
	"github.com/gin-gonic/gin"
)

func Routes(superRoute *gin.RouterGroup) {
	superRoute.GET("/", controller.HomePageHandler())

	superRoute.GET("/nigga", func(c *gin.Context) {
		Title := "nigga"
		render.RenderHtml(c, http.StatusAccepted, "base.html", Title)
	})

	// API for video streaming
	superRoute.GET("/stream/:filename", controller.StreamHandler())

	// frontend url -- The request responds to a url matching: /video?filename=example.mp4
	superRoute.GET("/video", controller.VideoPageHandler())
}
