package utils

import (
	"fmt"
	"net/http"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/dbConnector"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
	"github.com/gin-gonic/gin"
)

func DeleteDislikeFromDB(c *gin.Context, user models.User, video models.Video, dislike models.Dislikes) {
	if err := dbConnector.DB.Unscoped().Where("video_id = ? AND user_id = ?", video.ID, user.ID).Delete(&dislike).Error; err != nil {
		fmt.Printf("Error deleting Dislike: %v\n", err)
		render.RenderError(c, http.StatusInternalServerError, "Failed to delete the Dislike for the video. Please try again later.")
		return
	} else {
		fmt.Println("Deleted the dislike for the video.")
	}
}

func RemoveDislike(c *gin.Context, user models.User, video models.Video) {
	var dislike models.Dislikes
	result := dbConnector.DB.Where("video_id = ? AND user_id = ?", video.ID, user.ID).First(&dislike)
	if result.Error == nil {
		fmt.Println("User liked the video. Deleting the dislike, as user is liking the video.")
		DeleteDislikeFromDB(c, user, video, dislike)
	}
}

func AddDislike(c *gin.Context, videoID uint, userID uint) {
	newDislike := models.Dislikes{
		VideoId:    videoID,
		UserId:     userID,
		DislikedAt: models.GetCurrentTime(),
	}
	if err := dbConnector.DB.Create(&newDislike).Error; err != nil {
		fmt.Printf("Error creating dislike: %v\n", err)
		render.RenderError(c, http.StatusInternalServerError, "Failed to dislike the video. Please try again later.")
		return
	}
}
