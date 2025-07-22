package tests

import (
	"ProjectCinema/config"
	"ProjectCinema/handlers"
	"ProjectCinema/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SetupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.GET("/films", handlers.GetAllFilms)
	r.POST("/films", handlers.CreateFilms)
	r.GET("/films/:id", handlers.GetById)
	r.PUT("/films/:id", handlers.UpdateFilms)
	r.DELETE("/films/:id", handlers.DeleteFilms)

	return r
}

var DB *gorm.DB

func initDBTest() {
	dns := "host=localhost user=postgres password=2005b dbname=cinema_db_test port=5432 sslmode=disable TimeZone=Asia/Almaty"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Failed to connect to database: ", err)
	}
	DB = db
	config.DB = db

	err = db.Migrator().DropTable(&models.Film{})
	if err != nil {
		log.Fatal("Failed to drop table: ", err)
	}
	err = db.AutoMigrate(&models.Film{})
	if err != nil {
		log.Fatal("Migration failed", err)
	}

	db.Create(&models.Film{
		Title: "Inception", GenreID: 1, Description: "Dream movie", DirectorID: 1,
		DurationMinutes: 148, ReleaseYear: 2010,
	})
	db.Create(&models.Film{
		Title: "The Matrix", GenreID: 2, Description: "Virtual world", DirectorID: 2,
		DurationMinutes: 136, ReleaseYear: 1999,
	})

	fmt.Println("Database prepared for test")
}

func TestGetFilms(t *testing.T) {
	gin.SetMode(gin.TestMode)
	initDBTest()
	router := SetupTestRouter()

	req, _ := http.NewRequest("GET", "/films?page=1&limit=10", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Inception")
	assert.Contains(t, resp.Body.String(), "The Matrix")
}
func TestCreateFilms(t *testing.T) {
	gin.SetMode(gin.TestMode)
	initDBTest()
	router := SetupTestRouter()

	newFilmsJSON := `[
		{"title": "Interstellar", "genre_id": 3, "description": "Space movie", "director_id": 1, "duration_minutes": 169, "release_year": 2014},
		{"title": "Fight Club", "genre_id": 4, "description": "Mind-blowing", "director_id": 2, "duration_minutes": 139, "release_year": 1999}
	]`

	req, _ := http.NewRequest("POST", "/films",
		strings.NewReader(newFilmsJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Contains(t, resp.Body.String(), "Interstellar")
	assert.Contains(t, resp.Body.String(), "Fight Club")
}

func TestGetFilmById(t *testing.T) {
	initDBTest()
	router := SetupTestRouter()

	req, _ := http.NewRequest("GET", "/films/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Inception")
}

func TestUpdateFilm(t *testing.T) {
	gin.SetMode(gin.TestMode)
	initDBTest()

	router := gin.Default()
	router.PUT("/films/:id", handlers.UpdateFilms)

	updateJSON := `{
		"id": 1,
		"title": "Inception Updated",
		"genre_id": 5,
		"description": "Updated Description",
		"director_id": 3,
		"duration_minutes": 150,
		"release_year": 2011
	}`

	req, _ := http.NewRequest("PUT", "/films/1", strings.NewReader(updateJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Film updated Successfully!")
	assert.Contains(t, resp.Body.String(), "Inception Updated")
}

func TestDeleteFilm(t *testing.T) {
	initDBTest()
	router := SetupTestRouter()

	// Delete film
	req, _ := http.NewRequest("DELETE", "/films/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Film deleted successfully")

	// Try to get deleted film
	getReq, _ := http.NewRequest("GET", "/films/1", nil)
	getResp := httptest.NewRecorder()
	router.ServeHTTP(getResp, getReq)

	assert.Equal(t, http.StatusInternalServerError, getResp.Code)
	assert.Contains(t, getResp.Body.String(), "Film not found")
}
