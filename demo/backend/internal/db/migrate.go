package db

import "database/sql"

// Schema — все миграции встроены прямо в бинарник.
// Используем IF NOT EXISTS, поэтому повторный запуск безопасен.
const schema = `
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS services (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        VARCHAR(200) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    category    VARCHAR(100) NOT NULL DEFAULT '',
    status      VARCHAR(20) NOT NULL DEFAULT 'active',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS service_dependencies (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    service_id    UUID NOT NULL REFERENCES services(id) ON DELETE CASCADE,
    depends_on_id UUID NOT NULL REFERENCES services(id) ON DELETE CASCADE,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_dependency UNIQUE(service_id, depends_on_id),
    CONSTRAINT no_self_dependency CHECK (service_id != depends_on_id)
);
`

// Migrate создаёт схему БД. Безопасно запускать при каждом старте.
func Migrate(db *sql.DB) error {
	_, err := db.Exec(schema)
	return err
}
