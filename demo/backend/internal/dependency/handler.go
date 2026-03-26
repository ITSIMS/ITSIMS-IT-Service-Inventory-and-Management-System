package dependency

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"itsims/demo/internal/model"
)

type DependencyHandler struct {
	svc DependencyService
}

func NewDependencyHandler(svc DependencyService) *DependencyHandler {
	return &DependencyHandler{svc: svc}
}

func (h *DependencyHandler) RegisterRoutes(api *gin.RouterGroup) {
	api.GET("/services/:id/dependencies", h.GetDependencies)
	api.POST("/services/:id/dependencies", h.AddDependency)
	api.DELETE("/services/:id/dependencies/:dep_id", h.RemoveDependency)
	api.GET("/dependencies/graph", h.GetGraph)
}

func (h *DependencyHandler) GetDependencies(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}

	deps, err := h.svc.GetDependencies(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deps)
}

func (h *DependencyHandler) AddDependency(c *gin.Context) {
	idStr := c.Param("id")
	serviceID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}

	var req model.CreateDependencyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dep, err := h.svc.AddDependency(c.Request.Context(), serviceID, req.DependsOnID)
	if err != nil {
		if errors.Is(err, ErrSelfDependency) || errors.Is(err, ErrDuplicateDependency) || errors.Is(err, ErrCyclicDependency) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dep)
}

func (h *DependencyHandler) RemoveDependency(c *gin.Context) {
	idStr := c.Param("id")
	serviceID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid service UUID"})
		return
	}

	depIDStr := c.Param("dep_id")
	depID, err := uuid.Parse(depIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid dependency UUID"})
		return
	}

	err = h.svc.RemoveDependency(c.Request.Context(), serviceID, depID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "dependency not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *DependencyHandler) GetGraph(c *gin.Context) {
	graph, err := h.svc.GetGraph(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, graph)
}
