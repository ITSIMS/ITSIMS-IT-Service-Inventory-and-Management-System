package repository

import (
	"context"

	"github.com/google/uuid"
	"itsims/demo/internal/model"
)

type ServiceRepository interface {
	GetAll(ctx context.Context) ([]model.Service, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Service, error)
	Create(ctx context.Context, req model.CreateServiceRequest) (*model.Service, error)
	Update(ctx context.Context, id uuid.UUID, req model.UpdateServiceRequest) (*model.Service, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
