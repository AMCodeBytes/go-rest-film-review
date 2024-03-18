package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/users", getUsers)
	server.GET("/users/:id", getUser)
	server.POST("/users", createUser)
	server.PUT("/users/:id", updateUser)
	server.POST("/users/delete/:id", deleteUser)

	server.GET("/films", getFilms)
	server.GET("/films/:id", getFilm)
	server.POST("/films", createFilm)
	server.PUT("/films/:id", updateFilm)
	server.POST("/films/delete/:id", deleteFilm)
}
