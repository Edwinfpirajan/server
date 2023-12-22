package implementation

import (
	"context"
	"fmt"
	"math"

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
	var (
		users dto.Users
		count int64
	)

	// Set up the SQLBuilder with the initial SELECT and FROM clauses
	ur.builder.Select("id", "document", "f_name", "l_name", "smart_email", "fk_role_id", "created_at").From("users")

	// Apply filters if provided in the request
	if request.Filter.ID != uuid.Nil {
		ur.builder.Where(FilterCondition{"id", "=", request.Filter.ID})
	}
	if request.Filter.Document != "" {
		ur.builder.Where(FilterCondition{"document", "ILIKE", request.Filter.Document})
	}

	// Set up pagination
	ur.builder.SetPaginate(request.Paginate)

	// Execute the query and get the total count
	result := ur.builder.Query().Find(&users)
	if result.Error != nil {
		return dto.Pagination{}, result.Error
	}
	result.Count(&count)

	// Return the pagination information
	return dto.Pagination{
		Page:      request.Paginate.Page,
		Limit:     request.Paginate.Limit,
		TotalPage: int(math.Ceil(float64(count) / float64(request.Paginate.Limit))),
		TotalRows: count,
		Rows:      users,
	}, nil
}

// calculateTotalPages calculates the total number of pages based on the total rows and limit per page
// func calculateTotalPages(totalRows, limit int) int {
// 	if totalRows%limit == 0 {
// 		return totalRows / limit
// 	}
// 	return (totalRows / limit) + 1
// }
