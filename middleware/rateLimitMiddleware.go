package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var createVideoBucket = NewTokenBucket(10, time.Second) // 10 tokens, 1-second refill

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.FullPath() == "/create/video" && !createVideoBucket.TryConsume() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded. Please wait before trying again."})
			c.Abort()
			return
		}
		fmt.Println("Rate limit passed")
		c.Next()
	}
}
