package handlers

import (
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
	"github.com/theakshaygupta/go-authapi/dbo"
	"github.com/theakshaygupta/go-authapi/models"
	"net/http"
)

type AdminHandler struct {
	db dbo.DatabaseOps
}

func NewAdminHandler(db dbo.DatabaseOps) AdminHandler {
	return AdminHandler{db: db}
}

func (h *AdminHandler) GetUsers(c echo.Context) error {
	users, err := h.db.GetUsers(bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponse{Message: "cannot get users", Error: err})
	}
	return c.JSON(http.StatusOK, models.HttpResponse{Users: users})
}
