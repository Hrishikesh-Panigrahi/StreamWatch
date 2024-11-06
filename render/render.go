package render

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

func RenderTempl(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func RenderHtml(c *gin.Context, status int, html string, data interface{}) error {
	c.HTML(status, html, gin.H{
		"data": data,
	})
	return nil
}

func RenderError(c *gin.Context, status int, message string, template ...string) {
	type ErrorData struct {
		Title   string
		Message string
	}

	data := ErrorData{
		Title:   "Error",
		Message: message,
	}

	errorTemplate := "base.html"
	if len(template) > 0 {
		errorTemplate = template[0]
	}

	RenderHtml(c, status, errorTemplate, data)
}

func Redirect(c *gin.Context, url string, code int) {
	c.Render(-1, render.Redirect{
		Code:     code,
		Location: url,
		Request:  c.Request,
	})
}
