package main

import (
	"os"

	"github.com/cloutstrife13/ehealth-microservices-2024/services/auth-service/src/auth"
	"github.com/cloutstrife13/ehealth-microservices-2024/services/auth-service/src/users"
	"github.com/labstack/echo"
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

	app := echo.New()

	auth.AuthModule{App: app, Db: db}.RegisterModule()
	users.UserController(app)

	app.Start(":8081")
}
