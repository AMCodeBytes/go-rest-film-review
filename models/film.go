package models

import (
	"errors"
	"slices"
	"time"
)

type Film struct {
	ID          string
	Name        string `binding:"required"`
	Type        string `binding:"required"`
	Description string `binding:"required"`
	Thumbnail   string
	Categories  []string
	ReleasedAt  time.Time
	Likes       int
	Dislikes    int
	Comments    []string
	Locked      bool
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

var films = []Film{}

func GetFilmByID(id string) Film {
	idx := slices.IndexFunc(films, func(f Film) bool { return f.ID == id })
	return films[idx]
}

func GetAllFilms() []Film {
	return films
}

func (film Film) Create() {
	films = append(films, film)
}

func (film Film) Update(id string) error {
	idx := slices.IndexFunc(films, func(f Film) bool { return f.ID == id })

	if idx == -1 {
		return errors.New("no film exists")
	}

	updateFilm := &films[idx]

	(*updateFilm).Name = film.Name
	(*updateFilm).Type = film.Type
	(*updateFilm).Description = film.Description
	(*updateFilm).Thumbnail = film.Thumbnail
	(*updateFilm).Categories = film.Categories
	(*updateFilm).ReleasedAt = film.ReleasedAt
	(*updateFilm).Likes = film.Likes
	(*updateFilm).Dislikes = film.Dislikes
	(*updateFilm).Comments = film.Comments
	(*updateFilm).UpdatedAt = time.Now()

	return nil
}

func (film Film) Delete(id string) error {
	idx := slices.IndexFunc(films, func(f Film) bool { return f.ID == id })

	if idx == -1 {
		return errors.New("no film exists")
	}

	films = slices.Replace(films, idx, idx+1)
	return nil
}

func (film Film) UpdateLikes(id string, like int) error {
	idx := slices.IndexFunc(films, func(f Film) bool { return f.ID == id })

	if idx == -1 {
		return errors.New("no film exists")
	}

	updateFilm := &films[idx]

	(*updateFilm).Likes = (*updateFilm).Likes + like

	return nil
}

func (film Film) UpdateDislike(id string, dislike int) error {
	idx := slices.IndexFunc(films, func(f Film) bool { return f.ID == id })

	if idx == -1 {
		return errors.New("no film exists")
	}

	updateFilm := &films[idx]

	(*updateFilm).Dislikes = (*updateFilm).Dislikes + dislike

	return nil
}
