package interfaces

import (
	"context"

	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
	"github.com/Edwinfpirajan/server.git/internal/infra/adapters/db/models"
)

type Repository interface {
	// user
	CreateUser(ctx context.Context, user models.User) error
	GetUser(ctx context.Context, request dto.User) (dto.User, error)
	GetUsers(ctx context.Context, request dto.UsersRequest) (dto.Pagination, error)

	// role
	CreateRole(ctx context.Context, role dto.Role) (dto.Role, error)
	GetRoles(ctx context.Context) (dto.Roles, error)
}
