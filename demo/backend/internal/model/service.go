package model

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Category    string    `json:"category" db:"category"`
	Status      string    `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreateServiceRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=200"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Status      string `json:"status"`
}

type UpdateServiceRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=200"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Status      string `json:"status"`
}
