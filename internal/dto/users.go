package dto

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}

type Role string

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)
