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
	Likes     []string
	Dislikes  []string
	Comments  []string
	Bookmarks []string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
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
