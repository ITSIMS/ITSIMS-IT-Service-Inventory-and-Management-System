package handler

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"itsims/demo/internal/model"
)

// MockServiceService is a mock for service.ServiceService
type MockServiceService struct {
	mock.Mock
}

func (m *MockServiceService) GetAll(ctx context.Context, filter model.ServiceFilter) ([]model.Service, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).([]model.Service), args.Error(1)
}

func (m *MockServiceService) GetByID(ctx context.Context, id uuid.UUID) (*model.Service, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Service), args.Error(1)
}

func (m *MockServiceService) Create(ctx context.Context, req model.CreateServiceRequest) (*model.Service, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Service), args.Error(1)
}

func (m *MockServiceService) Update(ctx context.Context, id uuid.UUID, req model.UpdateServiceRequest) (*model.Service, error) {
	args := m.Called(ctx, id, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Service), args.Error(1)
}

func (m *MockServiceService) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockServiceService) GetStats(ctx context.Context) (*model.ServiceStats, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.ServiceStats), args.Error(1)
}

func setupRouter(svc *MockServiceService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := NewHandler(svc)
	h.RegisterRoutes(r)
	return r
}

func TestGetAll_Success(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	services := []model.Service{
		{ID: uuid.New(), Name: "Service 1", Status: "active", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	mockSvc.On("GetAll", mock.Anything, model.ServiceFilter{}).Return(services, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/services", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var result []model.Service
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	mockSvc.AssertExpectations(t)
}

func TestGetAll_WithFilters(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	services := []model.Service{
		{ID: uuid.New(), Name: "GitLab", Status: "active", Category: "DevOps"},
	}

	expectedFilter := model.ServiceFilter{Category: "DevOps", Status: "active", Search: "Git"}
	mockSvc.On("GetAll", mock.Anything, expectedFilter).Return(services, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/services?category=DevOps&status=active&search=Git", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestGetAll_Error(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	mockSvc.On("GetAll", mock.Anything, model.ServiceFilter{}).Return([]model.Service{}, errors.New("db error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/services", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestGetByID_Success(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	id := uuid.New()
	expected := &model.Service{ID: id, Name: "Service 1", Status: "active", CreatedAt: time.Now(), UpdatedAt: time.Now()}

	mockSvc.On("GetByID", mock.Anything, id).Return(expected, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/services/"+id.String(), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestGetByID_NotFound(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	id := uuid.New()

	mockSvc.On("GetByID", mock.Anything, id).Return(nil, sql.ErrNoRows)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/services/"+id.String(), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestGetByID_InvalidID(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/services/not-a-uuid", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockSvc.AssertNotCalled(t, "GetByID")
}

func TestCreate_Success(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	reqBody := model.CreateServiceRequest{
		Name:        "New Service",
		Description: "Description",
		Category:    "Category",
		Status:      "active",
	}

	expected := &model.Service{
		ID:     uuid.New(),
		Name:   reqBody.Name,
		Status: "active",
	}

	mockSvc.On("Create", mock.Anything, reqBody).Return(expected, nil)

	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/services", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestCreate_BadRequest(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/services", bytes.NewBufferString(`{"invalid json`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockSvc.AssertNotCalled(t, "Create")
}

func TestCreate_ServiceError(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	reqBody := model.CreateServiceRequest{
		Name:   "New Service",
		Status: "active",
	}

	mockSvc.On("Create", mock.Anything, reqBody).Return(nil, errors.New("service error"))

	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/services", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestUpdate_Success(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	id := uuid.New()
	reqBody := model.UpdateServiceRequest{
		Name:        "Updated Service",
		Description: "Updated Description",
		Category:    "Updated Category",
		Status:      "inactive",
	}

	expected := &model.Service{
		ID:     id,
		Name:   reqBody.Name,
		Status: "inactive",
	}

	mockSvc.On("Update", mock.Anything, id, reqBody).Return(expected, nil)

	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/services/"+id.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestUpdate_NotFound(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	id := uuid.New()
	reqBody := model.UpdateServiceRequest{
		Name:   "Updated Service",
		Status: "active",
	}

	mockSvc.On("Update", mock.Anything, id, reqBody).Return(nil, sql.ErrNoRows)

	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/services/"+id.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestUpdate_BadRequest(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	id := uuid.New()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/services/"+id.String(), bytes.NewBufferString(`{"invalid json`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockSvc.AssertNotCalled(t, "Update")
}

func TestUpdate_InvalidID(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	reqBody := model.UpdateServiceRequest{
		Name:   "Updated Service",
		Status: "active",
	}

	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/services/not-a-uuid", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockSvc.AssertNotCalled(t, "Update")
}

func TestDelete_Success(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	id := uuid.New()

	mockSvc.On("Delete", mock.Anything, id).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/services/"+id.String(), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestDelete_NotFound(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	id := uuid.New()

	mockSvc.On("Delete", mock.Anything, id).Return(sql.ErrNoRows)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/services/"+id.String(), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestDelete_InvalidID(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/services/not-a-uuid", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockSvc.AssertNotCalled(t, "Delete")
}

func TestDelete_InternalError(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	id := uuid.New()

	mockSvc.On("Delete", mock.Anything, id).Return(errors.New("internal error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/services/"+id.String(), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestGetByID_InternalError(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	id := uuid.New()

	mockSvc.On("GetByID", mock.Anything, id).Return(nil, errors.New("internal error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/services/"+id.String(), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestUpdate_InternalError(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	id := uuid.New()
	reqBody := model.UpdateServiceRequest{
		Name:   "Updated Service",
		Status: "active",
	}

	mockSvc.On("Update", mock.Anything, id, reqBody).Return(nil, errors.New("internal error"))

	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/services/"+id.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestHealthCheck(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetStats_Success(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	expected := &model.ServiceStats{
		Total: 5,
		ByStatus: []model.StatsItem{
			{Key: "active", Count: 3},
			{Key: "inactive", Count: 2},
		},
		ByCategory: []model.StatsItem{
			{Key: "DevOps", Count: 2},
		},
	}

	mockSvc.On("GetStats", mock.Anything).Return(expected, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/stats", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var result model.ServiceStats
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Equal(t, 5, result.Total)
	mockSvc.AssertExpectations(t)
}

func TestGetStats_Error(t *testing.T) {
	mockSvc := new(MockServiceService)
	router := setupRouter(mockSvc)

	mockSvc.On("GetStats", mock.Anything).Return(nil, errors.New("db error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/stats", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestRegisterAPIRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	api := r.Group("/api/v1")
	mockSvc := new(MockServiceService)
	h := NewHandler(mockSvc)
	h.RegisterAPIRoutes(api)

	stats := &model.ServiceStats{
		Total:      2,
		ByStatus:   []model.StatsItem{{Key: "active", Count: 2}},
		ByCategory: []model.StatsItem{{Key: "DevOps", Count: 2}},
	}
	mockSvc.On("GetStats", mock.Anything).Return(stats, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/stats", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}
