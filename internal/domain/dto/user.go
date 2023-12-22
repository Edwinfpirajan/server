package dto

import (
	"time"

	"github.com/google/uuid"
)

// user is the data transfer object of user
type User struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Document   string    `json:"document"`
	FName      string    `json:"f_name"`
	LName      string    `json:"l_name"`
	SmartEmail string    `json:"s_email"`
	Password   string    `json:"password"`
	FkRoleID   int       `json:"fk_role_id"`
	RoleName   string    `json:"role_name" gorm:"column:name"`
	CreatedAt  time.Time `json:"created_at"`
}

// Users is a slice of User
type Users []User
