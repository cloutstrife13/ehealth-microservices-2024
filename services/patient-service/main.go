package main

import (
	"os"

	"github.com/cloutstrife13/ehealth-microservices-2024/services/patient-service/src/patients"
	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Look for the DATABASE_URL environment variable first
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Fallback connection string for local development (standard DSN format)
		dsn = "host=localhost user=postgres password=SecurePassword dbname=patients-db port=5433 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database: " + err.Error())
	}

	app := echo.New()

	patients.PatientModule{App: app, Db: db}.RegisterModule()

	app.Start(":8080")
}
