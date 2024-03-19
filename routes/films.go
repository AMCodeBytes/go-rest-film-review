package routes

import (
	"net/http"

	"github.com/AMCodeBytes/go-rest-film-review/models"
	"github.com/gin-gonic/gin"
)

func getFilms(context *gin.Context) {
	films := models.GetAllFilms()
	context.JSON(http.StatusOK, films)
}

func getFilm(context *gin.Context) {
	id := context.Param("id")
	context.JSON(http.StatusOK, models.GetFilmByID(id))
}

func createFilm(context *gin.Context) {
	var film models.Film

	err := context.ShouldBindJSON(&film)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	// film.ID = "123-abc-qwerty"
	// film.Name = "Film name"
	// film.Description = "Film description will go here for the film."
	// film.Categories = []string{"Action", "Thriller", "Horror"}
	// film.Likes = 15
	// film.Dislikes = 2
	// film.Comments = []string{"This is a great film", "Could do with more action"}

	film.Create()

	context.JSON(http.StatusCreated, gin.H{"message": "Film created!", "film": film})
}

func updateFilm(context *gin.Context) {
	id := context.Param("id")
	var film models.Film

	err := context.ShouldBindJSON(&film)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse request data."})
		return
	}

	err2 := film.Update(id)

	if err2 != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "The film failed to update."})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Film was successfully updated."})
}

func deleteFilm(context *gin.Context) {
	id := context.Param("id")

	err := models.DeleteFilm(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "The film failed to be deleted."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Film was successfully deleted."})
}