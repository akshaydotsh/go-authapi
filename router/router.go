package router

import (
	"github.com/labstack/echo"
	"github.com/theakshaygupta/go-authapi/dbo"
	"github.com/theakshaygupta/go-authapi/handlers"
	"github.com/theakshaygupta/go-authapi/models"
	"net/http"
)

func rootHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, models.HttpResponse{Message: "Eureka! Your piece of shit server is running"})
}

func RegisterRoutes(db dbo.DatabaseOps, e *echo.Echo) {
	// root
	e.GET("/", rootHandler)

	// user routes
	userHandler := handlers.NewUserHandler(db)
	userGroup := e.Group("users")
	userGroup.POST("/register", userHandler.Register)
	userGroup.POST("/login", userHandler.Login)
}
