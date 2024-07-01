package main

import (
	"github.com/cloutstrife13/ehealth-microservices-2024/services/patient-service/src/patients"
	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("patients-db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	app := echo.New()

	patients.PatientModule{App: app, Db: db}.RegisterModule()

	app.Start(":8080")
}
