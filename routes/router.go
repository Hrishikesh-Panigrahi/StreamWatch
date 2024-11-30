package routes

import (
	"github.com/Hrishikesh-Panigrahi/StreamWatch/controller"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/middleware"
	"github.com/gin-gonic/gin"
	shardedmap "github.com/zutto/shardedmap"
)

func Routes(superRoute *gin.RouterGroup) {
	// Initialize the sharded map
	shardmap := shardedmap.NewShardMap(24)

	superRoute.GET("/", CheckEndpointInCache(shardmap), controller.HomePageHandler())

	superRoute.GET("/login", controller.LoginPageHandler())
	superRoute.POST("/login", controller.LoginHandler())

	superRoute.GET("/register", controller.RegisterPageHandler())
	superRoute.POST("/register", controller.RegisterHandler())

	superRoute.GET("/logout", controller.LogoutHandler())

	superRoute.POST("create/video", middleware.AuthMiddleware, middleware.RateLimitMiddleware(), controller.CreateVideoHandler())
	superRoute.GET("create/video", middleware.AuthMiddleware, middleware.EmailVerification(), controller.CreateVideoPageHandler())

	superRoute.POST("/video/:UUID/like", middleware.AuthMiddleware, middleware.RateLimitMiddleware(), controller.LikeHandler())
	superRoute.GET("/video/:UUID/getlike", controller.GetLikeHandler())

	superRoute.POST("/video/:UUID/dislike", middleware.AuthMiddleware, middleware.RateLimitMiddleware(), controller.DislikeHandler())
	superRoute.GET("/video/:UUID/getdislike", controller.GetDislikeHandler())

	superRoute.POST("/video/:UUID/watchlog", middleware.AuthMiddleware, controller.WatchLogHandler())
	superRoute.GET("/video/:UUID/getwatchlog", middleware.AuthMiddleware, controller.GetWatchLogHandler())

	// API for video streaming
	superRoute.GET("/stream/:UUID", controller.StreamHandler())
	// frontend url -- The request responds to a url matching: /video?UUID=xxxx-xxxx-xxxx-xxxx
	superRoute.GET("/video", middleware.AuthMiddleware, controller.VideoPageHandler())

	superRoute.GET("/trending-tags", controller.TrendingTagsHandler())

}

func InitRoutes() *gin.Engine {
	router := gin.Default()
	superRoute := router.Group("/api/v1")
	Routes(superRoute)
	return router
}

func CheckEndpointInCache(shardmap shardedmap.ShardMap) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Endpoint is working",
		})
	}
}
