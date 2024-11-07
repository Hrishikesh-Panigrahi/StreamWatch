package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/dbConnector"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware is a middleware to check if the user is authenticated
// it checks if the token is valid and not expired if valid it sets the user in the context
// if not valid it returns a 401
func AuthMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("token")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token missing"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Printf("Token parsing error: %v\n", err) // Avoid fatal errors that stop the server
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		userID, ok := claims["sub"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token payload"})
			c.Abort()
			return
		}

		var user models.User

		if err := dbConnector.DB.First(&user, uint(userID)).Error; err != nil {
			log.Printf("User retrieval error: %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		c.Abort()
	}

}
