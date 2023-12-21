package app

import (
	"context"
	"fmt"

	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
	"github.com/Edwinfpirajan/server.git/internal/domain/ports/db/interfaces"
)

// UserApp is the interface of user application
type UserApp interface {
	GetUsers(ctx context.Context) ([]dto.User, error)
}

// userApp is the implementation of UserApp
type userApp struct {
	repo interfaces.NewUserRepository
}

// NewUserApp is the constructor of UserApp
func NewUserApp(repo interfaces.NewUserRepository) UserApp {
	return &userApp{repo}
}

// GetUsers is the implementation of UserApp.GetUsers
func (ua *userApp) GetUsers(ctx context.Context) ([]dto.User, error) {
	users, err := ua.repo.GetUsers(ctx)
	fmt.Println(users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
