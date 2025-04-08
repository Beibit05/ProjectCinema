package handlers

import (
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

	var filterFilms []models.Film
	for _, film := range films {
		if (category == "" || strconv.Itoa(film.Genre) == category) &&
			(author == "" || strconv.Itoa(film.Genre) == author) {
			filterFilms = append(filterFilms, film)
		}
	}

	start := (page - 1) * limit
	end := page + limit

	if page >= len(filterFilms) {
		c.JSON(200, []models.Film{})
		return
	}

	if limit > len(filterFilms) {
		limit = len(filterFilms)
	}

	c.JSON(200, filterFilms[start:end])
}

func CreateFilms(c *gin.Context) {
	var newFilms []models.Film
	if err := c.ShouldBindJSON(&newFilms); err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	startId := len(films) + 1
	for i := range newFilms {
		newFilms[i].Id = startId + i
		films = append(films, newFilms[i])
	}
	c.JSON(200, newFilms)
}

func GetById(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid id films"})
		return
	}
	var oneFilms models.Film
	found := false
	mu.Lock()
	for _, film := range films {
		if film.Id == id {
			oneFilms = film
			found = true
			break
		}
	}
	mu.Unlock()

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Film not found "})
		return
	}
	c.JSON(200, gin.H{"This film": oneFilms})
}
func UpdateFilms(c *gin.Context) {
	idParam := c.Param("id")
	idB, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film id"})
		return
	}
	var oldFilm models.Film
	found := -1
	mu.Lock()
	for i, film := range films {
		if film.Id == idB {
			oldFilm = film
			found = i
			break
		}
	}
	mu.Unlock()
	if found == -1 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Film not found"})
		return
	}
	var updateFilm models.Film

	if err := c.ShouldBindJSON(&updateFilm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	mu.Lock()
	updateFilm.Id = films[found].Id
	films[found] = updateFilm
	mu.Unlock()
	c.JSON(http.StatusOK, gin.H{"message": "Book updated Successfully!",
		"OldBook":    oldFilm,
		"UpdateBook": films[found]})
}

func DeleteFilms(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Book id"})
	}
	var newFilms []models.Film
	found := false
	mu.Lock()
	defer mu.Unlock()
	for _, film := range films {
		if film.Id == id {
			found = true
			continue
		}
		newFilms = append(newFilms, film)
	}
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Book not found"})
		return
	}
	films = newFilms

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
