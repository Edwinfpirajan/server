package implementation

import (
	"context"
	"fmt"

	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
	"github.com/Edwinfpirajan/server.git/internal/domain/ports/db/interfaces"
	"github.com/google/uuid"
)

// userRepository is the implementation of NewUserRepository
type userRepository struct {
	builder *SQLBuilder
}

// NewUserRepository is the constructor of NewUserRepository
func NewUserRepository(builder *SQLBuilder) interfaces.NewUserRepository {
	return &userRepository{
		builder: builder,
	}
}

// CreateUser is the implementation of NewUserRepository.CreateUser
func (ur *userRepository) CreateUser(ctx context.Context, user dto.User) (dto.User, error) {
	id := uuid.New()
	fmt.Println(id)
	result := ur.builder.Insert("users", map[string]interface{}{
		"id":          id,
		"document":    user.Document,
		"f_name":      user.FName,
		"l_name":      user.LName,
		"smart_email": user.SmartEmail,
		"password":    user.Password,
		"fk_role_id":  user.FkRoleID,
		"created_at":  user.CreatedAt,
	}).Query()

	if result.Error != nil {
		return dto.User{}, result.Error
	}

	return user, nil
}

// GetUsers is the implementation of NewUserRepository.GetUsers
func (ur *userRepository) GetUsers(ctx context.Context, request dto.UsersRequest) (dto.Pagination, error) {

}
