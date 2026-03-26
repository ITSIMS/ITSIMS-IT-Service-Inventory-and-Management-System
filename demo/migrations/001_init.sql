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
