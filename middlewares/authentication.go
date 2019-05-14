package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/theakshaygupta/go-authapi/models"
)

func AuthenticationMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    models.JwtKey,
	})
}
