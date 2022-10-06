package routes

import (
	"golang_basic_gorm_gin/auth"
	"golang_basic_gorm_gin/config"
	"golang_basic_gorm_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messages": "Bad Request",
			"error":    err.Error(),
		})

		c.Abort()
		return
	}

	// hashPassword
	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": "Internal Server Error",
			"error":    err.Error(),
		})

		c.Abort()
		return
	}

	// insert data to tbl user
	insertUser := config.DB.Create(&user)
	if insertUser.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": "Internal Server Error",
			"error":    insertUser.Error.Error(),
		})

		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"UserID":   user.ID,
		"Email":    user.Email,
		"username": user.Username,
	})
}

func GenerateToken(c *gin.Context) {
	var request TokenRequest
	var user models.User

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messages": "Bad Request",
			"error":    err.Error(),
		})

		c.Abort()
		return
	}

	// check email
	checkEmail := config.DB.Where("email = ?", request.Email).First(&user)
	if checkEmail.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"messages": "Data Not FOund",
			"error":    checkEmail.Error.Error(),
		})

		c.Abort()
		return
	}

	// Check password
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"messages": "Invalid Credential",
			"error":    credentialError.Error(),
		})

		c.Abort()
		return
	}

	// generate token
	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": "Internal Server Error",
			"error":    err.Error(),
		})

		c.Abort()
		return
	}

	// response token
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
