package app

import (
	"context"
	"fmt"

	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
	"github.com/Edwinfpirajan/server.git/internal/domain/ports/db/interfaces"
)

// UserApp is the interface of user application
type UserApp interface {
	CreateUser(ctx context.Context, user dto.User) (dto.User, error)
	GetUsers(ctx context.Context) (dto.Users, error)
}

// userApp is the implementation of UserApp
type userApp struct {
	repo interfaces.NewUserRepository
}

// NewUserApp is the constructor of UserApp
func NewUserApp(repo interfaces.NewUserRepository) UserApp {
	return &userApp{repo}
}

func (ua *userApp) CreateUser(ctx context.Context, user dto.User) (dto.User, error) {
	user, err := ua.repo.CreateUser(ctx, user)
	fmt.Println(user)
	if err != nil {
		return dto.User{}, err
	}

	return user, nil
}

// GetUsers is the implementation of UserApp.GetUsers
func (ua *userApp) GetUsers(ctx context.Context) (dto.Users, error) {
	users, err := ua.repo.GetUsers(ctx)
	fmt.Println(users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
