package controller

import (
	"fmt"
	"net/http"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/dbConnector"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
	"github.com/gin-gonic/gin"
)

func LikeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		VideoUUID := c.Param("UUID")

		cookieuser, exists := c.Get("user")

		if !exists {
			render.RenderError(c, http.StatusUnauthorized, "User not logged in. Please login to like the video.")
			return
		}

		userID := cookieuser.(models.User).ID

		var user models.User
		if err := dbConnector.DB.First(&user, userID).Error; err != nil {
			fmt.Printf("Error retrieving User: %v\n", err)

			render.RenderError(c, http.StatusInternalServerError, "An error occurred while fetching the user. Please try again later.")
			return
		}

		var video models.Video
		if err := dbConnector.DB.Where("uuid = ?", VideoUUID).First(&video).Error; err != nil {
			fmt.Printf("Error retrieving Video: %v\n", err)

			render.RenderError(c, http.StatusInternalServerError, "An error occurred while fetching the video. Please try again later.")
			return
		}

		like := models.Likes{
			VideoId: video.ID,
			UserId:  user.ID,
		}

		if err := dbConnector.DB.Create(&like).Error; err != nil {
			fmt.Printf("Error creating like: %v\n", err)
			render.RenderError(c, http.StatusInternalServerError, "Failed to like the video. Please try again later.")
			return
		}

	}
}

func GetLikeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		VideoUUID := c.Param("UUID")

		var video models.Video
		if err := dbConnector.DB.Where("uuid = ?", VideoUUID).First(&video).Error; err != nil {
			fmt.Printf("Error retrieving Video: %v\n", err)

			render.RenderError(c, http.StatusInternalServerError, "An error occurred while fetching the video. Please try again later.")
			return
		}

		var likes []models.Likes
		if err := dbConnector.DB.Where("video_id = ?", video.ID).Find(&likes).Error; err != nil {
			fmt.Printf("Error retrieving Likes: %v\n", err)

			render.RenderError(c, http.StatusInternalServerError, "An error occurred while fetching the likes. Please try again later.")
			return
		}

	}
}
