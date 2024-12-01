package routes

import (
	"github.com/Hrishikesh-Panigrahi/StreamWatch/controller"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

type AllRoutes []Route

var routes = AllRoutes{
	Route{
		"GET",
		"/",
		controller.HomePageHandler(),
	},
	Route{
		"GET",
		"/login",
		controller.LoginPageHandler(),
	},
	Route{
		"POST",
		"/login",
		controller.LoginHandler(),
	},
	Route{
		"GET",
		"/register",
		controller.RegisterPageHandler(),
	},
	Route{
		"POST",
		"/register",
		controller.RegisterHandler(),
	},
	Route{
		"GET",
		"/logout",
		controller.LogoutHandler(),
	},
	Route{
		"POST",
		"/create/video",
		controller.CreateVideoHandler(),
	},
	Route{
		"GET",
		"/create/video",
		controller.CreateVideoPageHandler(),
	},
	Route{
		"POST",
		"/video/:UUID/like",

		controller.LikeHandler(),
	},
	Route{
		"GET",
		"/video/:UUID/getlike",
		controller.GetLikeHandler(),
	},
	Route{
		"POST",
		"/video/:UUID/dislike",
		controller.DislikeHandler(),
	},
	Route{
		"GET",

		"/video/:UUID/getdislike",
		controller.GetDislikeHandler(),
	},
	Route{
		"POST",
		"/video/:UUID/watchlog",
		controller.WatchLogHandler(),
	},
	Route{
		"GET",
		"/video/:UUID/getwatchlog",
		controller.GetWatchLogHandler(),
	},
	Route{
		"GET",
		"/stream/:UUID",
		controller.StreamHandler(),
	},
	Route{
		"GET",
		"/video",
		controller.VideoPageHandler(),
	},
	Route{
		"GET",
		"/trending-tags",
		controller.TrendingTagsHandler(),
	},
}

func GetRoutes() AllRoutes {
	return routes
}
