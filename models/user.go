package models

type User struct {
	Base
	Name     string
	Email    *string
	Password *string
}

type UserCreationEmailPw struct {
	Name     string
	Email    string
	Password string
}

type UserLoginEmailPw struct {
	Email    string
	Password string
}
