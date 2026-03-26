package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"itsims/demo/internal/model"
	"itsims/demo/internal/repository"
)

type serviceImpl struct {
	repo repository.ServiceRepository
}

func NewServiceImpl(repo repository.ServiceRepository) ServiceService {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) GetAll(ctx context.Context) ([]model.Service, error) {
	return s.repo.GetAll(ctx)
}

func (s *serviceImpl) GetByID(ctx context.Context, id uuid.UUID) (*model.Service, error) {
	return s.repo.GetByID(ctx, id)
}

func validateStatus(status string) error {
	if status != "active" && status != "inactive" {
		return errors.New("invalid status: must be active or inactive")
	}
	return nil
}

func (s *serviceImpl) Create(ctx context.Context, req model.CreateServiceRequest) (*model.Service, error) {
	if req.Status == "" {
		req.Status = "active"
	}

	if err := validateStatus(req.Status); err != nil {
		return nil, err
	}

	return s.repo.Create(ctx, req)
}

func (s *serviceImpl) Update(ctx context.Context, id uuid.UUID, req model.UpdateServiceRequest) (*model.Service, error) {
	if req.Status == "" {
		req.Status = "active"
	}

	if err := validateStatus(req.Status); err != nil {
		return nil, err
	}

	return s.repo.Update(ctx, id, req)
}

func (s *serviceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
