package handlers

import (
	"ProjectCinema/models"
	"ProjectCinema/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCinema(c *gin.Context) {
	var cinema models.Cinema
	if err := c.ShouldBindJSON(&cinema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateCinema(&cinema); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cinema"})
		return
	}
	c.JSON(http.StatusCreated, cinema)
}

func GetCinemas(c *gin.Context) {
	cinemas, err := services.GetAllCinemas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cinemas"})
		return
	}
	c.JSON(http.StatusOK, cinemas)
}

func GetCinemaByID(c *gin.Context) {
	id := c.Param("id")
	cinema, err := services.GetCinemaByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cinema not found"})
		return
	}
	c.JSON(http.StatusOK, cinema)
}

func UpdateCinema(c *gin.Context) {
	id := c.Param("id")
	cinema, err := services.GetCinemaByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cinema not found"})
		return
	}
	if err := c.ShouldBindJSON(&cinema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.UpdateCinema(cinema); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, cinema)
}

func DeleteCinema(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteCinema(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cinema deleted"})
}
