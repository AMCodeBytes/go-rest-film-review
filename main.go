package main

import (
	"net/http"

	"github.com/AMCodeBytes/go-rest-film-review/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Default returns an engine instance with logger & recovery middleware
	server := gin.Default()

	server.GET("/users", getUsers)
	server.POST("/users", createUser)

	server.GET("/films", getFilms)
	server.POST("/films", createFilm)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	server.Run()
}

func getUsers(context *gin.Context) {
	users := models.GetAllUsers()
	context.JSON(http.StatusOK, users)
}

func createUser(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse request data."})
		return
	}

	user.ID = "123-abc-qwerty"
	user.Name = "First Second"
	user.Email = "email@test.com"
	user.Password = "Password123!"

	user.Create()

	context.JSON(http.StatusCreated, gin.H{"message": "User created!", "user": user})
}

func getFilms(context *gin.Context) {
	films := models.GetAllFilms()
	context.JSON(http.StatusOK, films)
}

func createFilm(context *gin.Context) {
	var film models.Film

	err := context.ShouldBindJSON(&film)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	film.ID = "123-abc-qwerty"
	film.Name = "Film name"
	film.Description = "Film description will go here for the film."

	film.Create()

	context.JSON(http.StatusCreated, gin.H{"message": "Film created!", "film": film})
}
