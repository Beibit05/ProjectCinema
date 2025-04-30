package routes

import (
	"ProjectCinema/handlers"
	"ProjectCinema/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	authorized := r.Group("/")
	authorized.Use(middlewares.Authmiddleware())

	films := authorized.Group("/films")
	{
		films.GET("/", handlers.GetAllFilms)
		films.POST("/", handlers.CreateFilms)
		films.GET("/:id", handlers.GetById)
		films.PUT("/:id", handlers.UpdateFilms)
		films.DELETE("/:id", handlers.DeleteFilms)
	}

	genre := authorized.Group("/genres")
	{
		genre.GET("/", handlers.GetAllGenres)
		genre.POST("/", handlers.CreateGenre)
		genre.GET("/:id", handlers.GetGenreByID)
		genre.PUT("/:id", handlers.UpdateGenre)
		genre.DELETE("/:id", handlers.DeleteGenre)
	}
	director := authorized.Group("/directors")
	{
		director.GET("/", handlers.GetAllDirector)
		director.POST("/", handlers.CreateDirector)
		director.GET("/:id", handlers.GetDirectorByID)
		director.PUT("/:id", handlers.UpdateDirector)
		director.DELETE("/:id", handlers.DeleteDirector)
	}

}
