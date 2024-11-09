package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/dbConnector"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/models"
	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
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
			render.RenderError(c, http.StatusBadRequest, "Failed to Register the user. Please try again later.")
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		if err != nil {
			render.RenderError(c, http.StatusInternalServerError, "Failed to Hash the password. Please try again later.")
			return
		}

		user := models.User{Email: body.Email, Password: string(hash)}
		result := dbConnector.DB.Create(&user)

		if result.Error != nil {
			render.RenderError(c, http.StatusInternalServerError, "Failed to Register the user. Please try again later.")
			return
		}

		render.Redirect(c, "/login", http.StatusFound)
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
			render.RenderError(c, http.StatusBadRequest, "Failed to Login. Please try again later.")
			return
		}

		// find the user
		var user models.User
		dbConnector.DB.First(&user, "email = ?", body.Email)

		if user.ID == 0 {
			render.RenderError(c, http.StatusUnauthorized, "Invalid Email. Please try again later.")
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

		if err != nil {
			render.RenderError(c, http.StatusUnauthorized, "Invalid Password. Please try again later.")
			return
		}

		// generate a jwt toke
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

		if err != nil {
			render.RenderError(c, http.StatusInternalServerError, "Failed to generate JWT token. Please try again later.")
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