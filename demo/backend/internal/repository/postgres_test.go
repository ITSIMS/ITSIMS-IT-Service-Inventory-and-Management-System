package repository

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"itsims/demo/internal/model"
)

func newTestDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	return db, mock
}

func TestGetAll_Success(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	id1 := uuid.New()
	id2 := uuid.New()
	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "category", "status", "created_at", "updated_at"}).
		AddRow(id1, "Service 1", "Desc 1", "Cat 1", "active", now, now).
		AddRow(id2, "Service 2", "Desc 2", "Cat 2", "inactive", now, now)

	mock.ExpectQuery(`SELECT id, name, description, category, status, created_at, updated_at FROM services`).
		WillReturnRows(rows)

	services, err := repo.GetAll(ctx, model.ServiceFilter{})
	assert.NoError(t, err)
	assert.Len(t, services, 2)
	assert.Equal(t, id1, services[0].ID)
	assert.Equal(t, "Service 1", services[0].Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAll_WithFilter_Category(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	id1 := uuid.New()
	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "category", "status", "created_at", "updated_at"}).
		AddRow(id1, "Service 1", "Desc 1", "DevOps", "active", now, now)

	mock.ExpectQuery(`SELECT id, name, description, category, status, created_at, updated_at FROM services WHERE category`).
		WithArgs("DevOps").
		WillReturnRows(rows)

	services, err := repo.GetAll(ctx, model.ServiceFilter{Category: "DevOps"})
	assert.NoError(t, err)
	assert.Len(t, services, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAll_WithFilter_Status(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	id1 := uuid.New()
	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "category", "status", "created_at", "updated_at"}).
		AddRow(id1, "Service 1", "Desc 1", "Cat 1", "active", now, now)

	mock.ExpectQuery(`SELECT id, name, description, category, status, created_at, updated_at FROM services WHERE status`).
		WithArgs("active").
		WillReturnRows(rows)

	services, err := repo.GetAll(ctx, model.ServiceFilter{Status: "active"})
	assert.NoError(t, err)
	assert.Len(t, services, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAll_WithFilter_Search(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	id1 := uuid.New()
	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "category", "status", "created_at", "updated_at"}).
		AddRow(id1, "GitLab", "Desc 1", "DevOps", "active", now, now)

	mock.ExpectQuery(`SELECT id, name, description, category, status, created_at, updated_at FROM services WHERE name ILIKE`).
		WithArgs("%Git%").
		WillReturnRows(rows)

	services, err := repo.GetAll(ctx, model.ServiceFilter{Search: "Git"})
	assert.NoError(t, err)
	assert.Len(t, services, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAll_WithFilter_Multiple(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	id1 := uuid.New()
	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "category", "status", "created_at", "updated_at"}).
		AddRow(id1, "GitLab", "Desc 1", "DevOps", "active", now, now)

	mock.ExpectQuery(`SELECT id, name, description, category, status, created_at, updated_at FROM services WHERE`).
		WithArgs("DevOps", "active").
		WillReturnRows(rows)

	services, err := repo.GetAll(ctx, model.ServiceFilter{Category: "DevOps", Status: "active"})
	assert.NoError(t, err)
	assert.Len(t, services, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAll_QueryError(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	mock.ExpectQuery(`SELECT id, name, description, category, status, created_at, updated_at FROM services`).
		WillReturnError(errors.New("query error"))

	services, err := repo.GetAll(ctx, model.ServiceFilter{})
	assert.Error(t, err)
	assert.Nil(t, services)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAll_ScanError(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "category", "status", "created_at", "updated_at"}).
		AddRow("invalid-uuid", "Service 1", "Desc 1", "Cat 1", "active", "not-a-time", "not-a-time")

	mock.ExpectQuery(`SELECT id, name, description, category, status, created_at, updated_at FROM services`).
		WillReturnRows(rows)

	services, err := repo.GetAll(ctx, model.ServiceFilter{})
	assert.Error(t, err)
	assert.Nil(t, services)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByID_Success(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	id := uuid.New()
	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "category", "status", "created_at", "updated_at"}).
		AddRow(id, "Service 1", "Desc 1", "Cat 1", "active", now, now)

	mock.ExpectQuery(`SELECT id, name, description, category, status, created_at, updated_at FROM services WHERE id`).
		WithArgs(id).
		WillReturnRows(rows)

	service, err := repo.GetByID(ctx, id)
	assert.NoError(t, err)
	assert.NotNil(t, service)
	assert.Equal(t, id, service.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByID_NotFound(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	id := uuid.New()

	mock.ExpectQuery(`SELECT id, name, description, category, status, created_at, updated_at FROM services WHERE id`).
		WithArgs(id).
		WillReturnError(sql.ErrNoRows)

	service, err := repo.GetByID(ctx, id)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, sql.ErrNoRows))
	assert.Nil(t, service)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByID_QueryError(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	id := uuid.New()

	mock.ExpectQuery(`SELECT id, name, description, category, status, created_at, updated_at FROM services WHERE id`).
		WithArgs(id).
		WillReturnError(errors.New("connection error"))

	service, err := repo.GetByID(ctx, id)
	assert.Error(t, err)
	assert.Nil(t, service)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAll_RowsError(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	id1 := uuid.New()
	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "category", "status", "created_at", "updated_at"}).
		AddRow(id1, "Service 1", "Desc 1", "Cat 1", "active", now, now).
		RowError(0, errors.New("row error"))

	mock.ExpectQuery(`SELECT id, name, description, category, status, created_at, updated_at FROM services`).
		WillReturnRows(rows)

	services, err := repo.GetAll(ctx, model.ServiceFilter{})
	assert.Error(t, err)
	assert.Nil(t, services)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreate_Success(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	req := model.CreateServiceRequest{
		Name:        "New Service",
		Description: "Description",
		Category:    "Category",
		Status:      "active",
	}

	now := time.Now()
	returnedID := uuid.New()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "category", "status", "created_at", "updated_at"}).
		AddRow(returnedID, req.Name, req.Description, req.Category, req.Status, now, now)

	mock.ExpectQuery(`INSERT INTO services`).
		WillReturnRows(rows)

	service, err := repo.Create(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, service)
	assert.Equal(t, req.Name, service.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreate_QueryError(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	req := model.CreateServiceRequest{
		Name:   "New Service",
		Status: "active",
	}

	mock.ExpectQuery(`INSERT INTO services`).
		WillReturnError(errors.New("insert error"))

	service, err := repo.Create(ctx, req)
	assert.Error(t, err)
	assert.Nil(t, service)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdate_Success(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	id := uuid.New()
	req := model.UpdateServiceRequest{
		Name:        "Updated Service",
		Description: "Updated Description",
		Category:    "Updated Category",
		Status:      "inactive",
	}

	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "category", "status", "created_at", "updated_at"}).
		AddRow(id, req.Name, req.Description, req.Category, req.Status, now, now)

	mock.ExpectQuery(`UPDATE services SET`).
		WillReturnRows(rows)

	service, err := repo.Update(ctx, id, req)
	assert.NoError(t, err)
	assert.NotNil(t, service)
	assert.Equal(t, req.Name, service.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdate_NotFound(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	id := uuid.New()
	req := model.UpdateServiceRequest{
		Name:   "Updated Service",
		Status: "active",
	}

	mock.ExpectQuery(`UPDATE services SET`).
		WillReturnError(sql.ErrNoRows)

	service, err := repo.Update(ctx, id, req)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, sql.ErrNoRows))
	assert.Nil(t, service)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdate_QueryError(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	id := uuid.New()
	req := model.UpdateServiceRequest{
		Name:   "Updated Service",
		Status: "active",
	}

	mock.ExpectQuery(`UPDATE services SET`).
		WillReturnError(errors.New("update error"))

	service, err := repo.Update(ctx, id, req)
	assert.Error(t, err)
	assert.Nil(t, service)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete_Success(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	id := uuid.New()

	mock.ExpectExec(`DELETE FROM services WHERE id`).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Delete(ctx, id)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete_NotFound(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	id := uuid.New()

	mock.ExpectExec(`DELETE FROM services WHERE id`).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.Delete(ctx, id)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, sql.ErrNoRows))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete_QueryError(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	id := uuid.New()

	mock.ExpectExec(`DELETE FROM services WHERE id`).
		WithArgs(id).
		WillReturnError(errors.New("delete error"))

	err := repo.Delete(ctx, id)
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetStats_Success(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM services`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(5))

	statusRows := sqlmock.NewRows([]string{"status", "count"}).
		AddRow("active", 3).
		AddRow("inactive", 2)
	mock.ExpectQuery(`SELECT status, COUNT\(\*\) FROM services GROUP BY status`).
		WillReturnRows(statusRows)

	catRows := sqlmock.NewRows([]string{"category", "count"}).
		AddRow("DevOps", 2).
		AddRow("Monitoring", 1)
	mock.ExpectQuery(`SELECT category, COUNT\(\*\) FROM services GROUP BY category`).
		WillReturnRows(catRows)

	stats, err := repo.GetStats(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, stats)
	assert.Equal(t, 5, stats.Total)
	assert.Len(t, stats.ByStatus, 2)
	assert.Len(t, stats.ByCategory, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetStats_TotalError(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM services`).
		WillReturnError(errors.New("count error"))

	stats, err := repo.GetStats(ctx)
	assert.Error(t, err)
	assert.Nil(t, stats)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetStats_StatusError(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM services`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(5))

	mock.ExpectQuery(`SELECT status, COUNT\(\*\) FROM services GROUP BY status`).
		WillReturnError(errors.New("status query error"))

	stats, err := repo.GetStats(ctx)
	assert.Error(t, err)
	assert.Nil(t, stats)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetStats_CategoryError(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM services`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(5))

	statusRows := sqlmock.NewRows([]string{"status", "count"}).
		AddRow("active", 3)
	mock.ExpectQuery(`SELECT status, COUNT\(\*\) FROM services GROUP BY status`).
		WillReturnRows(statusRows)

	mock.ExpectQuery(`SELECT category, COUNT\(\*\) FROM services GROUP BY category`).
		WillReturnError(errors.New("category query error"))

	stats, err := repo.GetStats(ctx)
	assert.Error(t, err)
	assert.Nil(t, stats)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetStats_StatusScanError(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM services`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(5))

	statusRows := sqlmock.NewRows([]string{"status", "count"}).
		AddRow("active", "not-a-number")
	mock.ExpectQuery(`SELECT status, COUNT\(\*\) FROM services GROUP BY status`).
		WillReturnRows(statusRows)

	stats, err := repo.GetStats(ctx)
	assert.Error(t, err)
	assert.Nil(t, stats)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetStats_CategoryScanError(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM services`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(5))

	statusRows := sqlmock.NewRows([]string{"status", "count"}).
		AddRow("active", 3)
	mock.ExpectQuery(`SELECT status, COUNT\(\*\) FROM services GROUP BY status`).
		WillReturnRows(statusRows)

	catRows := sqlmock.NewRows([]string{"category", "count"}).
		AddRow("DevOps", "not-a-number")
	mock.ExpectQuery(`SELECT category, COUNT\(\*\) FROM services GROUP BY category`).
		WillReturnRows(catRows)

	stats, err := repo.GetStats(ctx)
	assert.Error(t, err)
	assert.Nil(t, stats)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAll_EmptyResult(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	mock.ExpectQuery(`SELECT id, name, description, category, status, created_at, updated_at FROM services`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "category", "status", "created_at", "updated_at"}))

	services, err := repo.GetAll(ctx, model.ServiceFilter{})
	assert.NoError(t, err)
	assert.Empty(t, services)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// errResult is a custom driver.Result that returns an error from RowsAffected.
type errResult struct{}

func (errResult) LastInsertId() (int64, error) { return 0, nil }
func (errResult) RowsAffected() (int64, error) { return 0, errors.New("rows affected error") }

func TestDelete_RowsAffectedError(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()
	id := uuid.New()

	mock.ExpectExec(`DELETE FROM services WHERE id`).
		WithArgs(id).
		WillReturnResult(errResult{})

	err := repo.Delete(ctx, id)
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetStats_StatusRowsError(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM services`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(3))

	statusRows := sqlmock.NewRows([]string{"status", "count"}).
		AddRow("active", 3).
		RowError(0, errors.New("row iteration error"))
	mock.ExpectQuery(`SELECT status, COUNT\(\*\) FROM services GROUP BY status`).
		WillReturnRows(statusRows)

	stats, err := repo.GetStats(ctx)
	assert.Error(t, err)
	assert.Nil(t, stats)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetStats_CategoryRowsError(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM services`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(3))

	statusRows := sqlmock.NewRows([]string{"status", "count"}).
		AddRow("active", 3)
	mock.ExpectQuery(`SELECT status, COUNT\(\*\) FROM services GROUP BY status`).
		WillReturnRows(statusRows)

	catRows := sqlmock.NewRows([]string{"category", "count"}).
		AddRow("DevOps", 2).
		RowError(0, errors.New("row iteration error"))
	mock.ExpectQuery(`SELECT category, COUNT\(\*\) FROM services GROUP BY category`).
		WillReturnRows(catRows)

	stats, err := repo.GetStats(ctx)
	assert.Error(t, err)
	assert.Nil(t, stats)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetStats_EmptyStats(t *testing.T) {
	db, mock := newTestDB(t)
	defer db.Close()

	repo := NewPostgresServiceRepository(db)
	ctx := context.Background()

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM services`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

	mock.ExpectQuery(`SELECT status, COUNT\(\*\) FROM services GROUP BY status`).
		WillReturnRows(sqlmock.NewRows([]string{"status", "count"}))

	mock.ExpectQuery(`SELECT category, COUNT\(\*\) FROM services GROUP BY category`).
		WillReturnRows(sqlmock.NewRows([]string{"category", "count"}))

	stats, err := repo.GetStats(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, stats)
	assert.Equal(t, 0, stats.Total)
	assert.Empty(t, stats.ByStatus)
	assert.Empty(t, stats.ByCategory)
	assert.NoError(t, mock.ExpectationsWereMet())
}
