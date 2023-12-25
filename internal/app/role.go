package app

import (
	"context"

	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
	"github.com/Edwinfpirajan/server.git/internal/domain/ports/db/interfaces"
)

type RoleApp interface {
	CreateRole(ctx context.Context, role dto.Role) (dto.Role, error)
	GetRoles(ctx context.Context) (dto.Roles, error)
}

type roleApp struct {
	repo interfaces.Repository
}

func NewRoleApp(repo interfaces.Repository) RoleApp {
	return &roleApp{repo}
}

func (ra *roleApp) CreateRole(ctx context.Context, role dto.Role) (dto.Role, error) {
	role, err := ra.repo.CreateRole(ctx, role)
	if err != nil {
		return dto.Role{}, err
	}

	return role, nil
}

func (ra *roleApp) GetRoles(ctx context.Context) (dto.Roles, error) {
	roles, err := ra.repo.GetRoles(ctx)
	if err != nil {
		return nil, err
	}

	return roles, nil
}
