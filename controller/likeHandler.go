package controller

import (
	"fmt"
	"net/http"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/dbConnector"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/utils"
	"github.com/gin-gonic/gin"
)

func LikeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		VideoUUID := c.Param("UUID")

		user := utils.GetUserFromCache(c)

		video := utils.GetVideoByUUID(c, VideoUUID, "An error occurred while fetching the video. Please try again later.")

		var like models.Likes
		result := dbConnector.DB.Where("video_id = ? AND user_id = ?", video.ID, user.ID).First(&like)
		if result.Error == nil {
			fmt.Println("User already liked the video. Deleting the like.")
			if err := dbConnector.DB.Unscoped().Where("video_id = ? AND user_id = ?", video.ID, user.ID).Delete(&like).Error; err != nil {
				fmt.Printf("Error deleting like: %v\n", err)
				render.RenderError(c, http.StatusInternalServerError, "Failed to delete the like for the video. Please try again later.")
				return
			}
		} else {
			newLike := models.Likes{
				VideoId: video.ID,
				UserId:  user.ID,
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
		video := utils.GetVideoByUUID(c, VideoUUID, "An error occurred while fetching the video. Please try again later.")

		var likeCount int64
		if err := dbConnector.DB.Model(&models.Likes{}).Where("video_id = ?", video.ID).Count(&likeCount).Error; err != nil {
			fmt.Printf("Error retrieving like count: %v\n", err)
			render.RenderError(c, http.StatusInternalServerError, "An error occurred while fetching the like count. Please try again later.")
			return
		}

		c.String(http.StatusOK, "%d", likeCount)
	}
}

func DislikeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		VideoUUID := c.Param("UUID")

		user := utils.GetUserFromCache(c)

		video := utils.GetVideoByUUID(c, VideoUUID, "An error occurred while fetching the video. Please try again later.")

		var dislike models.Dislikes
		result := dbConnector.DB.Where("video_id = ? AND user_id = ?", video.ID, user.ID).First(&dislike)
		if result.Error == nil {
			fmt.Println("User already disliked the video. Deleting the dislike.")
			if err := dbConnector.DB.Unscoped().Where("video_id = ? AND user_id = ?", video.ID, user.ID).Delete(&dislike).Error; err != nil {
				fmt.Printf("Error deleting dislike: %v\n", err)
				render.RenderError(c, http.StatusInternalServerError, "Failed to delete the dislike for the video. Please try again later.")
				return
			}
		} else {
			fmt.Println("User has disliked the video. Disliking the video.........")
			var like models.Likes
			result := dbConnector.DB.Where("video_id = ? AND user_id = ?", video.ID, user.ID).First(&like)
			if result.Error == nil {
				fmt.Println("User already liked the video. Deleting the like, as user is disliking the video.")
				if err := dbConnector.DB.Unscoped().Where("video_id = ? AND user_id = ?", video.ID, user.ID).Delete(&like).Error; err != nil {
					fmt.Printf("Error deleting like: %v\n", err)
					render.RenderError(c, http.StatusInternalServerError, "Failed to delete the like for the video. Please try again later.")
					return
				}else{
					fmt.Println("Deleted the like for the video.")
				}
			}
			newDislike := models.Dislikes{
				VideoId:    video.ID,
				UserId:     user.ID,
				DislikedAt: models.GetCurrentTime(),
			}
			if err := dbConnector.DB.Create(&newDislike).Error; err != nil {
				fmt.Printf("Error creating dislike: %v\n", err)
				render.RenderError(c, http.StatusInternalServerError, "Failed to dislike the video. Please try again later.")
				return
			}
		}

	}
}

func GetDislikeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		VideoUUID := c.Param("UUID")
		video := utils.GetVideoByUUID(c, VideoUUID, "An error occurred while fetching the video. Please try again later.")

		var dislikeCount int64
		if err := dbConnector.DB.Model(&models.Dislikes{}).Where("video_id = ?", video.ID).Count(&dislikeCount).Error; err != nil {
			fmt.Printf("Error retrieving dislike count: %v\n", err)
			render.RenderError(c, http.StatusInternalServerError, "An error occurred while fetching the dislike count. Please try again later.")
			return
		}

		c.String(http.StatusOK, "%d", dislikeCount)
	}
}
