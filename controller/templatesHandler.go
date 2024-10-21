package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VideoPageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		filename := c.Param("filename")

		if filename == "" {
			c.HTML(http.StatusNotFound, "Video not found", nil)
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Filename": filename,
		})

		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render template"})
		// 	// Optionally log the error
		// 	// log.Printf("Error rendering template: %v", err)
		// 	return
		// }
	}
}
