package dependency

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"itsims/demo/internal/model"
)

// DependencyRepository defines the interface for dependency data access.
type DependencyRepository interface {
	GetByServiceID(ctx context.Context, serviceID uuid.UUID) ([]model.ServiceDependency, error)
	GetUsedBy(ctx context.Context, serviceID uuid.UUID) ([]model.ServiceDependency, error)
	GetAll(ctx context.Context) ([]model.ServiceDependency, error)
	Create(ctx context.Context, serviceID, dependsOnID uuid.UUID) (*model.ServiceDependency, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetByPair(ctx context.Context, serviceID, dependsOnID uuid.UUID) (*model.ServiceDependency, error)
	ExistsByServiceIDAndDependsOnID(ctx context.Context, serviceID, dependsOnID uuid.UUID) (bool, error)
}

// PostgresDependencyRepository is the Postgres implementation.
type PostgresDependencyRepository struct {
	db *sql.DB
}

func NewPostgresDependencyRepository(db *sql.DB) *PostgresDependencyRepository {
	return &PostgresDependencyRepository{db: db}
}

func scanDependency(row interface {
	Scan(dest ...interface{}) error
}) (*model.ServiceDependency, error) {
	var d model.ServiceDependency
	err := row.Scan(&d.ID, &d.ServiceID, &d.DependsOnID, &d.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *PostgresDependencyRepository) GetByServiceID(ctx context.Context, serviceID uuid.UUID) ([]model.ServiceDependency, error) {
	query := `SELECT id, service_id, depends_on_id, created_at FROM service_dependencies WHERE service_id = $1`
	rows, err := r.db.QueryContext(ctx, query, serviceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanDependencies(rows)
}

func (r *PostgresDependencyRepository) GetUsedBy(ctx context.Context, serviceID uuid.UUID) ([]model.ServiceDependency, error) {
	query := `SELECT id, service_id, depends_on_id, created_at FROM service_dependencies WHERE depends_on_id = $1`
	rows, err := r.db.QueryContext(ctx, query, serviceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanDependencies(rows)
}

func (r *PostgresDependencyRepository) GetAll(ctx context.Context) ([]model.ServiceDependency, error) {
	query := `SELECT id, service_id, depends_on_id, created_at FROM service_dependencies`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanDependencies(rows)
}

func scanDependencies(rows *sql.Rows) ([]model.ServiceDependency, error) {
	var deps []model.ServiceDependency
	for rows.Next() {
		var d model.ServiceDependency
		if err := rows.Scan(&d.ID, &d.ServiceID, &d.DependsOnID, &d.CreatedAt); err != nil {
			return nil, err
		}
		deps = append(deps, d)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if deps == nil {
		deps = []model.ServiceDependency{}
	}
	return deps, nil
}

func (r *PostgresDependencyRepository) Create(ctx context.Context, serviceID, dependsOnID uuid.UUID) (*model.ServiceDependency, error) {
	query := `INSERT INTO service_dependencies (id, service_id, depends_on_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id, service_id, depends_on_id, created_at`
	id := uuid.New()
	now := time.Now().UTC()

	row := r.db.QueryRowContext(ctx, query, id, serviceID, dependsOnID, now)
	return scanDependency(row)
}

func (r *PostgresDependencyRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM service_dependencies WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *PostgresDependencyRepository) GetByPair(ctx context.Context, serviceID, dependsOnID uuid.UUID) (*model.ServiceDependency, error) {
	query := `SELECT id, service_id, depends_on_id, created_at FROM service_dependencies WHERE service_id = $1 AND depends_on_id = $2`
	row := r.db.QueryRowContext(ctx, query, serviceID, dependsOnID)
	return scanDependency(row)
}

func (r *PostgresDependencyRepository) ExistsByServiceIDAndDependsOnID(ctx context.Context, serviceID, dependsOnID uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM service_dependencies WHERE service_id = $1 AND depends_on_id = $2)`
	var exists bool
	err := r.db.QueryRowContext(ctx, query, serviceID, dependsOnID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
