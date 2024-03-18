package models

import (
	"errors"
	"slices"
	"time"
)

type User struct {
	ID        string
	Name      string `binding:"required"`
	Email     string `binding:"required"`
	Password  string `binding:"required"`
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

func (u User) Create() {
	users = append(users, u)
}

func UpdateUser(id string, user User) error {
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

func DeleteUser(id string) error {
	idx := slices.IndexFunc(users, func(u User) bool { return u.ID == id })

	if idx == -1 {
		return errors.New("no user exists")
	}

	users = slices.Replace(users, idx, idx+1)
	return nil
}
