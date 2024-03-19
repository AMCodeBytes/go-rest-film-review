package routes

import (
	"net/http"

	"github.com/AMCodeBytes/go-rest-film-review/models"
	"github.com/AMCodeBytes/go-rest-film-review/utils"
	"github.com/gin-gonic/gin"
)

func getUsers(context *gin.Context) {
	users := models.GetAllUsers()
	context.JSON(http.StatusOK, users)
}

func getUser(context *gin.Context) {
	id := context.Param("id")
	context.JSON(http.StatusOK, models.GetUserByID(id))
}

func createUser(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse request data."})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash password."})
	}

	user.Password = hashedPassword

	// user.ID = "123-abc-qwerty"
	// user.Name = "First Second"
	// user.Email = "email@test.com"
	// user.Password = "Password123!"

	user.Create()

	context.JSON(http.StatusCreated, gin.H{"message": "User created!", "user": user})
}

func updateUser(context *gin.Context) {
	id := context.Param("id")
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse request data."})
		return
	}

	err2 := user.Update(id)

	if err2 != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "The user failed to update."})
	}

	context.JSON(http.StatusOK, gin.H{"message": "User was successfully updated."})
}

func deleteUser(context *gin.Context) {
	id := context.Param("id")
	var user models.User

	err := user.Delete(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "The user failed to be deleted."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User was successfully deleted."})
}
