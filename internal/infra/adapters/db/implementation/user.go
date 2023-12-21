package implementation

import (
	"context"

	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
	"github.com/Edwinfpirajan/server.git/internal/domain/ports/db/interfaces"
	"github.com/google/uuid"
)

type userRepository struct {
	builder *SQLBuilder
}

func NewUserRepository(builder *SQLBuilder) interfaces.NewUserRepository {
	return &userRepository{
		builder: builder,
	}
}

func (ur *userRepository) CreateUser(ctx context.Context, user dto.User) (dto.User, error) {
	id := uuid.New()
	result := ur.builder.Insert("users", map[string]interface{}{
		"id":   id,
		"name": user.Name,
		"age":  user.Age,
	}).Query()

	if result.Error != nil {
		return dto.User{}, result.Error
	}

	return user, nil
}

func (ur *userRepository) GetUsers(ctx context.Context) (dto.Users, error) {
	var users dto.Users
	result := ur.builder.Select("id", "name", "age").From("users").Query().Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
