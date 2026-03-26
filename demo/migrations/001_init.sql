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

-- Тестовые данные
INSERT INTO services (name, description, category, status) VALUES
    ('GitLab', 'Система контроля версий и CI/CD', 'DevOps', 'active'),
    ('Jira', 'Система управления задачами', 'Project Management', 'active'),
    ('Confluence', 'База знаний и документация', 'Documentation', 'active'),
    ('Grafana', 'Мониторинг и визуализация метрик', 'Monitoring', 'active'),
    ('Vault', 'Управление секретами', 'Security', 'inactive');

CREATE TABLE IF NOT EXISTS service_dependencies (
    id           UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    service_id   UUID NOT NULL REFERENCES services(id) ON DELETE CASCADE,
    depends_on_id UUID NOT NULL REFERENCES services(id) ON DELETE CASCADE,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_dependency UNIQUE(service_id, depends_on_id),
    CONSTRAINT no_self_dependency CHECK (service_id != depends_on_id)
);

-- Тестовые зависимости
INSERT INTO service_dependencies (service_id, depends_on_id)
SELECT s1.id, s2.id FROM services s1, services s2
WHERE s1.name = 'GitLab' AND s2.name = 'Vault';

INSERT INTO service_dependencies (service_id, depends_on_id)
SELECT s1.id, s2.id FROM services s1, services s2
WHERE s1.name = 'Grafana' AND s2.name = 'GitLab';
