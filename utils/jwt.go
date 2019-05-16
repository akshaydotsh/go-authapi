package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/theakshaygupta/go-authapi/models"
	"time"
)

func CreateJWTToken(email, role, id string, isRefreshToken bool) (string, int64, error) {
	var expiry int64
	if isRefreshToken {
		expiry = int64(time.Now().Add(7 * 24 * time.Hour).Unix())
	} else {
		expiry = int64(time.Now().Add(24 * time.Hour).Unix())
	}
	claims := &models.JWTClaims{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			Id:        id,
			ExpiresAt: expiry,
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString(models.JWTSecret)
	if err != nil {
		return "", 0, err
	}
	return token, claims.StandardClaims.ExpiresAt, nil
}

func GetFieldFromToken(field string, c echo.Context) (string, bool) {
	user := c.Get("user")
	token := user.(*jwt.Token)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", false
	}
	return claims[field].(string), true
}
