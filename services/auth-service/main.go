package main

import (
	"log"
	"os"

	"github.com/cloutstrife13/ehealth-microservices-2024/packages/models"
	"github.com/cloutstrife13/ehealth-microservices-2024/services/auth-service/src/auth"
	"github.com/cloutstrife13/ehealth-microservices-2024/services/auth-service/src/users"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Look for the DATABASE_URL environment variable first
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Fallback connection string for local development (standard DSN format)
		dsn = "host=localhost user=postgres password=SecurePassword dbname=auth-db port=5433 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database: " + err.Error())
	}

	if err := db.AutoMigrate(&models.User{}, &models.VerificationStatus{}); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	app := echo.New()

	// Automatically recover from panics and log incoming HTTP requests
	app.Use(middleware.Recover())
	app.Use(middleware.Logger())

	auth.AuthModule{App: app, Db: db}.RegisterModule()
	users.UserController(app)

	for _, route := range app.Routes() {
		log.Printf("REGISTERED ROUTE: %s -> %s", route.Method, route.Path)
	}

	app.Start(":8081")
}
