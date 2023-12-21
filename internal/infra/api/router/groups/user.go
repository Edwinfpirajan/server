package groups

import (
	"github.com/Edwinfpirajan/server.git/internal/infra/api/handlers"
	"github.com/labstack/echo/v4"
)

// UserGroup is the interface of user group
type UserGroup interface {
	Resource(g *echo.Group)
}

// userGroup is the implementation of UserGroup
type userGroup struct {
	h handlers.UserHandler
}

// NewUserGroup is the constructor of userGroup
func NewUserGroup(h handlers.UserHandler) UserGroup {
	return &userGroup{
		h,
	}
}

// Resource is the implementation of UserGroup.Resource
func (groups userGroup) Resource(g *echo.Group) {
	//inyeccion de dependencias repository, app, handler, como hacerlas?

	group := g.Group("/user")

	group.POST("/create", groups.h.CreateUser)
	group.GET("/get", groups.h.GetUsers)
}
