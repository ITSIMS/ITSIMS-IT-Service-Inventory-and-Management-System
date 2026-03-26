# ITSIMS Demo — Каталог IT-сервисов

Демонстрационный модуль ITSIMS: каталог IT-сервисов с управлением зависимостями и графом.

## Стек

- **Backend**: Go (Gin) + PostgreSQL
- **Frontend**: React + TypeScript + Vite
- **Инфраструктура**: Docker + docker-compose

## Быстрый старт

### Первый запуск / после изменений схемы БД

```bash
cd demo
docker-compose down -v          # удалить старый volume
docker-compose up --build       # пересобрать и запустить
```

> Миграции и тестовые данные (15 сервисов + 12 зависимостей) применяются автоматически при старте бэкенда.

### Повторный запуск (БД уже есть)

```bash
docker-compose up --build
```

### Адреса

| Сервис | URL |
|--------|-----|
| Frontend | http://localhost:3000 |
| Backend API | http://localhost:8080 |
| PostgreSQL | localhost:5432 |

---

## API

### IT-сервисы

| Метод | Путь | Описание |
|-------|------|----------|
| GET | `/api/v1/services` | Список сервисов (фильтры: `?category=X&status=Y&search=Z`) |
| POST | `/api/v1/services` | Создать сервис |
| GET | `/api/v1/services/:id` | Получить сервис по ID |
| PUT | `/api/v1/services/:id` | Обновить сервис |
| DELETE | `/api/v1/services/:id` | Удалить сервис |
| GET | `/api/v1/stats` | Статистика по статусам и категориям |

### Зависимости

| Метод | Путь | Описание |
|-------|------|----------|
| GET | `/api/v1/services/:id/dependencies` | Зависимости сервиса |
| POST | `/api/v1/services/:id/dependencies` | Добавить зависимость |
| DELETE | `/api/v1/services/:id/dependencies/:dep_id` | Удалить зависимость |
| GET | `/api/v1/dependencies/graph` | Полный граф зависимостей |

| GET | `/health` | Health check |

---

## Тестовые данные

При первом запуске автоматически создаются **15 IT-сервисов** из 6 категорий и **12 зависимостей**:

| Категория | Сервисы |
|-----------|---------|
| CI/CD | GitLab, Jenkins, Docker Registry |
| Platform | Kubernetes |
| Security | Vault, Keycloak, LDAP |
| Monitoring | Prometheus, Grafana, ELK Stack |
| Infrastructure | PostgreSQL, Redis, RabbitMQ |
| Project Management | Jira |
| Documentation | Confluence |

---

## Тесты

```bash
cd demo/backend
go test ./internal/... -cover
```

Детальный отчёт:

```bash
go test ./internal/... -v -coverprofile=coverage.out
go tool cover -func=coverage.out
```

### Покрытие

| Пакет | Покрытие |
|-------|---------|
| `internal/service` | 100% |
| `internal/handler` | 98.7% |
| `internal/dependency` | 97.3% |
| `internal/repository` | 93.9% |
| **Итого** | **96.8%** |
