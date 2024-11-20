package controller

import (
	"net/http"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/utils"
	"github.com/gin-gonic/gin"
)

func TrendingTagsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		trendingTags := utils.GetTrendingTags()

		type Data struct {
			Title        string
			Message      string
			TrendingTags []models.TrendingTags
		}
		data := Data{Title: "trendingTags", Message: "this is index", TrendingTags: trendingTags}

		render.RenderHtml(c, http.StatusOK, "base.html", data)
	}
}