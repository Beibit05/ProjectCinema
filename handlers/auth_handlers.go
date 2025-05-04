package handlers

import (
	"ProjectCinema/config"
	"ProjectCinema/models"
	"ProjectCinema/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	newUser.Password = string(hashPass)

	if err := config.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not register"})
		return
	}

	c.JSON(200, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var user models.User
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"Error": "Invalid input"})
		return
	}
	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "User not found"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Incorrect password"})
		return
	}

	tokenString, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//func LoginHandler(c *gin.Context) {
//	var loginData LoginPayload
//	if err := c.ShouldBindJSON(&loginData); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
//		return
//	}
//
//	client := resty.New()
//	resp, err := client.R().
//		SetBody(loginData).
//		Post("http://localhost:8081/login") // Бұл — auth-service логикасы
//
//	if err != nil || resp.StatusCode() != http.StatusOK {
//		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
//		return
//	}
//
//	c.Data(resp.StatusCode(), "application/json", resp.Body())
//}
