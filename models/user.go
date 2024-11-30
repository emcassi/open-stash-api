package models

import "time"

type User struct {
	Id string
	Name string
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserCreationEmailPw struct {
	Name string
	Email string
	Password string
}
