package routes

import (
	"ProjectCinema/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	films := r.Group("/films")
	{
		films.GET("/", handlers.GetAllFilms)
		films.POST("/", handlers.CreateFilms)
		films.GET("/:id", handlers.GetById)
		films.PUT("/:id", handlers.UpdateFilms)
		films.DELETE("/:id", handlers.DeleteFilms)
	}

	genre := r.Group("/genres")
	{
		genre.GET("/", handlers.GetAllGenres)
		genre.POST("/", handlers.CreateGenre)
		genre.GET("/:id", handlers.GetGenreByID)
		genre.PUT("/:id", handlers.UpdateGenre)
		genre.DELETE("/:id", handlers.DeleteGenre)
	}
}
