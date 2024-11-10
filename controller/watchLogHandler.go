package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/dbConnector"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/utils"
	"github.com/gin-gonic/gin"
)

func WatchLogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get video UUID from URL parameter
		VideoUUID := c.Param("UUID")

		// Get watch duration from form data
		durationStr := c.PostForm("duration")

		fmt.Println("Duration: ", durationStr)

		duration, err := strconv.ParseFloat(durationStr, 64)
		if err != nil {
			render.RenderError(c, http.StatusBadRequest, "Invalid watch duration.")
			return
		}

		userID := utils.LoadUserFromCache(c)

		video := utils.GetVideoByUUID(c, VideoUUID, "Sorry your Video was not found.")

		var watchlog models.WatchLog
		result := dbConnector.DB.Where("video_id = ? AND user_id = ?", video.ID, userID).First(&watchlog)

		if result.Error == nil {
			fmt.Println("user has watched some part of the video. Updating the watch duration.")
			watchlog.Watch_duration = time.Duration(duration) * time.Second
			if err := dbConnector.DB.Where("video_id = ? AND user_id = ?", video.ID, userID).Save(&watchlog).Error; err != nil {
				render.RenderError(c, http.StatusInternalServerError, "Failed to update watch duration.")
				return
			}
		} else {
			fmt.Println("user has not watched the video. Creating a new watch log.")
			newWatchLog := models.WatchLog{
				VideoId:        video.ID,
				UserId:         userID,
				Watch_duration: time.Duration(duration) * time.Second, // Store in seconds
			}
			if err := dbConnector.DB.Create(&newWatchLog).Error; err != nil {
				render.RenderError(c, http.StatusInternalServerError, "Failed to log watch duration.")
				return
			}
		}

		c.Status(http.StatusOK)
	}
}

func GetWatchLogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		VideoUUID := c.Param("UUID")
		video := utils.GetVideoByUUID(c, VideoUUID, "Sorry your Video was not found.")

		userID := utils.LoadUserFromCache(c)

		var watchlog models.WatchLog
		result := dbConnector.DB.Where("video_id = ? AND user_id = ?", video.ID, userID).First(&watchlog)

		if result.Error != nil {
			// render.RenderError(c, http.StatusNotFound, "Watch log not found.")
			return
		}

		duration := watchlog.Watch_duration

		fmt.Println("Watch Duration: ", duration)
		c.String(http.StatusOK, "%v", duration)
	}
}
