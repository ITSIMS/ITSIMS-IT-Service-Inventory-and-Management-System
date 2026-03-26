package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"itsims/demo/internal/model"
)

type PostgresServiceRepository struct {
	db *sql.DB
}

func NewPostgresServiceRepository(db *sql.DB) *PostgresServiceRepository {
	return &PostgresServiceRepository{db: db}
}

func (r *PostgresServiceRepository) GetAll(ctx context.Context, filter model.ServiceFilter) ([]model.Service, error) {
	var conditions []string
	var args []interface{}
	argIdx := 1

	if filter.Category != "" {
		conditions = append(conditions, fmt.Sprintf("category = $%d", argIdx))
		args = append(args, filter.Category)
		argIdx++
	}
	if filter.Status != "" {
		conditions = append(conditions, fmt.Sprintf("status = $%d", argIdx))
		args = append(args, filter.Status)
		argIdx++
	}
	if filter.Search != "" {
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", argIdx))
		args = append(args, "%"+filter.Search+"%")
		argIdx++
	}

	query := "SELECT id, name, description, category, status, created_at, updated_at FROM services"
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}
	query += " ORDER BY created_at DESC"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []model.Service
	for rows.Next() {
		var s model.Service
		err := rows.Scan(&s.ID, &s.Name, &s.Description, &s.Category, &s.Status, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			return nil, err
		}
		services = append(services, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if services == nil {
		services = []model.Service{}
	}

	return services, nil
}

func (r *PostgresServiceRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Service, error) {
	query := `SELECT id, name, description, category, status, created_at, updated_at FROM services WHERE id = $1`

	var s model.Service
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&s.ID, &s.Name, &s.Description, &s.Category, &s.Status, &s.CreatedAt, &s.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *PostgresServiceRepository) Create(ctx context.Context, req model.CreateServiceRequest) (*model.Service, error) {
	now := time.Now().UTC()
	id := uuid.New()

	query := `INSERT INTO services (id, name, description, category, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, name, description, category, status, created_at, updated_at`

	var s model.Service
	err := r.db.QueryRowContext(ctx, query,
		id, req.Name, req.Description, req.Category, req.Status, now, now,
	).Scan(&s.ID, &s.Name, &s.Description, &s.Category, &s.Status, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *PostgresServiceRepository) Update(ctx context.Context, id uuid.UUID, req model.UpdateServiceRequest) (*model.Service, error) {
	now := time.Now().UTC()

	query := `UPDATE services SET name = $1, description = $2, category = $3, status = $4, updated_at = $5
		WHERE id = $6
		RETURNING id, name, description, category, status, created_at, updated_at`

	var s model.Service
	err := r.db.QueryRowContext(ctx, query,
		req.Name, req.Description, req.Category, req.Status, now, id,
	).Scan(&s.ID, &s.Name, &s.Description, &s.Category, &s.Status, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *PostgresServiceRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM services WHERE id = $1`

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

func (r *PostgresServiceRepository) GetStats(ctx context.Context) (*model.ServiceStats, error) {
	stats := &model.ServiceStats{}

	// Get total
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM services").Scan(&stats.Total)
	if err != nil {
		return nil, err
	}

	// By status
	rowsStatus, err := r.db.QueryContext(ctx, "SELECT status, COUNT(*) FROM services GROUP BY status ORDER BY status")
	if err != nil {
		return nil, err
	}
	defer rowsStatus.Close()

	for rowsStatus.Next() {
		var item model.StatsItem
		if err := rowsStatus.Scan(&item.Key, &item.Count); err != nil {
			return nil, err
		}
		stats.ByStatus = append(stats.ByStatus, item)
	}
	if err := rowsStatus.Err(); err != nil {
		return nil, err
	}

	// By category
	rowsCat, err := r.db.QueryContext(ctx, "SELECT category, COUNT(*) FROM services GROUP BY category ORDER BY category")
	if err != nil {
		return nil, err
	}
	defer rowsCat.Close()

	for rowsCat.Next() {
		var item model.StatsItem
		if err := rowsCat.Scan(&item.Key, &item.Count); err != nil {
			return nil, err
		}
		stats.ByCategory = append(stats.ByCategory, item)
	}
	if err := rowsCat.Err(); err != nil {
		return nil, err
	}

	if stats.ByStatus == nil {
		stats.ByStatus = []model.StatsItem{}
	}
	if stats.ByCategory == nil {
		stats.ByCategory = []model.StatsItem{}
	}

	return stats, nil
}
