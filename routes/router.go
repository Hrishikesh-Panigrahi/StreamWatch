package routes

import (
	"github.com/Hrishikesh-Panigrahi/StreamWatch/controller"
	"github.com/gin-gonic/gin"
)

func Routes(superRoute *gin.RouterGroup) {

	superRoute.GET("/login", controller.LoginPageHandler())
	superRoute.POST("/login", controller.LoginHandler())

	superRoute.GET("/register", controller.RegisterPageHandler())
	superRoute.POST("/register", controller.RegisterHandler())

	superRoute.GET("/", controller.HomePageHandler())

	superRoute.POST("create/video", controller.CreateVideo())

	// API for video streaming
	superRoute.GET("/stream/:UUID", controller.StreamHandler())

	// frontend url -- The request responds to a url matching: /video?UUID=xxxx-xxxx-xxxx-xxxx
	superRoute.GET("/video", controller.VideoPageHandler())
}
