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
	category := c.Query("cate")
	author := c.Query("author")
	var films []models.Film
	query := config.DB.Model(&models.Film{})

	// Фильтрация
	if category != "" {
		categoryInt, _ := strconv.Atoi(category)
		query = query.Where("genre = ?", categoryInt)
	}
	if author != "" {
		authorInt, _ := strconv.Atoi(author)
		query = query.Where("genre = ?", authorInt)
	}

	// Пагинация
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&films).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//var filterFilms []models.Film
	//for _, film := range films {
	//	if (category == "" || strconv.Itoa(film.Genre) == category) &&
	//		(author == "" || strconv.Itoa(film.Genre) == author) {
	//		filterFilms = append(filterFilms, film)
	//	}
	//}
	//if err := config.DB.Find(&filterFilms).Error; err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	//	return
	//}
	//
	//start := (page - 1) * limit
	//end := page + limit
	//
	//if page >= len(filterFilms) {
	//	c.JSON(200, []models.Film{})
	//	return
	//}
	//
	//if limit > len(filterFilms) {
	//	limit = len(filterFilms)
	//}

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
	film.Id = updateFilm.Id
	film.Title = updateFilm.Title
	film.Genre = updateFilm.Genre
	film.Description = updateFilm.Description
	film.Duration = updateFilm.Duration
	film.VideoURL = updateFilm.VideoURL

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
