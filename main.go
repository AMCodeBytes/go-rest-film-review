package main

import (
	"net/http"

	"github.com/AMCodeBytes/go-rest-film-review/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Default returns an engine instance with logger & recovery middleware
	server := gin.Default()

	// Setup a route for /ping
	server.GET("/ping", func(context *gin.Context) {
		// Send a response of pong
		context.JSON(http.StatusOK, gin.H{
			"message": "ponging",
		})
	})

	server.GET("/films", getFilms)
	server.POST("/films", createFilm)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	server.Run()
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

	film.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Film created!", "film": film})
}
