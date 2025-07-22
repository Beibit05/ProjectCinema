package tests

import (
	"ProjectCinema/config"
	"ProjectCinema/handlers"
	"ProjectCinema/models"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Таза тест базасын дайындау
func initSessionTestDB() {
	dsn := "host=localhost user=postgres password=2005b dbname=cinema_db_test port=5432 sslmode=disable TimeZone=Asia/Almaty"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to test DB:", err)
	}
	config.DB = db

	err = db.Migrator().DropTable(&models.Session{})
	if err != nil {
		log.Fatal("❌ Failed to drop sessions table:", err)
	}

	err = db.AutoMigrate(&models.Session{})
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	// Тест мәліметін қосу
	db.Create(&models.Session{
		FilmID:    1,
		CinemaID:  1,
		StartTime: time.Now().Add(2 * time.Hour),
		HallName:  "Blue Hall",
	})
}

// Барлық маршруттарды бір жерде тіркеу
func setupSessionRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.POST("/sessions", handlers.CreateSession)
	r.GET("/sessions", handlers.GetSessions)
	r.GET("/sessions/:id", handlers.GetSessionByID)
	r.PUT("/sessions/:id", handlers.UpdateSession)
	r.DELETE("/sessions/:id", handlers.DeleteSession)

	return r
}

func TestSessionHandlers(t *testing.T) {
	// Әр тесттің алдында база тазаланып, дайын болады
	initSessionTestDB()
	router := setupSessionRouter()

	t.Run("Create Session", func(t *testing.T) {
		session := models.Session{
			FilmID:    2,
			CinemaID:  3,
			StartTime: time.Now().Add(4 * time.Hour),
			HallName:  "Red Hall",
		}
		body, _ := json.Marshal(session)
		req, _ := http.NewRequest("POST", "/sessions", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)
	})

	t.Run("Get All Sessions", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/sessions", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("Get Session By ID", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/sessions/1", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("Update Session", func(t *testing.T) {
		update := models.Session{
			FilmID:    1,
			CinemaID:  1,
			StartTime: time.Now().Add(5 * time.Hour),
			HallName:  "Updated Hall",
		}
		body, _ := json.Marshal(update)
		req, _ := http.NewRequest("PUT", "/sessions/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("Delete Session", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/sessions/1", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})
}
