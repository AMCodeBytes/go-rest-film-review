package models

import "time"

type User struct {
	ID        string
	Name      string `binding:"required"`
	Email     string `binding:"required"`
	Password  string `binding:"required"`
	Likes     string
	Dislikes  string
	Comments  string
	Bookmarks string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

var users = []User{}

func GetAllUsers() []User {
	return users
}

func (u User) Create() {
	users = append(users, u)
}
