package utils

import (
	"net/http"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/dbConnector"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
	"github.com/gin-gonic/gin"
)

func GetVideoByID(id uint) (models.Video, error) {
	var video models.Video
	if err := dbConnector.DB.First(&video, id).Error; err != nil {
		return video, err
	}
	return video, nil
}

func GetVideoByUUID(c *gin.Context, uuid string, message string) models.Video {
	var video models.Video
	if err := dbConnector.DB.Where("uuid = ?", uuid).First(&video).Error; err != nil {
		render.RenderError(c, http.StatusNotFound, message)
		return video
	}
	return video
}
