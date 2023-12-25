package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
	"github.com/Edwinfpirajan/server.git/internal/domain/ports/db/interfaces"
	"github.com/Edwinfpirajan/server.git/internal/infra/adapters/db/models"
	"github.com/Edwinfpirajan/server.git/utils/password"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
)

// UserApp is the interface of user application
type UserApp interface {
	CreateUser(ctx context.Context, user dto.User) error
	GetUsers(ctx context.Context, request dto.UsersRequest) (dto.Pagination, error)
}

// userApp is the implementation of UserApp
type userApp struct {
	repo interfaces.Repository
}

// NewUserApp is the constructor of UserApp
func NewUserApp(repo interfaces.Repository) UserApp {
	return &userApp{repo}
}

func (app *userApp) CreateUser(ctx context.Context, user dto.User) error {
	userDB, err := app.GetUser(ctx, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if userDB.ID != uuid.Nil.String() {
		return echo.NewHTTPError(http.StatusBadRequest, "User already exists")
	}

	// Generate hashed password
	hashedPassword, err := password.HashPassword(user.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error hashing password")
	}

	// Create user in database
	err = app.repo.CreateUser(ctx, models.User{
		ID:         uuid.New(),
		Document:   user.Document,
		FName:      user.FName,
		LName:      user.LName,
		Password:   hashedPassword,
		SmartEmail: user.SmartEmail,
		FkRoleID:   user.FkRoleID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (app *userApp) GetUser(ctx context.Context, request dto.User) (dto.User, error) {
	user, err := app.repo.GetUser(ctx, request)
	if err != nil {
		return dto.User{}, err
	}

	return user, nil
}

// GetUsers is the implementation of UserApp.GetUsers
func (app *userApp) GetUsers(ctx context.Context, request dto.UsersRequest) (dto.Pagination, error) {

	request.Paginate.SetDefault()

	fmt.Println(color.Red("request.Filter.ID: "), request.Filter.ID)

	users, err := app.repo.GetUsers(ctx, request)
	if err != nil {
		return dto.Pagination{}, err
	}

	return users, nil
}
