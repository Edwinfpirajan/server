package groups

import (
	"github.com/Edwinfpirajan/server.git/internal/infra/api/handlers"
	"github.com/labstack/echo/v4"
)

type RoleGroup interface {
	Resource(g *echo.Group)
}

type roleGroup struct {
	rh handlers.RoleHandler
}

func NewRoleGroup(rh handlers.RoleHandler) RoleGroup {
	return &roleGroup{
		rh,
	}
}

func (groups roleGroup) Resource(g *echo.Group) {

	group := g.Group("/role")

	group.POST("/create", groups.rh.CreateRole)
	group.GET("/get", groups.rh.GetRoles)
}
