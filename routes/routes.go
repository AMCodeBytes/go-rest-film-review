package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", createUser)
	server.POST("/login", login)

	server.GET("/users", getUsers)
	server.GET("/users/:id", getUser)
	server.POST("/users", createUser)
	server.PUT("/users/:id", updateUser)
	server.DELETE("/users/:id", deleteUser)

	server.GET("/films", getFilms)
	server.GET("/films/:id", getFilm)
	server.POST("/films", createFilm)
	server.PUT("/films/:id", updateFilm)
	server.DELETE("/films/:id", deleteFilm)
}
