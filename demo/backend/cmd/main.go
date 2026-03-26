package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"itsims/demo/internal/dependency"
	"itsims/demo/internal/handler"
	"itsims/demo/internal/repository"
	"itsims/demo/internal/service"
)

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func main() {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "itsims")
	dbPassword := getEnv("DB_PASSWORD", "itsims_pass")
	dbName := getEnv("DB_NAME", "itsims")
	port := getEnv("PORT", "8080")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Println("connected to database")

	repo := repository.NewPostgresServiceRepository(db)
	svc := service.NewServiceImpl(repo)
	h := handler.NewHandler(svc)

	depRepo := dependency.NewPostgresDependencyRepository(db)
	depSvc := dependency.NewDependencyService(depRepo, repo)
	depHandler := dependency.NewDependencyHandler(depSvc)

	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	h.RegisterRoutes(r)

	api := r.Group("/api/v1")
	depHandler.RegisterRoutes(api)

	log.Printf("server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
