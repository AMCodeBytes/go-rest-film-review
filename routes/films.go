package routes

import (
	"net/http"

	"github.com/AMCodeBytes/go-rest-film-review/models"
	"github.com/AMCodeBytes/go-rest-film-review/utils"
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

	filmID, err := utils.GenerateUUID()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate UUID."})
	}

	film.ID = filmID

	userId := context.GetString("userId")

	film.CreatedBy = userId

	film.Create()

	context.JSON(http.StatusCreated, gin.H{"message": "Film created!", "film": film})
}

func updateFilm(context *gin.Context) {
	id := context.Param("id")
	var film models.Film

	userId := context.GetString("userId")
	film = models.GetFilmByID(id)

	if film.Locked {
		if film.CreatedBy != userId {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorised to update film"})
			return
		}
	}

	err := context.ShouldBindJSON(&film)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse request data."})
		return
	}

	err = film.Update(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "The film failed to update."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Film was successfully updated."})
}

func deleteFilm(context *gin.Context) {
	id := context.Param("id")
	var film models.Film

	userId := context.GetString("userId")
	film = models.GetFilmByID(id)

	if film.Locked {
		if film.CreatedBy != userId {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorised to delete film"})
			return
		}
	}

	err := film.Delete(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "The film failed to be deleted."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Film was successfully deleted."})
}

func likeFilm(context *gin.Context) {
	id := context.Param("id")
	var film models.Film
	var user models.User
	userId := context.GetString("userId")
	film = models.GetFilmByID(id)
	user = models.GetUserByID(userId)

	like, err := user.Like(userId, id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update the user's like."})
		return
	}

	err = film.UpdateLikes(id, like)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update the film's like."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully updated the likes."})
}

func dislikeFilm(context *gin.Context) {
	id := context.Param("id")
	var film models.Film
	var user models.User
	userId := context.GetString("userId")
	film = models.GetFilmByID(id)
	user = models.GetUserByID(userId)

	dislike, err := user.Dislike(userId, id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update the user's dislike."})
		return
	}

	err = film.UpdateDislike(id, dislike)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update the film's dislike."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully updated the dislikes."})
}
