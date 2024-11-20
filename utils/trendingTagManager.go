package utils

import (
	"fmt"
	"net/http"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/dbConnector"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
	"github.com/gin-gonic/gin"
)

func TrendingTagManager(c *gin.Context, tag string) {
	var trendingTag models.TrendingTags

	if err := dbConnector.DB.Where("tag = ?", tag).First(&trendingTag).Error; err != nil {
		trendingTag = models.TrendingTags{
			Tag:         tag,
			Usage_count: 1,
		}
		if err := dbConnector.DB.Create(&trendingTag).Error; err != nil {
			fmt.Printf("Error creating trending tag: %v\n", err)
			render.RenderError(c, http.StatusInternalServerError, "An error occurred while creating the trending tag. Please try again later.")
			return
		}
	} else {
		trendingTag.Usage_count++
		if err := dbConnector.DB.Save(&trendingTag).Error; err != nil {
			fmt.Printf("Error updating trending tag: %v\n", err)
			render.RenderError(c, http.StatusInternalServerError, "An error occurred while updating the trending tag. Please try again later.")
			return
		}
	}
}
