package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Username string
	Password string
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
