package router

import (
	"github.com/labstack/echo"
	"github.com/theakshaygupta/go-authapi/dbo"
	"github.com/theakshaygupta/go-authapi/handlers"
	"github.com/theakshaygupta/go-authapi/middlewares"
	"github.com/theakshaygupta/go-authapi/models"
	"net/http"
)

func rootHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, models.HttpResponse{Message: "Eureka! Your piece of shit server is running"})
}

func RegisterRoutes(db dbo.DatabaseOps, e *echo.Echo) {
	// root
	e.GET("/", rootHandler)

	// auth routes
	authHandler := handlers.NewAuthHandler(db)
	authGroup := e.Group("auth")
	authGroup.POST("/register", authHandler.Register)
	authGroup.POST("/login", authHandler.Login)
	authGroup.POST("/refresh", authHandler.Refresh, middlewares.AuthenticationMiddleware(), middlewares.AuthorizationMiddleware("refresh"))

	// admin routes
	adminHandler := handlers.NewAdminHandler(db)
	adminGroup := e.Group("admin")
	adminGroup.Use(middlewares.AuthenticationMiddleware(), middlewares.AuthorizationMiddleware("basicAdminPermissions"))
	adminGroup.GET("/users", adminHandler.GetUsers)
}
