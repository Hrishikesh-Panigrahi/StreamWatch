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
		fmt.Println(VideoUUID)

		cookieuser, exists := c.Get("user")
		fmt.Println(cookieuser)

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

		var like models.Likes
		result := dbConnector.DB.Where("video_id = ? AND user_id = ?", video.ID, userID).First(&like)
		if result.Error == nil {
			fmt.Println("User already liked the video. Deleting the like.")
			if err := dbConnector.DB.Unscoped().Where("video_id = ? AND user_id = ?", video.ID, userID).Delete(&like).Error; err != nil {
				fmt.Printf("Error deleting like: %v\n", err)
				render.RenderError(c, http.StatusInternalServerError, "Failed to delete the like for the video. Please try again later.")
				return
			}
		} else {
			newLike := models.Likes{
				VideoId: video.ID,
				UserId:  userID,
				LikedAt: models.GetCurrentTime(),
			}
			if err := dbConnector.DB.Create(&newLike).Error; err != nil {
				fmt.Printf("Error creating like: %v\n", err)
				render.RenderError(c, http.StatusInternalServerError, "Failed to like the video. Please try again later.")
				return
			}
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

		var likeCount int64
		if err := dbConnector.DB.Model(&models.Likes{}).Where("video_id = ?", video.ID).Count(&likeCount).Error; err != nil {
			fmt.Printf("Error retrieving like count: %v\n", err)
			render.RenderError(c, http.StatusInternalServerError, "An error occurred while fetching the like count. Please try again later.")
			return
		}

		c.String(http.StatusOK, "%d", likeCount)
	}
}
