# ITSIMS Demo — Каталог IT-сервисов

Демонстрационный модуль для ITSIMS: полный CRUD для каталога IT-сервисов.

## Стек

- **Backend**: Go (Gin) + PostgreSQL
- **Frontend**: React + TypeScript + Vite
- **Инфраструктура**: Docker + docker-compose

## Запуск

```bash
docker-compose up --build
```

- Frontend: http://localhost:3000
- Backend API: http://localhost:8080
- PostgreSQL: localhost:5432

## API

| Метод  | Путь                    | Описание              |
|--------|-------------------------|-----------------------|
| GET    | /api/v1/services        | Список всех сервисов  |
| POST   | /api/v1/services        | Создать сервис        |
| GET    | /api/v1/services/:id    | Получить сервис по ID |
| PUT    | /api/v1/services/:id    | Обновить сервис       |
| DELETE | /api/v1/services/:id    | Удалить сервис        |
| GET    | /health                 | Health check          |

## Тесты

```bash
cd backend
go test ./... -cover
```

Подробный вывод с покрытием:

```bash
cd backend
go test ./... -v -coverprofile=coverage.out
go tool cover -func=coverage.out
```
