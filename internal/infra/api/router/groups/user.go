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
	uh handlers.UserHandler
}

// NewUserGroup is the constructor of userGroup
func NewUserGroup(uh handlers.UserHandler) UserGroup {
	return &userGroup{
		uh,
	}
}

// Resource is the implementation of UserGroup.Resource
func (groups userGroup) Resource(g *echo.Group) {

	group := g.Group("/user")

	group.POST("/create", groups.uh.CreateUser)
	group.GET("/get", groups.uh.GetUsers)
}
