package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        int64     `json:"id"`
	UUID      uuid.UUID `json:"uuid"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Role      string    `json:"role" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
