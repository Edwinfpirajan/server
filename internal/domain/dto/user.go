package dto

import "github.com/google/uuid"

// user is the data transfer object of user
type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}

type Users []User
