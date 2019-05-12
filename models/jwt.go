package models

import "github.com/dgrijalva/jwt-go"

var JwtKey = []byte("fuckyouandyouandyou")

type JWTClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
