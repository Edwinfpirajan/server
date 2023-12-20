package groups

import (
	"github.com/Edwinfpirajan/server.git/internal/collaborator/handler"
	"github.com/labstack/echo/v4"
)

// CollaboratorGroup is the interface of collaborator group
type CollaboratorGroup interface {
	Resource(g *echo.Group)
}

// collaboratorGroup is the implementation of CollaboratorGroup
type collaboratorGroup struct {
	h handler.CollaboratorHandler
}

// NewCollaboratorGroup is the constructor of collaboratorGroup
func NewCollaboratorGroup(h handler.CollaboratorHandler) CollaboratorGroup {
	return &collaboratorGroup{
		h: h,
	}
}

// Resource is the implementation of CollaboratorGroup.Resource
func (c collaboratorGroup) Resource(g *echo.Group) {

	collaboratorHandler := handler.NewCollaboratorHandler()

	collaborator := g.Group("/collaborator")

	collaborator.GET("/get", collaboratorHandler.GetCollaborator)
}
