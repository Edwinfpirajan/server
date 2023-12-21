package interfaces

import (
	"context"

	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
)

// NewUserRepository is the interface of user repository
type NewRoleRepository interface {
	CreateRole(ctx context.Context, role dto.Role) (dto.Role, error)
	GetRoles(ctx context.Context) (dto.Roles, error)
}
