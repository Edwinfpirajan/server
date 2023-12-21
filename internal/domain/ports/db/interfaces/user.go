package interfaces

import (
	"context"

	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
)

type NewUserRepository interface {
	GetUsers(ctx context.Context) (dto.Users, error)
}
