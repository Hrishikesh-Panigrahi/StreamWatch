package utils

import (
	"net/http"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
	"github.com/gin-gonic/gin"
)

func LoadUserFromCache(c *gin.Context) uint {
	user, exists := c.Get("user")
	if !exists {
		render.RenderError(c, http.StatusUnauthorized, "User not logged in.")
		return 0
	}

	return user.(models.User).ID
}
