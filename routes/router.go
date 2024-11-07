package routes

import (
	"github.com/Hrishikesh-Panigrahi/StreamWatch/controller"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(superRoute *gin.RouterGroup) {
	superRoute.GET("/", controller.HomePageHandler())

	superRoute.GET("/login", controller.LoginPageHandler())
	superRoute.POST("/login", controller.LoginHandler())

	superRoute.GET("/register", controller.RegisterPageHandler())
	superRoute.POST("/register", controller.RegisterHandler())

	superRoute.GET("/logout", controller.LogoutHandler())

	superRoute.POST("create/video", middleware.AuthMiddleware, middleware.RateLimitMiddleware(), controller.CreateVideoHandler())
	superRoute.GET("create/video", middleware.AuthMiddleware, middleware.EmailVerification(), controller.CreateVideoPageHandler())

	// API for video streaming
	superRoute.GET("/stream/:UUID", controller.StreamHandler())
	// frontend url -- The request responds to a url matching: /video?UUID=xxxx-xxxx-xxxx-xxxx
	superRoute.GET("/video", controller.VideoPageHandler())

}
