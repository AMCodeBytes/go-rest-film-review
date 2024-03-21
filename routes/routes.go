package routes

import (
	"github.com/AMCodeBytes/go-rest-film-review/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	protected := server.Group("/")
	protected.Use(middlewares.Auth)

	// Auth
	server.POST("/signup", createUser)
	server.POST("/login", login)

	// Users
	server.GET("/users", getUsers)
	server.GET("/users/:id", getUser)
	server.POST("/users", createUser)
	server.PUT("/users/:id", updateUser)
	server.DELETE("/users/:id", deleteUser)

	// Films
	server.GET("/films", getFilms)
	server.GET("/films/:id", getFilm)

	protected.POST("/films", createFilm)
	protected.PUT("/films/:id", updateFilm)
	protected.DELETE("/films/:id", deleteFilm)
}
