package handlers

import (
	"ProjectCinema/models"
	"ProjectCinema/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSession(c *gin.Context) {
	var session models.Session
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateSession(&session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session"})
		return
	}
	c.JSON(http.StatusCreated, session)
}

func GetSessions(c *gin.Context) {
	sessions, err := services.GetAllSessions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sessions"})
		return
	}
	c.JSON(http.StatusOK, sessions)
}

func GetSessionByID(c *gin.Context) {
	id := c.Param("id")
	session, err := services.GetSessionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}
	c.JSON(http.StatusOK, session)
}

func UpdateSession(c *gin.Context) {
	id := c.Param("id")
	session, err := services.GetSessionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.UpdateSession(session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, session)
}

func DeleteSession(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteSession(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Session deleted"})
}
