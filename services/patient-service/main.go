package main

import (
	"cloutstrife13/patient-service/src/patients"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	patients.PatientController(app)

	app.Listen(":8080")
}
