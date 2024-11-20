package controller

import (
	"github.com/Hrishikesh-Panigrahi/StreamWatch/utils"
	"github.com/gin-gonic/gin"
)

func TrendingTagsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		trendingTags := utils.GetTrendingTags()

		c.JSON(200, trendingTags)
	}
}
