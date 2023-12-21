package repository

import (
	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
	"github.com/google/uuid"
)

// UserRepositoryPort define el puerto para la gestión de usuarios
type UserRepository interface {
}

// UserRepository es la implementación concreta del puerto UserRepositoryPort
type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (ur *userRepository) GetUsers() ([]dto.User, error) {
	// Do some database stuff here
	return []dto.User{
		{
			ID:   uuid.New(),
			Name: "Edwin",
			Age:  23,
		},
		{
			ID:   uuid.New(),
			Name: "Edwin",
			Age:  23,
		},
		{
			ID:   uuid.New(),
			Name: "Edwin",
			Age:  23,
		},
	}, nil
}
