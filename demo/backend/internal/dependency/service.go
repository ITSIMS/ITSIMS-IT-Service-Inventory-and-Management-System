package dependency

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"itsims/demo/internal/model"
	servicerepository "itsims/demo/internal/repository"
)

var ErrCyclicDependency = errors.New("cyclic dependency detected")
var ErrSelfDependency = errors.New("service cannot depend on itself")
var ErrDuplicateDependency = errors.New("dependency already exists")

// DependencyService defines the interface for dependency business logic.
type DependencyService interface {
	GetDependencies(ctx context.Context, serviceID uuid.UUID) (*model.ServiceDependencies, error)
	AddDependency(ctx context.Context, serviceID, dependsOnID uuid.UUID) (*model.ServiceDependency, error)
	RemoveDependency(ctx context.Context, serviceID, dependencyID uuid.UUID) error
	GetGraph(ctx context.Context) (*model.DependencyGraph, error)
}

type dependencyServiceImpl struct {
	repo        DependencyRepository
	serviceRepo servicerepository.ServiceRepository
}

func NewDependencyService(repo DependencyRepository, serviceRepo servicerepository.ServiceRepository) DependencyService {
	return &dependencyServiceImpl{
		repo:        repo,
		serviceRepo: serviceRepo,
	}
}

func (s *dependencyServiceImpl) GetDependencies(ctx context.Context, serviceID uuid.UUID) (*model.ServiceDependencies, error) {
	depDeps, err := s.repo.GetByServiceID(ctx, serviceID)
	if err != nil {
		return nil, err
	}

	usedByDeps, err := s.repo.GetUsedBy(ctx, serviceID)
	if err != nil {
		return nil, err
	}

	result := &model.ServiceDependencies{
		DependsOn: []model.Service{},
		UsedBy:    []model.Service{},
	}

	for _, d := range depDeps {
		svc, err := s.serviceRepo.GetByID(ctx, d.DependsOnID)
		if err != nil {
			continue // skip not found
		}
		result.DependsOn = append(result.DependsOn, *svc)
	}

	for _, d := range usedByDeps {
		svc, err := s.serviceRepo.GetByID(ctx, d.ServiceID)
		if err != nil {
			continue // skip not found
		}
		result.UsedBy = append(result.UsedBy, *svc)
	}

	return result, nil
}

func (s *dependencyServiceImpl) AddDependency(ctx context.Context, serviceID, dependsOnID uuid.UUID) (*model.ServiceDependency, error) {
	// 1. Check self-dependency
	if serviceID == dependsOnID {
		return nil, ErrSelfDependency
	}

	// 2. Check duplicate
	exists, err := s.repo.ExistsByServiceIDAndDependsOnID(ctx, serviceID, dependsOnID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrDuplicateDependency
	}

	// 3. Check cycle via DFS
	allDeps, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Build adjacency list: for each dep, service_id -> depends_on_id
	adj := make(map[uuid.UUID][]uuid.UUID)
	for _, d := range allDeps {
		adj[d.ServiceID] = append(adj[d.ServiceID], d.DependsOnID)
	}

	// If we can reach serviceID starting from dependsOnID, adding serviceID->dependsOnID would create a cycle
	if hasCycle(dependsOnID, serviceID, adj) {
		return nil, ErrCyclicDependency
	}

	// 4. Create dependency
	return s.repo.Create(ctx, serviceID, dependsOnID)
}

func (s *dependencyServiceImpl) RemoveDependency(ctx context.Context, serviceID, dependsOnID uuid.UUID) error {
	dep, err := s.repo.GetByPair(ctx, serviceID, dependsOnID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, dep.ID)
}

func (s *dependencyServiceImpl) GetGraph(ctx context.Context) (*model.DependencyGraph, error) {
	allDeps, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	allServices, err := s.serviceRepo.GetAll(ctx, model.ServiceFilter{})
	if err != nil {
		return nil, err
	}

	// Build a set of service IDs participating in dependencies
	participantIDs := make(map[uuid.UUID]bool)
	for _, d := range allDeps {
		participantIDs[d.ServiceID] = true
		participantIDs[d.DependsOnID] = true
	}

	// Build service map
	serviceMap := make(map[uuid.UUID]model.Service)
	for _, svc := range allServices {
		serviceMap[svc.ID] = svc
	}

	graph := &model.DependencyGraph{
		Nodes: []model.GraphNode{},
		Edges: []model.GraphEdge{},
	}

	for id := range participantIDs {
		if svc, ok := serviceMap[id]; ok {
			graph.Nodes = append(graph.Nodes, model.GraphNode{
				ID:       svc.ID,
				Name:     svc.Name,
				Category: svc.Category,
				Status:   svc.Status,
			})
		}
	}

	for _, d := range allDeps {
		graph.Edges = append(graph.Edges, model.GraphEdge{
			ID:          d.ID,
			ServiceID:   d.ServiceID,
			DependsOnID: d.DependsOnID,
		})
	}

	return graph, nil
}

func hasCycle(from, target uuid.UUID, adj map[uuid.UUID][]uuid.UUID) bool {
	visited := make(map[uuid.UUID]bool)
	var dfs func(node uuid.UUID) bool
	dfs = func(node uuid.UUID) bool {
		if node == target {
			return true
		}
		if visited[node] {
			return false
		}
		visited[node] = true
		for _, neighbor := range adj[node] {
			if dfs(neighbor) {
				return true
			}
		}
		return false
	}
	return dfs(from)
}
