package implementation

import (
	"context"

	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
	"github.com/Edwinfpirajan/server.git/internal/domain/ports/db/interfaces"
)

type roleRepository struct {
	builder *SQLBuilder
}

func NewRoleRepository(builder *SQLBuilder) interfaces.NewRoleRepository {
	return &roleRepository{
		builder: builder,
	}
}

func (rr *roleRepository) CreateRole(ctx context.Context, role dto.Role) (dto.Role, error) {
	result := rr.builder.Insert("roles", map[string]interface{}{
		"name": role.Name,
	}).Query()

	if result.Error != nil {
		return dto.Role{}, result.Error
	}

	return role, nil
}

func (rr *roleRepository) GetRoles(ctx context.Context) (dto.Roles, error) {
	var roles dto.Roles
	result := rr.builder.Select("id", "name").From("roles").Query().Find(&roles)

	if result.Error != nil {
		return nil, result.Error
	}

	return roles, nil
}
