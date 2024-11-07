package middleware

import (
	"fmt"
	"net/http"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/dbConnector"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
	"github.com/gin-gonic/gin"
)

func EmailVerification() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieuser, exists := c.Get("user")

		if !exists {
			render.RenderError(c, http.StatusUnauthorized, "User not logged in. Please login to upload video.")
			return
		}

		userID := cookieuser.(models.User).ID

		fmt.Println(userID)

		var user models.User
		if err := dbConnector.DB.First(&user, userID).Error; err != nil {
			fmt.Printf("Error retrieving User: %v\n", err)

			render.RenderError(c, http.StatusInternalServerError, "An error occurred while fetching the user. Please try again later.")
			return
		}

		if !user.Is_verified {
			render.RenderError(c, http.StatusUnauthorized, "User not verified. Please verify your email.")
			return
		}

		c.Next()
	}
}
