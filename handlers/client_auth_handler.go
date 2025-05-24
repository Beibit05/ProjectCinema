package handlers

import (
	"ProjectCinema/utils"
	"github.com/gin-gonic/gin"

	"net/http"
)

func RegisterUser(c *gin.Context) {
	var body map[string]interface{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	resp, err := utils.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post("http://user-service:8081/users/register")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call user services"})
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}

func LoginUser(c *gin.Context) {
	var body map[string]interface{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	resp, err := utils.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post("http://user-service:8081/users/login")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call user services"})
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}
