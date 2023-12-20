package handlers

import (
	"net/http"

	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
	"github.com/Edwinfpirajan/server.git/internal/domain/entity"
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
	return c.JSON(http.StatusOK, entity.Response{
		Message: "Hello World!",
		Data: dto.Collaborator{
			Name: "Edwin",
			Age:  25,
		},
	})
}
