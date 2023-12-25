package dto

import (
	"time"

	"github.com/Edwinfpirajan/server.git/utils/db"
)

type Collaborator struct {
	Id             int    `json:"id" query:"id"`
	Document       string `json:"document"  query:"document"`
	FName          string
	LName          string
	Email          string
	Bmail          string
	State          string
	Leader         string
	LeaderDocument string
	Subprocess     string
	Headquarters   string
	Position       string
	CreatedAt      time.Time
}

type Collaborators []Collaborator

func (c *Collaborator) Validate() error {
	return db.Validate.Struct(c)
}
