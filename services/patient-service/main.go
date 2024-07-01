package main

import (
	. "github.com/cloutstrife13/ehealth-microservices-2024/services/patient-service/src/patients"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	PatientModule{App: app}.RegisterModule()

	app.Listen(":8080")
}
