package http_server

import (
	"github.com/labstack/echo"
	"github.com/theakshaygupta/go-authapi/dbo"
	"github.com/theakshaygupta/go-authapi/middlewares"
	"github.com/theakshaygupta/go-authapi/router"
)

func NewEchoServer(db dbo.DatabaseOps) *echo.Echo {
	e := echo.New()

	// set global middlewares
	middlewares.SetGlobalMiddlewares(e)

	// register Routes
	router.RegisterRoutes(db, e)

	return e
}
