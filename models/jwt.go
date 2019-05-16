package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/theakshaygupta/go-authapi/config"
)

var JWTSecret = []byte(config.Config.JWTSecret)

type JWTClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
