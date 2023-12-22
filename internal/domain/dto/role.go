package dto

// Role is the data transfer object of role
type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Roles is a slice of Role
type Roles []Role
