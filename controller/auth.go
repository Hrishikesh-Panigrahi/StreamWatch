package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/dbConnector"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// RegisterHandler is a function to register an user
// it creates an user with the email and password
func RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			Email    string
			Password string
		}

		if c.Bind(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to read body",
			})
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to hash password",
			})
			return
		}

		user := models.User{Email: body.Email, Password: string(hash)}
		result := dbConnector.DB.Create(&user)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to create user",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	}
}

// LoginHandler is a function to login a user it checks if the email and password are correct
// if correct it then generates a jwt token and sets it as a cookie
// if not correct it returns a 401
func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			Email    string
			Password string
		}

		if c.Bind(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to read body",
			})
			return
		}

		// find the user
		var user models.User
		dbConnector.DB.First(&user, "email = ?", body.Email)

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid credentials",
			})
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid pssword",
			})
			return
		}

		// generate a jwt toke
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to generate token",
			})
			return
		}

		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("token", tokenString, 3600*24*30, "", "", false, true)
		c.JSON(http.StatusOK, gin.H{})
	}
}

func LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("token", "", -1, "", "", false, true)
		c.JSON(http.StatusOK, gin.H{})
	}
}
