package main

import (
	"github.com/cloutstrife13/ehealth-microservices-2024/services/patient-service/patients"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	patients.PatientController(app)

	app.Listen(":8080")
}
