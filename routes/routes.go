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
}
