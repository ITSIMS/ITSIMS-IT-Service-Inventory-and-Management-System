package dependency

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

// MockDependencyService mocks DependencyService
type MockDependencyService struct {
	mock.Mock
}

func (m *MockDependencyService) GetDependencies(ctx context.Context, serviceID uuid.UUID) (*model.ServiceDependencies, error) {
	args := m.Called(ctx, serviceID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.ServiceDependencies), args.Error(1)
}

func (m *MockDependencyService) AddDependency(ctx context.Context, serviceID, dependsOnID uuid.UUID) (*model.ServiceDependency, error) {
	args := m.Called(ctx, serviceID, dependsOnID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.ServiceDependency), args.Error(1)
}

func (m *MockDependencyService) RemoveDependency(ctx context.Context, serviceID, dependencyID uuid.UUID) error {
	args := m.Called(ctx, serviceID, dependencyID)
	return args.Error(0)
}

func (m *MockDependencyService) GetGraph(ctx context.Context) (*model.DependencyGraph, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.DependencyGraph), args.Error(1)
}

func setupDepRouter(svc *MockDependencyService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := NewDependencyHandler(svc)
	api := r.Group("/api/v1")
	h.RegisterRoutes(api)
	return r
}

func TestDepHandler_GetDependencies_Success(t *testing.T) {
	mockSvc := new(MockDependencyService)
	router := setupDepRouter(mockSvc)

	id := uuid.New()
	expected := &model.ServiceDependencies{
		DependsOn: []model.Service{{ID: uuid.New(), Name: "B"}},
		UsedBy:    []model.Service{},
	}

	mockSvc.On("GetDependencies", mock.Anything, id).Return(expected, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/services/"+id.String()+"/dependencies", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestDepHandler_GetDependencies_InvalidUUID(t *testing.T) {
	mockSvc := new(MockDependencyService)
	router := setupDepRouter(mockSvc)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/services/bad-uuid/dependencies", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockSvc.AssertNotCalled(t, "GetDependencies")
}

func TestDepHandler_GetDependencies_NotFound(t *testing.T) {
	mockSvc := new(MockDependencyService)
	router := setupDepRouter(mockSvc)

	id := uuid.New()
	mockSvc.On("GetDependencies", mock.Anything, id).Return(nil, sql.ErrNoRows)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/services/"+id.String()+"/dependencies", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestDepHandler_GetDependencies_InternalError(t *testing.T) {
	mockSvc := new(MockDependencyService)
	router := setupDepRouter(mockSvc)

	id := uuid.New()
	mockSvc.On("GetDependencies", mock.Anything, id).Return(nil, errors.New("internal error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/services/"+id.String()+"/dependencies", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestDepHandler_AddDependency_Success(t *testing.T) {
	mockSvc := new(MockDependencyService)
	router := setupDepRouter(mockSvc)

	serviceID := uuid.New()
	dependsOnID := uuid.New()
	created := &model.ServiceDependency{
		ID:          uuid.New(),
		ServiceID:   serviceID,
		DependsOnID: dependsOnID,
		CreatedAt:   time.Now(),
	}

	mockSvc.On("AddDependency", mock.Anything, serviceID, dependsOnID).Return(created, nil)

	reqBody := model.CreateDependencyRequest{DependsOnID: dependsOnID}
	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/services/"+serviceID.String()+"/dependencies", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestDepHandler_AddDependency_InvalidUUID(t *testing.T) {
	mockSvc := new(MockDependencyService)
	router := setupDepRouter(mockSvc)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/services/bad-uuid/dependencies", bytes.NewBufferString(`{}`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockSvc.AssertNotCalled(t, "AddDependency")
}

func TestDepHandler_AddDependency_BadJSON(t *testing.T) {
	mockSvc := new(MockDependencyService)
	router := setupDepRouter(mockSvc)

	id := uuid.New()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/services/"+id.String()+"/dependencies", bytes.NewBufferString(`{bad json`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockSvc.AssertNotCalled(t, "AddDependency")
}

func TestDepHandler_AddDependency_SelfDependency(t *testing.T) {
	mockSvc := new(MockDependencyService)
	router := setupDepRouter(mockSvc)

	id := uuid.New()
	mockSvc.On("AddDependency", mock.Anything, id, id).Return(nil, ErrSelfDependency)

	reqBody := model.CreateDependencyRequest{DependsOnID: id}
	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/services/"+id.String()+"/dependencies", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestDepHandler_AddDependency_Duplicate(t *testing.T) {
	mockSvc := new(MockDependencyService)
	router := setupDepRouter(mockSvc)

	svcID := uuid.New()
	depID := uuid.New()
	mockSvc.On("AddDependency", mock.Anything, svcID, depID).Return(nil, ErrDuplicateDependency)

	reqBody := model.CreateDependencyRequest{DependsOnID: depID}
	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/services/"+svcID.String()+"/dependencies", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestDepHandler_AddDependency_CyclicDependency(t *testing.T) {
	mockSvc := new(MockDependencyService)
	router := setupDepRouter(mockSvc)

	svcID := uuid.New()
	depID := uuid.New()
	mockSvc.On("AddDependency", mock.Anything, svcID, depID).Return(nil, ErrCyclicDependency)

	reqBody := model.CreateDependencyRequest{DependsOnID: depID}
	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/services/"+svcID.String()+"/dependencies", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestDepHandler_AddDependency_InternalError(t *testing.T) {
	mockSvc := new(MockDependencyService)
	router := setupDepRouter(mockSvc)

	svcID := uuid.New()
	depID := uuid.New()
	mockSvc.On("AddDependency", mock.Anything, svcID, depID).Return(nil, errors.New("internal error"))

	reqBody := model.CreateDependencyRequest{DependsOnID: depID}
	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/services/"+svcID.String()+"/dependencies", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestDepHandler_RemoveDependency_Success(t *testing.T) {
	mockSvc := new(MockDependencyService)
	router := setupDepRouter(mockSvc)

	svcID := uuid.New()
	depID := uuid.New()
	mockSvc.On("RemoveDependency", mock.Anything, svcID, depID).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/services/"+svcID.String()+"/dependencies/"+depID.String(), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestDepHandler_RemoveDependency_InvalidServiceUUID(t *testing.T) {
	mockSvc := new(MockDependencyService)
	router := setupDepRouter(mockSvc)

	depID := uuid.New()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/services/bad-uuid/dependencies/"+depID.String(), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockSvc.AssertNotCalled(t, "RemoveDependency")
}

func TestDepHandler_RemoveDependency_InvalidDepUUID(t *testing.T) {
	mockSvc := new(MockDependencyService)
	router := setupDepRouter(mockSvc)

	svcID := uuid.New()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/services/"+svcID.String()+"/dependencies/bad-dep-uuid", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockSvc.AssertNotCalled(t, "RemoveDependency")
}

func TestDepHandler_RemoveDependency_NotFound(t *testing.T) {
	mockSvc := new(MockDependencyService)
	router := setupDepRouter(mockSvc)

	svcID := uuid.New()
	depID := uuid.New()
	mockSvc.On("RemoveDependency", mock.Anything, svcID, depID).Return(sql.ErrNoRows)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/services/"+svcID.String()+"/dependencies/"+depID.String(), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestDepHandler_RemoveDependency_InternalError(t *testing.T) {
	mockSvc := new(MockDependencyService)
	router := setupDepRouter(mockSvc)

	svcID := uuid.New()
	depID := uuid.New()
	mockSvc.On("RemoveDependency", mock.Anything, svcID, depID).Return(errors.New("internal error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/services/"+svcID.String()+"/dependencies/"+depID.String(), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestDepHandler_GetGraph_Success(t *testing.T) {
	mockSvc := new(MockDependencyService)
	router := setupDepRouter(mockSvc)

	graph := &model.DependencyGraph{
		Nodes: []model.GraphNode{{ID: uuid.New(), Name: "A"}},
		Edges: []model.GraphEdge{{ID: uuid.New(), ServiceID: uuid.New(), DependsOnID: uuid.New()}},
	}
	mockSvc.On("GetGraph", mock.Anything).Return(graph, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/dependencies/graph", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestDepHandler_GetGraph_Error(t *testing.T) {
	mockSvc := new(MockDependencyService)
	router := setupDepRouter(mockSvc)

	mockSvc.On("GetGraph", mock.Anything).Return(nil, errors.New("graph error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/dependencies/graph", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockSvc.AssertExpectations(t)
}
