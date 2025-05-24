package routes

import (
	"ProjectCinema/handlers"
	"ProjectCinema/middlewares"
	"ProjectCinema/utils"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	//r.POST("/register", handlers.Register)
	//r.POST("/login", handlers.Login)
	authorized := r.Group("/")
	authorized.Use(middlewares.AuthMiddleware())

	adminRoutes := r.Group("/admin", middlewares.AuthMiddleware(), middlewares.AdminOnly())
	{
		// Films
		adminRoutes.GET("/films", handlers.GetAllFilms)
		adminRoutes.POST("/films", handlers.CreateFilms)
		adminRoutes.GET("/films/:id", handlers.GetById)
		adminRoutes.PUT("/films/:id", handlers.UpdateFilms)
		adminRoutes.DELETE("/films/:id", handlers.DeleteFilms)

		// Genres
		adminRoutes.GET("/genres", handlers.GetAllGenres)
		adminRoutes.POST("/genres", handlers.CreateGenre)
		adminRoutes.GET("/genres/:id", handlers.GetGenreByID)
		adminRoutes.PUT("/genres/:id", handlers.UpdateGenre)
		adminRoutes.DELETE("/genres/:id", handlers.DeleteGenre)

		// Directors
		adminRoutes.GET("/directors", handlers.GetAllDirector)
		adminRoutes.POST("/directors", handlers.CreateDirector)
		adminRoutes.GET("/directors/:id", handlers.GetDirectorByID)
		adminRoutes.PUT("/directors/:id", handlers.UpdateDirector)
		adminRoutes.DELETE("/directors/:id", handlers.DeleteDirector)

		// Cinemas
		adminRoutes.GET("/cinemas", handlers.GetCinemas)
		adminRoutes.POST("/cinemas", handlers.CreateCinema)
		adminRoutes.GET("/cinemas/:id", handlers.GetCinemaByID)
		adminRoutes.PUT("/cinemas/:id", handlers.UpdateCinema)
		adminRoutes.DELETE("/cinemas/:id", handlers.DeleteCinema)

		// Sessions
		adminSession := adminRoutes.Group("/sessions")
		{
			adminSession.GET("/", handlers.GetSessions)
			adminSession.POST("/", handlers.CreateSession)
			adminSession.GET("/:id", handlers.GetSessionByID)
			adminSession.PUT("/:id", handlers.UpdateSession)
			adminSession.DELETE("/:id", handlers.DeleteSession)
		}
	}
	user := r.Group("/", middlewares.AuthMiddleware())
	{
		// Фильмдер тізімі мен біреуін көру
		user.GET("/films", handlers.GetAllFilms)
		user.GET("/films/:id", handlers.GetById)

		// Жанрлар тізімі
		user.GET("/genres", handlers.GetAllGenres)
		user.GET("/genres/:id", handlers.GetGenreByID)

		// Режиссерлер
		user.GET("/directors", handlers.GetAllDirector)
		user.GET("/directors/:id", handlers.GetDirectorByID)

		// Кинозалдар
		user.GET("/cinemas", handlers.GetCinemas)
		user.GET("/cinemas/:id", handlers.GetCinemaByID)

		// Сеанстар
		user.GET("/sessions", handlers.GetSessions)
		user.GET("/sessions/:id", handlers.GetSessionByID)

		// Тапсырыстар (орын брондау)
		user.POST("/orders", handlers.CreateOrder) // Жаңа тапсырыс жасау
		//user.GET("/orders", handlers.GetUserOrders)              // Өз тапсырыстарының тізімі
		//user.GET("/orders/:id", handlers.GetOrderByID)           // Белгілі бір тапсырысты көру
		//user.DELETE("/orders/:id", handlers.CancelOrder)         // Тапсырысты жою (мысалы, бронды болдырмау)

	}

	utils.InitResty()
	client := r.Group("/clients")
	{
		client.POST("/register", handlers.RegisterUser)
		client.POST("/login", handlers.LoginUser)

		client.POST("/order", handlers.CreateOrder)
		client.GET("/order", handlers.GetUserOrders)
		client.GET("/order/:id", handlers.GetOrderByID)
		client.DELETE("/order?:id", handlers.CancelOrder)

	}

}
