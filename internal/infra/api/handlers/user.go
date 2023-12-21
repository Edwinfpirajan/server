package handlers

import (
	"net/http"

	"github.com/Edwinfpirajan/server.git/internal/app"
	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
	"github.com/Edwinfpirajan/server.git/internal/domain/entity"
	"github.com/labstack/echo/v4"
)

// UserHandler is the interface of user handler
type UserHandler interface {
	CreateUser(c echo.Context) error
	GetUsers(c echo.Context) error
}

// userHandler is the implementation of UserHandler
type userHandler struct {
	ua app.UserApp
}

// NewUserHandler is the constructor of userHandler
func NewUserHandler(ua app.UserApp) UserHandler {
	return &userHandler{
		ua,
	}
}

// CreateUser is the implementation of UserHandler.CreateUser
func (uh *userHandler) CreateUser(c echo.Context) error {
	var user dto.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{
			Message: "Bad Request",
			Data:    nil,
		})
	}

	user, err = uh.ua.CreateUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.Response{
			Message: "Internal Server Error",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, entity.Response{
		Message: "User created successfully",
		Data:    user,
	})
}

// GetUsers is the implementation of UserHandler.GetUsers
func (uh *userHandler) GetUsers(c echo.Context) error {
	users, err := uh.ua.GetUsers(c.Request().Context())
	if err != nil {

		return c.JSON(http.StatusInternalServerError, entity.Response{
			Message: "Internal Server Error",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, entity.Response{
		Message: "Success",
		Data:    users,
	})
}
