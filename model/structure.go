package model

import "github.com/golang-jwt/jwt"

type Member struct {
	Name  string
	Age   int
	Phone string
}

type Users struct {
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

type UserFilter struct {
	Firstname string
}

type LoginFilter struct {
	Email    string
	Password string
}

type StructEncodeToken struct {
	Email string
	jwt.RegisteredClaims
}
