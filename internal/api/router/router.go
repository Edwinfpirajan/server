package router

import (
	"github.com/Edwinfpirajan/server.git/internal/api/groups"
	"github.com/Edwinfpirajan/server.git/internal/health/handler"
	"github.com/labstack/echo/v4"
)

// Router is the implementation of Group
type Router struct {
	server       *echo.Echo
	collaborator groups.CollaboratorGroup
}

// NewRouter is the constructor of Router
func NewRouter(
	server *echo.Echo,
	collaborator groups.CollaboratorGroup,
) *Router {
	return &Router{
		server,
		collaborator,
	}
}

// Init is the implementation of Group.Init
func (r *Router) Init() {

	r.server.GET("/health", handler.HealthCheck)

	basePath := r.server.Group("/api/v1")

	r.collaborator.Resource(basePath)

}
