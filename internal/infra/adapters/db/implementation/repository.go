package implementation

import (
	"context"
	"fmt"
	"math"

	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
	"github.com/Edwinfpirajan/server.git/internal/domain/ports/db/interfaces"
	"github.com/Edwinfpirajan/server.git/internal/infra/adapters/db/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) interfaces.Repository {
	return &repository{
		db,
	}
}

// Users
func (repo *repository) CreateUser(ctx context.Context, user models.User) error {
	if err := repo.db.Table("users").Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (repo *repository) GetUser(ctx context.Context, request dto.User) (dto.User, error) {
	var (
		modelUser models.UserWithRole
	)

	if err := repo.db.Table("users u").
		Select("u.id", "u.document", "u.f_name", "u.l_name", "u.smart_email", "u.created_at", "r.name as role_name").
		Where("u.document = ? OR u.smart_email = ?", request.Document, request.SmartEmail).
		Joins("JOIN roles r ON r.id = u.fk_role_id").
		Scan(&modelUser).Error; err != nil {
		return dto.User{}, err
	}

	fmt.Printf("%+v", modelUser)

	return modelUser.ToDomainDto(), nil
}

// GetUsers is the implementation of NewUserRepository.GetUsers
func (repo *repository) GetUsers(ctx context.Context, request dto.UsersRequest) (dto.Pagination, error) {
	var (
		users dto.Users
		count int64
		where string
	)

	offset := (request.Paginate.Page - 1) * request.Paginate.Limit

	BuildFilters("document", request.Filter.Document, "OR", &where)
	BuildFilters("f_name", request.Filter.FName, "OR", &where)

	if err := repo.db.Table("users u").
		Select("u.id", "u.document", "u.f_name", "u.l_name", "u.smart_email", "u.created_at", "r.name as role_name").
		Where(where).
		Joins("JOIN roles r ON r.id = u.fk_role_id").
		Offset(offset).
		Limit(request.Paginate.Limit).
		Scan(&users).Error; err != nil {
		return dto.Pagination{}, err
	}

	return dto.Pagination{
		Page:      request.Paginate.Page,
		Limit:     request.Paginate.Limit,
		TotalPage: int(math.Ceil(float64(count) / float64(request.Paginate.Limit))),
		TotalRows: count,
		Rows:      users}, nil
}

// Roles
func (repo *repository) CreateRole(ctx context.Context, role dto.Role) (dto.Role, error) {
	result := repo.db.Select("roles").Create(&dto.Role{
		Name: role.Name,
	})

	if result.Error != nil {
		return dto.Role{}, result.Error
	}

	return dto.Role{}, nil
}

func (repo *repository) GetRoles(ctx context.Context) (dto.Roles, error) {
	var roles dto.Roles

	result := repo.db.Select("id", "name").Find(&roles)

	if result.Error != nil {
		return nil, result.Error
	}

	return roles, nil
}
