package dependency

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
)

func newDepTestDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	return db, mock
}

var depCols = []string{"id", "service_id", "depends_on_id", "created_at"}

func TestDepRepo_GetByServiceID_Success(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()

	svcID := uuid.New()
	depID := uuid.New()
	id := uuid.New()
	now := time.Now()

	rows := sqlmock.NewRows(depCols).AddRow(id, svcID, depID, now)
	mock.ExpectQuery(`SELECT id, service_id, depends_on_id, created_at FROM service_dependencies WHERE service_id`).
		WithArgs(svcID).
		WillReturnRows(rows)

	result, err := repo.GetByServiceID(ctx, svcID)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, id, result[0].ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_GetByServiceID_Error(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()

	svcID := uuid.New()
	mock.ExpectQuery(`SELECT id, service_id, depends_on_id, created_at FROM service_dependencies WHERE service_id`).
		WithArgs(svcID).
		WillReturnError(errors.New("query error"))

	result, err := repo.GetByServiceID(ctx, svcID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_GetUsedBy_Success(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()

	svcID := uuid.New()
	depID := uuid.New()
	id := uuid.New()
	now := time.Now()

	rows := sqlmock.NewRows(depCols).AddRow(id, svcID, depID, now)
	mock.ExpectQuery(`SELECT id, service_id, depends_on_id, created_at FROM service_dependencies WHERE depends_on_id`).
		WithArgs(depID).
		WillReturnRows(rows)

	result, err := repo.GetUsedBy(ctx, depID)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_GetUsedBy_Error(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()

	depID := uuid.New()
	mock.ExpectQuery(`SELECT id, service_id, depends_on_id, created_at FROM service_dependencies WHERE depends_on_id`).
		WithArgs(depID).
		WillReturnError(errors.New("query error"))

	result, err := repo.GetUsedBy(ctx, depID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_GetAll_Success(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()

	id1 := uuid.New()
	s1 := uuid.New()
	d1 := uuid.New()
	now := time.Now()

	rows := sqlmock.NewRows(depCols).AddRow(id1, s1, d1, now)
	mock.ExpectQuery(`SELECT id, service_id, depends_on_id, created_at FROM service_dependencies$`).
		WillReturnRows(rows)

	result, err := repo.GetAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_GetAll_Error(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()

	mock.ExpectQuery(`SELECT id, service_id, depends_on_id, created_at FROM service_dependencies$`).
		WillReturnError(errors.New("query error"))

	result, err := repo.GetAll(ctx)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_Create_Success(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()

	svcID := uuid.New()
	depID := uuid.New()
	returnedID := uuid.New()
	now := time.Now()

	rows := sqlmock.NewRows(depCols).AddRow(returnedID, svcID, depID, now)
	mock.ExpectQuery(`INSERT INTO service_dependencies`).
		WillReturnRows(rows)

	result, err := repo.Create(ctx, svcID, depID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, returnedID, result.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_Create_Error(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()

	svcID := uuid.New()
	depID := uuid.New()

	mock.ExpectQuery(`INSERT INTO service_dependencies`).
		WillReturnError(errors.New("insert error"))

	result, err := repo.Create(ctx, svcID, depID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_Delete_Success(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()

	id := uuid.New()
	mock.ExpectExec(`DELETE FROM service_dependencies WHERE id`).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Delete(ctx, id)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_Delete_NotFound(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()

	id := uuid.New()
	mock.ExpectExec(`DELETE FROM service_dependencies WHERE id`).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.Delete(ctx, id)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, sql.ErrNoRows))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_Delete_Error(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()

	id := uuid.New()
	mock.ExpectExec(`DELETE FROM service_dependencies WHERE id`).
		WithArgs(id).
		WillReturnError(errors.New("delete error"))

	err := repo.Delete(ctx, id)
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_GetByPair_Success(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()

	svcID := uuid.New()
	depID := uuid.New()
	retID := uuid.New()
	now := time.Now()

	rows := sqlmock.NewRows(depCols).AddRow(retID, svcID, depID, now)
	mock.ExpectQuery(`SELECT id, service_id, depends_on_id, created_at FROM service_dependencies WHERE service_id`).
		WithArgs(svcID, depID).
		WillReturnRows(rows)

	result, err := repo.GetByPair(ctx, svcID, depID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, retID, result.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_GetByPair_Error(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()

	svcID := uuid.New()
	depID := uuid.New()

	mock.ExpectQuery(`SELECT id, service_id, depends_on_id, created_at FROM service_dependencies WHERE service_id`).
		WithArgs(svcID, depID).
		WillReturnError(sql.ErrNoRows)

	result, err := repo.GetByPair(ctx, svcID, depID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_ExistsByServiceIDAndDependsOnID_True(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()

	svcID := uuid.New()
	depID := uuid.New()

	rows := sqlmock.NewRows([]string{"exists"}).AddRow(true)
	mock.ExpectQuery(`SELECT EXISTS`).
		WithArgs(svcID, depID).
		WillReturnRows(rows)

	exists, err := repo.ExistsByServiceIDAndDependsOnID(ctx, svcID, depID)
	assert.NoError(t, err)
	assert.True(t, exists)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_ExistsByServiceIDAndDependsOnID_False(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()

	svcID := uuid.New()
	depID := uuid.New()

	rows := sqlmock.NewRows([]string{"exists"}).AddRow(false)
	mock.ExpectQuery(`SELECT EXISTS`).
		WithArgs(svcID, depID).
		WillReturnRows(rows)

	exists, err := repo.ExistsByServiceIDAndDependsOnID(ctx, svcID, depID)
	assert.NoError(t, err)
	assert.False(t, exists)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_ExistsByServiceIDAndDependsOnID_Error(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()

	svcID := uuid.New()
	depID := uuid.New()

	mock.ExpectQuery(`SELECT EXISTS`).
		WithArgs(svcID, depID).
		WillReturnError(errors.New("query error"))

	exists, err := repo.ExistsByServiceIDAndDependsOnID(ctx, svcID, depID)
	assert.Error(t, err)
	assert.False(t, exists)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_GetAll_ScanError(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()

	rows := sqlmock.NewRows(depCols).AddRow("bad-uuid", "bad-uuid", "bad-uuid", "bad-time")
	mock.ExpectQuery(`SELECT id, service_id, depends_on_id, created_at FROM service_dependencies$`).
		WillReturnRows(rows)

	result, err := repo.GetAll(ctx)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_GetAll_Empty(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()

	mock.ExpectQuery(`SELECT id, service_id, depends_on_id, created_at FROM service_dependencies$`).
		WillReturnRows(sqlmock.NewRows(depCols))

	result, err := repo.GetAll(ctx)
	assert.NoError(t, err)
	assert.Empty(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_GetAll_RowsError(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()

	aID, bID := uuid.New(), uuid.New()
	d := makeDep(aID, bID)
	// RowError(0) triggers after the first row is iterated, causing rows.Err() to be non-nil
	rows := sqlmock.NewRows(depCols).
		AddRow(d.ID, d.ServiceID, d.DependsOnID, d.CreatedAt).
		RowError(0, errors.New("row iteration error"))
	mock.ExpectQuery(`SELECT id, service_id, depends_on_id, created_at FROM service_dependencies$`).
		WillReturnRows(rows)

	result, err := repo.GetAll(ctx)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_GetByServiceID_ScanError(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()
	svcID := uuid.New()

	rows := sqlmock.NewRows(depCols).AddRow("bad-uuid", "bad-uuid", "bad-uuid", "bad-time")
	mock.ExpectQuery(`SELECT id, service_id, depends_on_id, created_at FROM service_dependencies WHERE service_id`).
		WithArgs(svcID).
		WillReturnRows(rows)

	result, err := repo.GetByServiceID(ctx, svcID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDepRepo_GetByServiceID_RowsError(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()
	svcID := uuid.New()

	d := makeDep(svcID, uuid.New())
	rows := sqlmock.NewRows(depCols).
		AddRow(d.ID, d.ServiceID, d.DependsOnID, d.CreatedAt).
		RowError(0, errors.New("row error"))
	mock.ExpectQuery(`SELECT id, service_id, depends_on_id, created_at FROM service_dependencies WHERE service_id`).
		WithArgs(svcID).
		WillReturnRows(rows)

	result, err := repo.GetByServiceID(ctx, svcID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// errDepResult is a custom driver.Result that returns an error from RowsAffected.
type errDepResult struct{}

func (errDepResult) LastInsertId() (int64, error) { return 0, nil }
func (errDepResult) RowsAffected() (int64, error) { return 0, errors.New("rows affected error") }

func TestDepRepo_Delete_RowsAffectedError(t *testing.T) {
	db, mock := newDepTestDB(t)
	defer db.Close()
	repo := NewPostgresDependencyRepository(db)
	ctx := context.Background()
	id := uuid.New()

	mock.ExpectExec(`DELETE FROM service_dependencies WHERE id`).
		WithArgs(id).
		WillReturnResult(errDepResult{})

	err := repo.Delete(ctx, id)
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
