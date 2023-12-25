package router

import (
	"github.com/Edwinfpirajan/server.git/internal/infra/api/handlers"
	"github.com/Edwinfpirajan/server.git/internal/infra/api/router/groups"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Router is the implementation of Group
type Router struct {
	server *echo.Echo
	user   groups.UserGroup
	role   groups.RoleGroup
}

// NewRouter is the constructor of Router
func NewRouter(
	server *echo.Echo,
	user groups.UserGroup,
	role groups.RoleGroup,
) *Router {
	return &Router{
		server,
		user,
		role,
	}
}

// Init is the implementation of Group.Init
func (r *Router) Init() {

	r.server.Use(middleware.Recover())

	r.server.GET("/health", handlers.HealthCheck)

	base := r.server.Group("/api/v1")

	r.user.Resource(base)
	r.role.Resource(base)

}
