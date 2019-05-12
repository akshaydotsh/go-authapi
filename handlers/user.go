package handlers

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
	"github.com/theakshaygupta/go-authapi/dbo"
	"github.com/theakshaygupta/go-authapi/models"
	"github.com/theakshaygupta/go-authapi/utils"
	"net/http"
	"time"
)

type UserHandler struct {
	db dbo.DatabaseOps
}

func NewUserHandler(db dbo.DatabaseOps) UserHandler {
	return UserHandler{db: db}
}

func (h *UserHandler) Register(c echo.Context) error {
	// bind
	user := &models.User{Id: bson.NewObjectId()}
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{Message: "Cannot Signup User", Error: err})
	}

	// validate
	if err, ok := user.Validate(); !ok {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{Message: "ValidationError", Error: err})
	}

	// check for existing user
	existing, err := h.db.FindUser(bson.M{"email": user.Email})
	if (models.User{}) != existing {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{Message: "User Already Exists"})
	}
	if err != nil && err != mgo.ErrNotFound {
		return c.JSON(http.StatusInternalServerError, models.HttpResponse{
			Message: "Error Occured while registering user",
			Error:   err})
	}

	// register user
	user.CreatedAt = int(time.Now().Unix())
	if err := h.db.InsertUser(*user); err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponse{
			Message: "Error Occured while registering user",
			Error:   err})
	}
	return c.JSON(http.StatusCreated, models.HttpResponse{Message: "User Registered"})
}

func (h *UserHandler) Login(c echo.Context) error {
	// bind
	user := &models.UserLoginCreds{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{Message: "Cannot Login User", Error: err})
	}

	// validate
	if err, ok := user.Validate(); !ok {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{Message: "ValidationError", Error: err})
	}

	found, err := h.db.FindUser(bson.M{"email": user.Email, "password": user.Password, "role": user.Role})
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, models.HttpResponse{
				Message: "No User with this email password combination found"})
		} else {
			return c.JSON(http.StatusInternalServerError, models.HttpResponse{Message: "Cannot login", Error: err})
		}
	}

	// access token
	accessToken, accessTokenExpiresAt, err := utils.CreateJWTToken(user.Email, user.Role, found.Id.Hex(), false)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponse{Message: "Cannot login user", Error: err})
	}

	// refresh token
	refreshToken, refreshTokenExpiresAt, err := utils.CreateJWTToken(user.Email, user.Role, found.Id.Hex(), true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponse{Message: "Cannot login user", Error: err})
	}

	return c.JSON(http.StatusOK, models.HttpResponse{
		Message:               "Logged in",
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessTokenExpiresAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshTokenExpiresAt,
	})

}
