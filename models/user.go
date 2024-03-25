package models

import (
	"errors"
	"slices"
	"time"

	"github.com/AMCodeBytes/go-rest-film-review/utils"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Likes     []Like
	Dislikes  []Dislike
	Comments  []Comment
	Bookmarks []Bookmark
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Like struct {
	FilmID string
}

type Dislike struct {
	FilmID string
}

type Comment struct {
	FilmID  string
	Comment string
}

type Bookmark struct {
	FilmID string
}

var users = []User{}

func GetUserByID(id string) User {
	idx := slices.IndexFunc(users, func(u User) bool { return u.ID == id })
	return users[idx]
}

func GetAllUsers() []User {
	return users
}

func (user User) Login() (string, error) {
	idx := slices.IndexFunc(users, func(u User) bool { return u.Email == user.Email })

	if idx == -1 {
		return "", errors.New("no user exists")
	}

	u := &users[idx]

	match := utils.AuthenticatePassword(user.Password, u.Password)

	if !match {
		return "", errors.New("invalid credentials")
	}

	return u.ID, nil
}

func (user User) Create() {
	users = append(users, user)
}

func (user User) Update(id string) error {
	idx := slices.IndexFunc(users, func(u User) bool { return u.ID == id })

	if idx == -1 {
		return errors.New("no user exists")
	}

	updateUser := &users[idx]

	(*updateUser).Name = user.Name
	(*updateUser).Email = user.Email
	(*updateUser).Password = user.Password
	(*updateUser).Likes = user.Likes
	(*updateUser).Dislikes = user.Dislikes
	(*updateUser).Comments = user.Comments
	(*updateUser).UpdatedAt = time.Now()

	return nil
}

func (user User) Delete(id string) error {
	idx := slices.IndexFunc(users, func(u User) bool { return u.ID == id })

	if idx == -1 {
		return errors.New("no user exists")
	}

	users = slices.Replace(users, idx, idx+1)
	return nil
}

func (user User) Like(id string, filmId string) (int, error) {
	var like Like
	idx := slices.IndexFunc(users, func(u User) bool { return u.ID == id })

	if idx == -1 {
		return 0, errors.New("no user exists")
	}

	selectedUser := &users[idx]

	existsId := slices.IndexFunc(selectedUser.Likes, func(l Like) bool { return l.FilmID == filmId })

	if existsId == -1 {
		like.FilmID = filmId

		(*selectedUser).Likes = append((*selectedUser).Likes, like)
		return 1, nil
	}

	(*selectedUser).Likes = slices.Replace((*selectedUser).Likes, existsId, existsId+1)
	return -1, nil
}

func (user User) Dislike(id string, filmId string) (int, error) {
	var dislike Dislike
	idx := slices.IndexFunc(users, func(u User) bool { return u.ID == id })

	if idx == -1 {
		return 0, errors.New("no user exists")
	}

	selectedUser := &users[idx]

	existsId := slices.IndexFunc(selectedUser.Dislikes, func(d Dislike) bool { return d.FilmID == filmId })

	if existsId == -1 {
		dislike.FilmID = filmId

		(*selectedUser).Dislikes = append((*selectedUser).Dislikes, dislike)
		return 1, nil
	}

	(*selectedUser).Dislikes = slices.Replace((*selectedUser).Dislikes, existsId, existsId+1)
	return -1, nil
}

func (user User) Bookmark(id string, filmId string) error {
	var bookmark Bookmark
	idx := slices.IndexFunc(users, func(u User) bool { return u.ID == id })

	if idx == -1 {
		return errors.New("no user exists")
	}

	selectedUser := &users[idx]

	existsId := slices.IndexFunc(selectedUser.Bookmarks, func(b Bookmark) bool { return b.FilmID == filmId })

	if existsId == -1 {
		bookmark.FilmID = filmId

		(*selectedUser).Bookmarks = append((*selectedUser).Bookmarks, bookmark)
		return nil
	}

	(*selectedUser).Bookmarks = slices.Replace((*selectedUser).Bookmarks, existsId, existsId+1)
	return nil
}
