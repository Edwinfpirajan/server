package interfaces

import (
	"context"

	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
)

type NewUserRepository interface {
	CreateUser(ctx context.Context, user dto.User) (dto.User, error)
	GetUsers(ctx context.Context) (dto.Users, error)
}
