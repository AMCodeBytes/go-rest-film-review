package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AMCodeBytes/go-rest-film-review/models"
	"github.com/stretchr/testify/assert"
)

// Helper function to create a user
func CreateUser() {

}

// Helper function to create a film
func CreateFilm() {

}

func TestSignUpRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	user := models.User{
		Name:     "First Second",
		Email:    "email@test.com",
		Password: "Password123!",
	}

	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonValue))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestLoginRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	rec := httptest.NewRecorder()

	user := models.User{
		Name:     "First Second",
		Email:    "email@test.com",
		Password: "Password123!",
	}

	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	user = models.User{
		Email:    "email@test.com",
		Password: "Password123!",
	}

	jsonValue, _ = json.Marshal(user)

	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUsersRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/users", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
