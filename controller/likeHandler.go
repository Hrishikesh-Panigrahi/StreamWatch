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
			utils.DeleteLikeFromDB(c, user, video, like)
		} else {
			fmt.Println("User is liking the video.")
			utils.RemoveDislike(c, user, video)
			utils.AddLike(c, video.ID, user.ID)
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
			utils.DeleteDislikeFromDB(c, user, video, dislike)
		} else {
			fmt.Println("User is disliking the video.")
			utils.RemoveLike(c, user, video)
			utils.AddDislike(c, video.ID, user.ID)
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