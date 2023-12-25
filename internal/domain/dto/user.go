package dto

import (
	"time"

	"github.com/go-playground/validator"
)

var validate = validator.New()

// user is the data transfer object of user
type User struct {
	ID         string    `json:"id" query:"id" `
	Document   string    `json:"document"  query:"document" validate:"required"`
	FName      string    `json:"f_name" query:"f_name" validate:"required"`
	LName      string    `json:"l_name" query:"l_name" validate:"required"`
	SmartEmail string    `json:"s_email"  query:"s_email" validate:"email"`
	Password   string    `json:"password" query:"password" validate:"required"`
	FkRoleID   int64     `json:"fk_role_id" validate:"required"`
	RoleName   string    `json:"role_name"  query:"role_name"`
	CreatedAt  time.Time `json:"created_at" `
}

func (u *User) Validate() error {
	return validate.Struct(u)
}

// Users is a slice of User
type Users []User

type UsersRequest struct {
	Filter   User
	Paginate Paginate
}

func (u *UsersRequest) Validate() error {
	return validate.Struct(u)
}
