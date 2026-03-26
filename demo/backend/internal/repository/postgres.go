package repository

import (
	"context"
	"database/sql"
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

func (r *PostgresServiceRepository) GetAll(ctx context.Context) ([]model.Service, error) {
	query := `SELECT id, name, description, category, status, created_at, updated_at FROM services ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query)
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
