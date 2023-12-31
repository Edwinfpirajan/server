package handlers

import (
	"github.com/Edwinfpirajan/server.git/internal/app"
	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
	"github.com/labstack/echo/v4"
)

type RoleHandler interface {
	CreateRole(c echo.Context) error
	GetRoles(c echo.Context) error
}

type roleHandler struct {
	app app.RoleApp
}

func NewRoleHandler(ra app.RoleApp) RoleHandler {
	return &roleHandler{ra}
}

func (hand *roleHandler) CreateRole(c echo.Context) error {
	var role dto.Role
	if err := c.Bind(&role); err != nil {
		return err
	}

	role, err := hand.app.CreateRole(c.Request().Context(), role)
	if err != nil {
		return err
	}

	return c.JSON(200, role)
}

func (hand *roleHandler) GetRoles(c echo.Context) error {
	roles, err := hand.app.GetRoles(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(200, roles)
}
