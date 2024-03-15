package models

import "time"

type Film struct {
	ID          string
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Thumbnail   string
	Categories  string
	ReleasedAt  time.Time
	Likes       int
	Dislikes    int
	Comments    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

// Create a slice of films
var films = []Film{}

func GetAllFilms() []Film {
	return films
}

// Save a film
func (f Film) Save() {
	films = append(films, f)
}
