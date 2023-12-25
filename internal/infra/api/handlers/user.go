package handlers

import (
	"fmt"
	"net/http"

	"github.com/Edwinfpirajan/server.git/internal/app"
	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
	"github.com/Edwinfpirajan/server.git/internal/domain/entity"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
)

// UserHandler is the interface of user handler
type UserHandler interface {
	CreateUser(c echo.Context) error
	GetUsers(c echo.Context) error
}

// userHandler is the implementation of UserHandler
type userHandler struct {
	app app.UserApp
}

// NewUserHandler is the constructor of userHandler
func NewUserHandler(app app.UserApp) UserHandler {
	return &userHandler{
		app,
	}
}

// CreateUser is the implementation of UserHandler.CreateUser
func (hand *userHandler) CreateUser(c echo.Context) error {
	var request dto.User

	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := request.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := hand.app.CreateUser(c.Request().Context(), request); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.Response{
		Message: "User created successfully",
	})
}

// GetUsers is the implementation of UserHandler.GetUsers
func (hand *userHandler) GetUsers(c echo.Context) error {
	var request dto.UsersRequest
	fmt.Println(color.Red("request.Filter.ID: "), request.Filter.ID)

	err := c.Bind(&request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{
			Message: "No fue posible obtener los usuarios",
			Data:    nil,
		})
	}

	// err = request.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{
			Message: "No fue posible obtener los usuarios",
			Data:    nil,
		})
	}

	users, err := hand.app.GetUsers(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.Response{
			Message: "Internal Server Error",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, entity.Response{
		Message: "Users retrieved successfully",
		Data:    users,
	})
}
