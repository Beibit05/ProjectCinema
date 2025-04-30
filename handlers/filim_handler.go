package handlers

import (
	"ProjectCinema/config"
	"ProjectCinema/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync"
)

var films []models.Film

var mu sync.Mutex

func GetAllFilms(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	genre := c.Query("genre")
	//author := c.Query("author")
	var films []models.Film
	query := config.DB.Model(&models.Film{})

	if genre != "" {
		genreInt, _ := strconv.Atoi(genre)
		query = query.Where("genre = ?", genreInt)
	}
	//if author != "" {
	//	authorInt, _ := strconv.Atoi(author)
	//	query = query.Where("genre = ?", authorInt)
	//}

	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&films).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, films)
}

func CreateFilms(c *gin.Context) {
	var newFilms []models.Film
	if err := c.ShouldBindJSON(&newFilms); err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}

	if err := config.DB.Create(&newFilms).Error; err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newFilms)
}

func GetById(c *gin.Context) {
	var oneFilms models.Film
	paramId := c.Param("id")

	if err := config.DB.First(&oneFilms, paramId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Film not found"})
		return
	}
	c.JSON(200, gin.H{"This film": oneFilms})
}
func UpdateFilms(c *gin.Context) {
	idParam := c.Param("id")
	var oldFilm models.Film
	if err := config.DB.First(&oldFilm, idParam).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Film not found"})
		return
	}
	var film models.Film
	if err := config.DB.First(&film, idParam).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Film not found"})
		return
	}
	var updateFilm models.Film
	if err := c.ShouldBindJSON(&updateFilm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	film.ID = updateFilm.ID
	film.Title = updateFilm.Title
	film.GenreID = updateFilm.GenreID
	film.Description = updateFilm.Description
	film.DurationMinutes = updateFilm.DurationMinutes
	film.DirectorID = updateFilm.DirectorID
	film.ReleaseYear = updateFilm.ReleaseYear

	if err := config.DB.Save(&film).Error; err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Film updated Successfully!",
		"OldFilm":    oldFilm,
		"UpdateFilm": film})
}

func DeleteFilms(c *gin.Context) {
	idParam := c.Param("id")
	if err := config.DB.Delete(&models.Film{}, idParam).Error; err != nil {
		c.JSON(500, gin.H{"Error": "Film not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Film deleted successfully"})
}
