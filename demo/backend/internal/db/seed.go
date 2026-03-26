package db

import (
	"database/sql"
	"log"
)

// seedServices — список базовых IT-сервисов для демонстрации.
var seedServices = []struct {
	name        string
	description string
	category    string
	status      string
}{
	// CI/CD
	{"GitLab", "Система контроля версий, CI/CD и DevSecOps", "CI/CD", "active"},
	{"Jenkins", "Сервер автоматизации сборки и деплоя", "CI/CD", "active"},
	{"Docker Registry", "Реестр Docker-образов для хранения артефактов", "CI/CD", "active"},
	// Platform
	{"Kubernetes", "Платформа оркестрации контейнеров", "Platform", "active"},
	{"Vault", "Управление секретами и криптографическими ключами", "Security", "active"},
	// Monitoring
	{"Prometheus", "Сбор и хранение метрик", "Monitoring", "active"},
	{"Grafana", "Визуализация метрик и дашборды", "Monitoring", "active"},
	{"ELK Stack", "Централизованное логирование (Elasticsearch + Logstash + Kibana)", "Monitoring", "inactive"},
	// Infrastructure
	{"PostgreSQL", "Реляционная СУБД", "Infrastructure", "active"},
	{"Redis", "In-memory хранилище данных и брокер сообщений", "Infrastructure", "active"},
	{"RabbitMQ", "Брокер сообщений на основе AMQP", "Infrastructure", "active"},
	// Business
	{"Jira", "Система управления задачами и проектами", "Project Management", "active"},
	{"Confluence", "База знаний и корпоративная документация", "Documentation", "active"},
	{"Keycloak", "Управление идентификацией и доступом (SSO/OIDC)", "Security", "active"},
	{"LDAP", "Служба каталогов и корпоративная аутентификация", "Security", "active"},
}

// seedDependencies — граф зависимостей между сервисами.
// A -> B означает: сервис A зависит от сервиса B (A uses B).
var seedDependencies = [][2]string{
	{"Grafana", "Prometheus"},        // Grafana использует Prometheus как источник данных
	{"Jenkins", "GitLab"},            // Jenkins берёт код из GitLab
	{"Jenkins", "Docker Registry"},   // Jenkins пушит образы в Docker Registry
	{"Kubernetes", "Docker Registry"}, // Kubernetes тянет образы из Registry
	{"ELK Stack", "Kubernetes"},      // ELK Stack работает поверх Kubernetes
	{"Prometheus", "Kubernetes"},     // Prometheus мониторит Kubernetes
	{"Keycloak", "PostgreSQL"},       // Keycloak хранит данные в PostgreSQL
	{"Keycloak", "LDAP"},             // Keycloak федерирует пользователей из LDAP
	{"Confluence", "LDAP"},           // Confluence использует LDAP для аутентификации
	{"Jira", "PostgreSQL"},           // Jira хранит данные в PostgreSQL
	{"Vault", "LDAP"},                // Vault использует LDAP как auth backend
	{"Jenkins", "Vault"},             // Jenkins получает секреты из Vault
}

// Seed заполняет базу тестовыми данными, если таблица пустая.
func Seed(db *sql.DB) error {
	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM services").Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		log.Printf("seed: таблица services уже содержит %d записей, пропускаем", count)
		return nil
	}

	log.Println("seed: заполняем тестовыми данными...")

	// Вставляем сервисы
	nameToID := make(map[string]string)
	for _, s := range seedServices {
		var id string
		err := db.QueryRow(
			`INSERT INTO services (name, description, category, status)
			 VALUES ($1, $2, $3, $4) RETURNING id`,
			s.name, s.description, s.category, s.status,
		).Scan(&id)
		if err != nil {
			return err
		}
		nameToID[s.name] = id
	}
	log.Printf("seed: добавлено %d сервисов", len(seedServices))

	// Вставляем зависимости
	inserted := 0
	for _, dep := range seedDependencies {
		fromID, okFrom := nameToID[dep[0]]
		toID, okTo := nameToID[dep[1]]
		if !okFrom || !okTo {
			log.Printf("seed: сервис не найден (%s → %s), пропускаем", dep[0], dep[1])
			continue
		}
		_, err := db.Exec(
			`INSERT INTO service_dependencies (service_id, depends_on_id) VALUES ($1, $2)
			 ON CONFLICT DO NOTHING`,
			fromID, toID,
		)
		if err != nil {
			return err
		}
		inserted++
	}
	log.Printf("seed: добавлено %d зависимостей", inserted)
	return nil
}
