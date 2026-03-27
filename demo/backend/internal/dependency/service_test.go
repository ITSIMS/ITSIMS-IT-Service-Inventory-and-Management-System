package dependency

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

// MockDependencyRepository mocks DependencyRepository
type MockDependencyRepository struct {
	mock.Mock
}

func (m *MockDependencyRepository) GetByServiceID(ctx context.Context, serviceID uuid.UUID) ([]model.ServiceDependency, error) {
	args := m.Called(ctx, serviceID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.ServiceDependency), args.Error(1)
}

func (m *MockDependencyRepository) GetUsedBy(ctx context.Context, serviceID uuid.UUID) ([]model.ServiceDependency, error) {
	args := m.Called(ctx, serviceID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.ServiceDependency), args.Error(1)
}

func (m *MockDependencyRepository) GetAll(ctx context.Context) ([]model.ServiceDependency, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.ServiceDependency), args.Error(1)
}

func (m *MockDependencyRepository) Create(ctx context.Context, serviceID, dependsOnID uuid.UUID) (*model.ServiceDependency, error) {
	args := m.Called(ctx, serviceID, dependsOnID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.ServiceDependency), args.Error(1)
}

func (m *MockDependencyRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockDependencyRepository) GetByPair(ctx context.Context, serviceID, dependsOnID uuid.UUID) (*model.ServiceDependency, error) {
	args := m.Called(ctx, serviceID, dependsOnID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.ServiceDependency), args.Error(1)
}

func (m *MockDependencyRepository) ExistsByServiceIDAndDependsOnID(ctx context.Context, serviceID, dependsOnID uuid.UUID) (bool, error) {
	args := m.Called(ctx, serviceID, dependsOnID)
	return args.Bool(0), args.Error(1)
}

// MockServiceRepository mocks repository.ServiceRepository
type MockServiceRepository struct {
	mock.Mock
}

func (m *MockServiceRepository) GetAll(ctx context.Context, filter model.ServiceFilter) ([]model.Service, error) {
	args := m.Called(ctx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
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

func (m *MockServiceRepository) GetStats(ctx context.Context) (*model.ServiceStats, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.ServiceStats), args.Error(1)
}

// Helpers
func makeService(id uuid.UUID, name string) model.Service {
	return model.Service{ID: id, Name: name, Category: "test", Status: "active", CreatedAt: time.Now(), UpdatedAt: time.Now()}
}

func makeDep(serviceID, dependsOnID uuid.UUID) model.ServiceDependency {
	return model.ServiceDependency{
		ID:          uuid.New(),
		ServiceID:   serviceID,
		DependsOnID: dependsOnID,
		CreatedAt:   time.Now(),
	}
}

// Tests

func TestGetDependencies_Success(t *testing.T) {
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	aID := uuid.New()
	bID := uuid.New()
	cID := uuid.New()

	depRepo.On("GetByServiceID", ctx, aID).Return([]model.ServiceDependency{makeDep(aID, bID)}, nil)
	depRepo.On("GetUsedBy", ctx, aID).Return([]model.ServiceDependency{makeDep(cID, aID)}, nil)
	svcRepo.On("GetByID", ctx, bID).Return(&model.Service{ID: bID, Name: "B"}, nil)
	svcRepo.On("GetByID", ctx, cID).Return(&model.Service{ID: cID, Name: "C"}, nil)

	result, err := svc.GetDependencies(ctx, aID)
	assert.NoError(t, err)
	assert.Len(t, result.DependsOn, 1)
	assert.Equal(t, bID, result.DependsOn[0].ID)
	assert.Len(t, result.UsedBy, 1)
	assert.Equal(t, cID, result.UsedBy[0].ID)
	depRepo.AssertExpectations(t)
	svcRepo.AssertExpectations(t)
}

func TestGetDependencies_ErrorOnDependsOn(t *testing.T) {
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	aID := uuid.New()

	depRepo.On("GetByServiceID", ctx, aID).Return(nil, errors.New("repo error"))

	result, err := svc.GetDependencies(ctx, aID)
	assert.Error(t, err)
	assert.Nil(t, result)
	depRepo.AssertExpectations(t)
}

func TestGetDependencies_ErrorOnUsedBy(t *testing.T) {
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	aID := uuid.New()

	depRepo.On("GetByServiceID", ctx, aID).Return([]model.ServiceDependency{}, nil)
	depRepo.On("GetUsedBy", ctx, aID).Return(nil, errors.New("repo error"))

	result, err := svc.GetDependencies(ctx, aID)
	assert.Error(t, err)
	assert.Nil(t, result)
	depRepo.AssertExpectations(t)
}

func TestAddDependency_SelfDependency(t *testing.T) {
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	id := uuid.New()
	result, err := svc.AddDependency(ctx, id, id)
	assert.ErrorIs(t, err, ErrSelfDependency)
	assert.Nil(t, result)
}

func TestAddDependency_DuplicateDependency(t *testing.T) {
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	aID := uuid.New()
	bID := uuid.New()

	depRepo.On("ExistsByServiceIDAndDependsOnID", ctx, aID, bID).Return(true, nil)

	result, err := svc.AddDependency(ctx, aID, bID)
	assert.ErrorIs(t, err, ErrDuplicateDependency)
	assert.Nil(t, result)
	depRepo.AssertExpectations(t)
}

func TestAddDependency_CyclicDependency(t *testing.T) {
	// A->B exists, try to add B->A
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	aID := uuid.New()
	bID := uuid.New()

	// B->A: serviceID=bID, dependsOnID=aID
	depRepo.On("ExistsByServiceIDAndDependsOnID", ctx, bID, aID).Return(false, nil)
	// GetAll returns A->B
	depRepo.On("GetAll", ctx).Return([]model.ServiceDependency{makeDep(aID, bID)}, nil)

	result, err := svc.AddDependency(ctx, bID, aID)
	assert.ErrorIs(t, err, ErrCyclicDependency)
	assert.Nil(t, result)
	depRepo.AssertExpectations(t)
}

func TestAddDependency_TransitiveCycle(t *testing.T) {
	// A->B->C exist, try to add C->A
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	aID := uuid.New()
	bID := uuid.New()
	cID := uuid.New()

	depRepo.On("ExistsByServiceIDAndDependsOnID", ctx, cID, aID).Return(false, nil)
	depRepo.On("GetAll", ctx).Return([]model.ServiceDependency{
		makeDep(aID, bID), // A->B
		makeDep(bID, cID), // B->C
	}, nil)

	result, err := svc.AddDependency(ctx, cID, aID)
	assert.ErrorIs(t, err, ErrCyclicDependency)
	assert.Nil(t, result)
	depRepo.AssertExpectations(t)
}

func TestAddDependency_Success(t *testing.T) {
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	aID := uuid.New()
	bID := uuid.New()

	created := &model.ServiceDependency{ID: uuid.New(), ServiceID: aID, DependsOnID: bID, CreatedAt: time.Now()}

	depRepo.On("ExistsByServiceIDAndDependsOnID", ctx, aID, bID).Return(false, nil)
	depRepo.On("GetAll", ctx).Return([]model.ServiceDependency{}, nil)
	depRepo.On("Create", ctx, aID, bID).Return(created, nil)

	result, err := svc.AddDependency(ctx, aID, bID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, created.ID, result.ID)
	depRepo.AssertExpectations(t)
}

func TestAddDependency_GetAllError(t *testing.T) {
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	aID := uuid.New()
	bID := uuid.New()

	depRepo.On("ExistsByServiceIDAndDependsOnID", ctx, aID, bID).Return(false, nil)
	depRepo.On("GetAll", ctx).Return(nil, errors.New("db error"))

	result, err := svc.AddDependency(ctx, aID, bID)
	assert.Error(t, err)
	assert.Nil(t, result)
	depRepo.AssertExpectations(t)
}

func TestAddDependency_ExistsError(t *testing.T) {
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	aID := uuid.New()
	bID := uuid.New()

	depRepo.On("ExistsByServiceIDAndDependsOnID", ctx, aID, bID).Return(false, errors.New("db error"))

	result, err := svc.AddDependency(ctx, aID, bID)
	assert.Error(t, err)
	assert.Nil(t, result)
	depRepo.AssertExpectations(t)
}

func TestAddDependency_CreateError(t *testing.T) {
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	aID := uuid.New()
	bID := uuid.New()

	depRepo.On("ExistsByServiceIDAndDependsOnID", ctx, aID, bID).Return(false, nil)
	depRepo.On("GetAll", ctx).Return([]model.ServiceDependency{}, nil)
	depRepo.On("Create", ctx, aID, bID).Return(nil, errors.New("insert error"))

	result, err := svc.AddDependency(ctx, aID, bID)
	assert.Error(t, err)
	assert.Nil(t, result)
	depRepo.AssertExpectations(t)
}

func TestRemoveDependency_Success(t *testing.T) {
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	svcID := uuid.New()
	dependsOnID := uuid.New()
	recordID := uuid.New()

	depRecord := &model.ServiceDependency{
		ID:          recordID,
		ServiceID:   svcID,
		DependsOnID: dependsOnID,
	}
	depRepo.On("GetByPair", ctx, svcID, dependsOnID).Return(depRecord, nil)
	depRepo.On("Delete", ctx, recordID).Return(nil)

	err := svc.RemoveDependency(ctx, svcID, dependsOnID)
	assert.NoError(t, err)
	depRepo.AssertExpectations(t)
}

func TestRemoveDependency_GetByPairError(t *testing.T) {
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	svcID := uuid.New()
	dependsOnID := uuid.New()

	depRepo.On("GetByPair", ctx, svcID, dependsOnID).Return(nil, sql.ErrNoRows)

	err := svc.RemoveDependency(ctx, svcID, dependsOnID)
	assert.ErrorIs(t, err, sql.ErrNoRows)
	depRepo.AssertExpectations(t)
}

func TestRemoveDependency_DeleteError(t *testing.T) {
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	svcID := uuid.New()
	dependsOnID := uuid.New()
	recordID := uuid.New()

	depRecord := &model.ServiceDependency{
		ID:          recordID,
		ServiceID:   svcID,
		DependsOnID: dependsOnID,
	}
	depRepo.On("GetByPair", ctx, svcID, dependsOnID).Return(depRecord, nil)
	depRepo.On("Delete", ctx, recordID).Return(sql.ErrNoRows)

	err := svc.RemoveDependency(ctx, svcID, dependsOnID)
	assert.ErrorIs(t, err, sql.ErrNoRows)
	depRepo.AssertExpectations(t)
}

func TestGetGraph_Success(t *testing.T) {
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	aID := uuid.New()
	bID := uuid.New()

	svcA := makeService(aID, "ServiceA")
	svcB := makeService(bID, "ServiceB")

	depRepo.On("GetAll", ctx).Return([]model.ServiceDependency{makeDep(aID, bID)}, nil)
	svcRepo.On("GetAll", ctx, model.ServiceFilter{}).Return([]model.Service{svcA, svcB}, nil)

	graph, err := svc.GetGraph(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, graph)
	assert.Len(t, graph.Edges, 1)
	assert.Len(t, graph.Nodes, 2)
	depRepo.AssertExpectations(t)
	svcRepo.AssertExpectations(t)
}

func TestGetGraph_Error(t *testing.T) {
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	depRepo.On("GetAll", ctx).Return(nil, errors.New("db error"))

	graph, err := svc.GetGraph(ctx)
	assert.Error(t, err)
	assert.Nil(t, graph)
	depRepo.AssertExpectations(t)
}

func TestGetGraph_ServiceRepoError(t *testing.T) {
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	aID := uuid.New()
	bID := uuid.New()

	depRepo.On("GetAll", ctx).Return([]model.ServiceDependency{makeDep(aID, bID)}, nil)
	svcRepo.On("GetAll", ctx, model.ServiceFilter{}).Return(nil, errors.New("service repo error"))

	graph, err := svc.GetGraph(ctx)
	assert.Error(t, err)
	assert.Nil(t, graph)
	depRepo.AssertExpectations(t)
	svcRepo.AssertExpectations(t)
}

func TestGetDependencies_ServiceNotFound_Skipped(t *testing.T) {
	// Verify that if GetByID fails for a dep, it's skipped (not an error)
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	aID := uuid.New()
	bID := uuid.New()

	depRepo.On("GetByServiceID", ctx, aID).Return([]model.ServiceDependency{makeDep(aID, bID)}, nil)
	depRepo.On("GetUsedBy", ctx, aID).Return([]model.ServiceDependency{}, nil)
	svcRepo.On("GetByID", ctx, bID).Return(nil, sql.ErrNoRows)

	result, err := svc.GetDependencies(ctx, aID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.DependsOn, 0) // skipped
	depRepo.AssertExpectations(t)
	svcRepo.AssertExpectations(t)
}

func TestGetDependencies_UsedByServiceNotFound_Skipped(t *testing.T) {
	// Verify that if GetByID fails for a usedBy dep, it's skipped (not an error)
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	aID := uuid.New()
	cID := uuid.New()

	depRepo.On("GetByServiceID", ctx, aID).Return([]model.ServiceDependency{}, nil)
	depRepo.On("GetUsedBy", ctx, aID).Return([]model.ServiceDependency{makeDep(cID, aID)}, nil)
	svcRepo.On("GetByID", ctx, cID).Return(nil, sql.ErrNoRows)

	result, err := svc.GetDependencies(ctx, aID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.UsedBy, 0) // skipped
	depRepo.AssertExpectations(t)
	svcRepo.AssertExpectations(t)
}

func TestAddDependency_NoCycleWithSharedDep(t *testing.T) {
	// Diamond graph: A->B, A->C, B->D, C->D
	// Add E->A: hasCycle(from=A, target=E) — DFS visits D twice; second visit hits the visited branch
	depRepo := new(MockDependencyRepository)
	svcRepo := new(MockServiceRepository)
	svc := NewDependencyService(depRepo, svcRepo)
	ctx := context.Background()

	aID := uuid.New()
	bID := uuid.New()
	cID := uuid.New()
	dID := uuid.New()
	eID := uuid.New()

	depRepo.On("ExistsByServiceIDAndDependsOnID", ctx, eID, aID).Return(false, nil)
	depRepo.On("GetAll", ctx).Return([]model.ServiceDependency{
		makeDep(aID, bID), // A->B
		makeDep(aID, cID), // A->C
		makeDep(bID, dID), // B->D
		makeDep(cID, dID), // C->D  (diamond — D visited twice)
	}, nil)

	created := &model.ServiceDependency{ID: uuid.New(), ServiceID: eID, DependsOnID: aID, CreatedAt: time.Now()}
	depRepo.On("Create", ctx, eID, aID).Return(created, nil)

	result, err := svc.AddDependency(ctx, eID, aID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	depRepo.AssertExpectations(t)
}
