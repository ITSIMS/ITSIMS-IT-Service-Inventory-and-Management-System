package model

import (
	"time"

	"github.com/google/uuid"
)

type ServiceDependency struct {
	ID          uuid.UUID `json:"id"`
	ServiceID   uuid.UUID `json:"service_id"`
	DependsOnID uuid.UUID `json:"depends_on_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateDependencyRequest struct {
	DependsOnID uuid.UUID `json:"depends_on_id" binding:"required"`
}

// GraphNode представляет узел в графе зависимостей
type GraphNode struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Category string    `json:"category"`
	Status   string    `json:"status"`
}

// GraphEdge представляет ребро в графе зависимостей
type GraphEdge struct {
	ID          uuid.UUID `json:"id"`
	ServiceID   uuid.UUID `json:"service_id"`
	DependsOnID uuid.UUID `json:"depends_on_id"`
}

// DependencyGraph - полный граф зависимостей
type DependencyGraph struct {
	Nodes []GraphNode `json:"nodes"`
	Edges []GraphEdge `json:"edges"`
}

// ServiceDependencies - зависимости конкретного сервиса
type ServiceDependencies struct {
	DependsOn []Service `json:"depends_on"` // сервисы, от которых зависит данный
	UsedBy    []Service `json:"used_by"`    // сервисы, которые зависят от данного
}
