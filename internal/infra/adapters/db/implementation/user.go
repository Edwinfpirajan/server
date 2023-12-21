package implementation

import (
	"context"

	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
	"github.com/Edwinfpirajan/server.git/internal/domain/ports/db/interfaces"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewDbUserRepository(db *gorm.DB) interfaces.NewUserRepository {
	return &userRepository{
		db,
	}
}

func (ur *userRepository) GetUsers(ctx context.Context) ([]dto.User, error) {
	var users []dto.User
	ur.db.WithContext(ctx).Find(&users)
	return users, nil
}
