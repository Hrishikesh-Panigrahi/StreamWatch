package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/dbConnector"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
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

		// Retrieve user from context
		user, exists := c.Get("user")
		if !exists {
			render.RenderError(c, http.StatusUnauthorized, "User not logged in.")
			return
		}

		// Retrieve video based on UUID
		var video models.Video
		if err := dbConnector.DB.Where("uuid = ?", VideoUUID).First(&video).Error; err != nil {
			render.RenderError(c, http.StatusNotFound, "Video not found.")
			return
		}

		// // Log the watch duration
		// watchLog := models.WatchLog{
		// 	VideoId:        video.ID,
		// 	UserId:         user.(models.User).ID,
		// 	Watch_duration: time.Duration(duration) * time.Second, // Store in seconds
		// }
		// if err := dbConnector.DB.Create(&watchLog).Error; err != nil {
		// 	render.RenderError(c, http.StatusInternalServerError, "Failed to log watch duration.")
		// 	return
		// }

		fmt.Println("Watch duration: ", time.Duration(duration)*time.Second)
		fmt.Println("User: ", user.(models.User).ID)

		c.Status(http.StatusOK)
	}
}
