package controller

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

func VideoPageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		filename := c.Param("filename")
		tmpl, err := template.ParseFiles("./templates/index.html")

		if err != nil {
			c.String(http.StatusInternalServerError, "Error loading template file", err)
			return
		}

		data := struct {
			Filename string
		}{
			Filename: filename,
		}

		c.Header("Content-Type", "text/html")

		if err := tmpl.Execute(c.Writer, data); err != nil {
			log.Printf("Error executing template: %v", err)
			c.String(http.StatusInternalServerError, "Template execution error: %v", err)
		}
	}
}
