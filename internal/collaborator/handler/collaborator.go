package handler

import (
	"github.com/labstack/echo/v4"
)

type CollaboratorHandler interface {
	GetCollaborator(c echo.Context) error
}

type collaborator struct{}

func NewCollaboratorHandler() CollaboratorHandler {
	return &collaborator{}
}

func (h *collaborator) GetCollaborator(c echo.Context) error {
	return c.JSON(200, "Hello World")
}
