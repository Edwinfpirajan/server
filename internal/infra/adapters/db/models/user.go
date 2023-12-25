package models

import (
	"time"

	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key"`
	Document   string
	FName      string
	LName      string
	SmartEmail string
	Password   string
	FkRoleID   int64 `gorm:"column:fk_role_id"`
	CreatedAt  time.Time
}

func (u *User) ToDomainDto() dto.User {
	return dto.User{
		ID:         u.ID.String(),
		Document:   u.Document,
		FName:      u.FName,
		LName:      u.LName,
		SmartEmail: u.SmartEmail,
		CreatedAt:  u.CreatedAt,
	}
}

type UserWithRole struct {
	User
	RoleName string
}

func (u *UserWithRole) ToDomainDto() dto.User {
	return dto.User{
		ID:         u.ID.String(),
		Document:   u.Document,
		FName:      u.FName,
		LName:      u.LName,
		SmartEmail: u.SmartEmail,
		RoleName:   u.RoleName,
		CreatedAt:  u.CreatedAt,
	}
}

// Users is a slice of User
type Users []User
