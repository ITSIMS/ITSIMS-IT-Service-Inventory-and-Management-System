package service

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"itsims/demo/internal/model"
)

// MockServiceRepository is a mock for repository.ServiceRepository
type MockServiceRepository struct {
	mock.Mock
}

func (m *MockServiceRepository) GetAll(ctx context.Context) ([]model.Service, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.Service), args.Error(1)
}

func (m *MockServiceRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Service, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Service), args.Error(1)
}

func (m *MockServiceRepository) Create(ctx context.Context, req model.CreateServiceRequest) (*model.Service, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Service), args.Error(1)
}

func (m *MockServiceRepository) Update(ctx context.Context, id uuid.UUID, req model.UpdateServiceRequest) (*model.Service, error) {
	args := m.Called(ctx, id, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Service), args.Error(1)
}

func (m *MockServiceRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestGetAll_Success(t *testing.T) {
	mockRepo := new(MockServiceRepository)
	svc := NewServiceImpl(mockRepo)
	ctx := context.Background()

	expected := []model.Service{
		{ID: uuid.New(), Name: "Service 1", Status: "active", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	mockRepo.On("GetAll", ctx).Return(expected, nil)

	result, err := svc.GetAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAll_Error(t *testing.T) {
	mockRepo := new(MockServiceRepository)
	svc := NewServiceImpl(mockRepo)
	ctx := context.Background()

	mockRepo.On("GetAll", ctx).Return([]model.Service{}, errors.New("db error"))

	result, err := svc.GetAll(ctx)
	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestGetByID_Success(t *testing.T) {
	mockRepo := new(MockServiceRepository)
	svc := NewServiceImpl(mockRepo)
	ctx := context.Background()

	id := uuid.New()
	expected := &model.Service{ID: id, Name: "Service 1", Status: "active"}

	mockRepo.On("GetByID", ctx, id).Return(expected, nil)

	result, err := svc.GetByID(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetByID_Error(t *testing.T) {
	mockRepo := new(MockServiceRepository)
	svc := NewServiceImpl(mockRepo)
	ctx := context.Background()

	id := uuid.New()

	mockRepo.On("GetByID", ctx, id).Return(nil, sql.ErrNoRows)

	result, err := svc.GetByID(ctx, id)
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestCreate_Success(t *testing.T) {
	mockRepo := new(MockServiceRepository)
	svc := NewServiceImpl(mockRepo)
	ctx := context.Background()

	req := model.CreateServiceRequest{
		Name:        "New Service",
		Description: "Description",
		Category:    "Category",
		Status:      "active",
	}

	expected := &model.Service{
		ID:     uuid.New(),
		Name:   req.Name,
		Status: "active",
	}

	mockRepo.On("Create", ctx, req).Return(expected, nil)

	result, err := svc.Create(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestCreate_DefaultStatus(t *testing.T) {
	mockRepo := new(MockServiceRepository)
	svc := NewServiceImpl(mockRepo)
	ctx := context.Background()

	req := model.CreateServiceRequest{
		Name:   "New Service",
		Status: "",
	}

	expectedReq := model.CreateServiceRequest{
		Name:   "New Service",
		Status: "active",
	}

	expected := &model.Service{
		ID:     uuid.New(),
		Name:   req.Name,
		Status: "active",
	}

	mockRepo.On("Create", ctx, expectedReq).Return(expected, nil)

	result, err := svc.Create(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestCreate_InvalidStatus(t *testing.T) {
	mockRepo := new(MockServiceRepository)
	svc := NewServiceImpl(mockRepo)
	ctx := context.Background()

	req := model.CreateServiceRequest{
		Name:   "New Service",
		Status: "unknown",
	}

	result, err := svc.Create(ctx, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid status: must be active or inactive")
	mockRepo.AssertNotCalled(t, "Create")
}

func TestCreate_RepoError(t *testing.T) {
	mockRepo := new(MockServiceRepository)
	svc := NewServiceImpl(mockRepo)
	ctx := context.Background()

	req := model.CreateServiceRequest{
		Name:   "New Service",
		Status: "active",
	}

	mockRepo.On("Create", ctx, req).Return(nil, errors.New("db error"))

	result, err := svc.Create(ctx, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdate_Success(t *testing.T) {
	mockRepo := new(MockServiceRepository)
	svc := NewServiceImpl(mockRepo)
	ctx := context.Background()

	id := uuid.New()
	req := model.UpdateServiceRequest{
		Name:   "Updated Service",
		Status: "inactive",
	}

	expected := &model.Service{
		ID:     id,
		Name:   req.Name,
		Status: "inactive",
	}

	mockRepo.On("Update", ctx, id, req).Return(expected, nil)

	result, err := svc.Update(ctx, id, req)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdate_DefaultStatus(t *testing.T) {
	mockRepo := new(MockServiceRepository)
	svc := NewServiceImpl(mockRepo)
	ctx := context.Background()

	id := uuid.New()
	req := model.UpdateServiceRequest{
		Name:   "Updated Service",
		Status: "",
	}

	expectedReq := model.UpdateServiceRequest{
		Name:   "Updated Service",
		Status: "active",
	}

	expected := &model.Service{
		ID:     id,
		Name:   req.Name,
		Status: "active",
	}

	mockRepo.On("Update", ctx, id, expectedReq).Return(expected, nil)

	result, err := svc.Update(ctx, id, req)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdate_InvalidStatus(t *testing.T) {
	mockRepo := new(MockServiceRepository)
	svc := NewServiceImpl(mockRepo)
	ctx := context.Background()

	id := uuid.New()
	req := model.UpdateServiceRequest{
		Name:   "Updated Service",
		Status: "pending",
	}

	result, err := svc.Update(ctx, id, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid status: must be active or inactive")
	mockRepo.AssertNotCalled(t, "Update")
}

func TestUpdate_Error(t *testing.T) {
	mockRepo := new(MockServiceRepository)
	svc := NewServiceImpl(mockRepo)
	ctx := context.Background()

	id := uuid.New()
	req := model.UpdateServiceRequest{
		Name:   "Updated Service",
		Status: "active",
	}

	mockRepo.On("Update", ctx, id, req).Return(nil, sql.ErrNoRows)

	result, err := svc.Update(ctx, id, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestDelete_Success(t *testing.T) {
	mockRepo := new(MockServiceRepository)
	svc := NewServiceImpl(mockRepo)
	ctx := context.Background()

	id := uuid.New()

	mockRepo.On("Delete", ctx, id).Return(nil)

	err := svc.Delete(ctx, id)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDelete_Error(t *testing.T) {
	mockRepo := new(MockServiceRepository)
	svc := NewServiceImpl(mockRepo)
	ctx := context.Background()

	id := uuid.New()

	mockRepo.On("Delete", ctx, id).Return(sql.ErrNoRows)

	err := svc.Delete(ctx, id)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, sql.ErrNoRows))
	mockRepo.AssertExpectations(t)
}
