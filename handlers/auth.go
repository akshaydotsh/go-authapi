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

type AuthHandler struct {
	db dbo.DatabaseOps
}

func NewAuthHandler(db dbo.DatabaseOps) AuthHandler {
	return AuthHandler{db: db}
}

func (h *AuthHandler) Register(c echo.Context) error {
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

func (h *AuthHandler) Login(c echo.Context) error {
	// bind
	user := &models.UserLoginCreds{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{Message: "Cannot Login User", Error: err})
	}

	// validate
	if err, ok := user.Validate(); !ok {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{Message: "ValidationError", Error: err})
	}

	// find user
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

func (h *AuthHandler) Refresh(c echo.Context) error {
	email, ok := utils.GetFieldFromToken("email", c)
	if !ok {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{Message: "Cannot identify accesstoken"})
	}

	// get the fckn user
	user, err := h.db.FindUser(bson.M{"email": email})
	if err != nil {
		var status int
		if err == mgo.ErrNotFound {
			status = http.StatusBadRequest
		} else {
			status = http.StatusInternalServerError
		}
		return c.JSON(status, models.HttpResponse{Message: "Cannot refresh token", Error: err})
	}

	// give the mf some access token
	accessToken, accessTokenExpiresAt, err := utils.CreateJWTToken(user.Email, user.Role, user.Id.Hex(), false)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponse{Message: "Cannot refresh token", Error: err})
	}
	return c.JSON(http.StatusOK, models.HttpResponse{
		Message:              "New access token created",
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessTokenExpiresAt,
	})
}
