package utils

import (
	"fmt"
	"net/http"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/dbConnector"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
	"github.com/gin-gonic/gin"
)

func DeleteLikeFromDB(c *gin.Context, user models.User, video models.Video, like models.Likes) {
	if err := dbConnector.DB.Unscoped().Where("video_id = ? AND user_id = ?", video.ID, user.ID).Delete(&like).Error; err != nil {
		fmt.Printf("Error deleting like: %v\n", err)
		render.RenderError(c, http.StatusInternalServerError, "Failed to delete the like for the video. Please try again later.")
		return
	} else {
		fmt.Println("Deleted the like for the video.")
	}
}

func RemoveLike(c *gin.Context, user models.User, video models.Video) {
	var like models.Likes
	result := dbConnector.DB.Where("video_id = ? AND user_id = ?", video.ID, user.ID).First(&like)
	if result.Error == nil {
		fmt.Println("User disliked the video. Deleting the like, as user is disliking the video.")
		DeleteLikeFromDB(c, user, video, like)
	}
}

func AddLike(c *gin.Context, videoID uint, userID uint) {
	newLike := models.Likes{
		VideoId: videoID,
		UserId:  userID,
		LikedAt: models.GetCurrentTime(),
	}
	if err := dbConnector.DB.Create(&newLike).Error; err != nil {
		fmt.Printf("Error creating like: %v\n", err)
		render.RenderError(c, http.StatusInternalServerError, "Failed to like the video. Please try again later.")
		return
	}
}
