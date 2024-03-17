package models

import (
	"errors"
	"slices"
	"time"
)

type Film struct {
	ID          string
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Thumbnail   string
	Categories  []string
	ReleasedAt  time.Time
	Likes       int
	Dislikes    int
	Comments    []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

// Create a slice of films
var films = []Film{}

func GetFilmByID(id string) Film {
	idx := slices.IndexFunc(films, func(f Film) bool { return f.ID == id })
	return films[idx]
}

func GetAllFilms() []Film {
	return films
}

// Save a film
func (f Film) Create() {
	films = append(films, f)
}

func DeleteFilm(id string) error {
	idx := slices.IndexFunc(films, func(f Film) bool { return f.ID == id })

	if idx != 0 {
		return errors.New("no film exists")
	}

	films = slices.Replace(films, idx, idx+1)
	return nil
}
