package db

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newMock(t *testing.T) (*sqlmock.Sqlmock, func()) {
	t.Helper()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	return &mock, func() { db.Close() }
}

// ---- Migrate ----

func TestMigrate_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("CREATE EXTENSION").WillReturnResult(sqlmock.NewResult(0, 0))

	assert.NoError(t, Migrate(db))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMigrate_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("CREATE EXTENSION").WillReturnError(errors.New("migrate error"))

	assert.Error(t, Migrate(db))
}

// ---- Seed ----

func TestSeed_CountError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT COUNT").WillReturnError(errors.New("count error"))

	assert.Error(t, Seed(db))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSeed_AlreadySeeded(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT COUNT").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(5))

	assert.NoError(t, Seed(db))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSeed_InsertServiceError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT COUNT").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
	mock.ExpectQuery("INSERT INTO services").
		WillReturnError(errors.New("insert service error"))

	assert.Error(t, Seed(db))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSeed_InsertDepError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT COUNT").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

	for range seedServices {
		mock.ExpectQuery("INSERT INTO services").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(uuid.New().String()))
	}
	mock.ExpectExec("INSERT INTO service_dependencies").
		WillReturnError(errors.New("dep insert error"))

	assert.Error(t, Seed(db))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSeed_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT COUNT").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

	for range seedServices {
		mock.ExpectQuery("INSERT INTO services").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(uuid.New().String()))
	}
	for range seedDependencies {
		mock.ExpectExec("INSERT INTO service_dependencies").
			WillReturnResult(sqlmock.NewResult(0, 1))
	}

	assert.NoError(t, Seed(db))
	assert.NoError(t, mock.ExpectationsWereMet())
}
