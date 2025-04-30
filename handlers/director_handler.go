package handlers

import (
	"ProjectCinema/config"
	"ProjectCinema/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllDirector(c *gin.Context) {
	var directors []models.Director
	if err := config.DB.Find(&directors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, directors)
}

func GetDirectorByID(c *gin.Context) {
	var director models.Director
	id := c.Param("id")
	if err := config.DB.First(&director, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Director not found"})
		return
	}
	c.JSON(http.StatusOK, director)
}

func CreateDirector(c *gin.Context) {
	var directors []models.Director
	if err := c.ShouldBindJSON(&directors); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&directors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, directors)
}

func UpdateDirector(c *gin.Context) {
	var director models.Director
	id := c.Param("id")
	if err := config.DB.First(&director, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Director not found"})
		return
	}
	if err := c.ShouldBindJSON(&director); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&director)
	c.JSON(http.StatusOK, director)
}

func DeleteDirector(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Director{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Director deleted"})
}
