package dto

import "github.com/google/uuid"

// user is the data transfer object of user
type User struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}

type Users []User
