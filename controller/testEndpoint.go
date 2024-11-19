package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func TestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("TestHandler endpoint hit")
		
		var payload struct {
			Action    string `json:"action"`
			LikeCount int    `json:"likeCount"`
		}

		// Bind JSON data sent in the request body
		if err := c.ShouldBindJSON(&payload); err != nil {
			fmt.Printf("Error binding JSON: %v\n", err)
			c.JSON(400, gin.H{"error": "Invalid payload"})
			return
		}

		// Log the received data
		fmt.Printf("Action: %s, LikeCount: %d\n", payload.Action, payload.LikeCount)

		c.JSON(200, gin.H{
			"message": "Test endpoint",
		})
	}
}
