package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/theakshaygupta/go-authapi/models"
	"github.com/theakshaygupta/go-authapi/static"
	"net/http"
)

var codesMapping = static.Mapping

func apiCodeValidForRole(apicode, role string) bool {
	for key := range codesMapping {
		if key == role {
			for _, allowedCode := range codesMapping[key] {
				if allowedCode == apicode {
					return true
				}
			}
		}
	}
	return false
}

func AuthorizationMiddleware(apicode string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user")
			token := user.(*jwt.Token)

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, models.HttpResponse{Message: "Access Denied"})
			}
			role := claims["role"].(string)

			if apiCodeValidForRole(apicode, role) {
				return next(c)
			}
			return c.JSON(http.StatusUnauthorized, models.HttpResponse{Message: "Access Denied"})
		}
	}
}
