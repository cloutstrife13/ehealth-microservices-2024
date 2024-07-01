package patients

import (
	"github.com/kataras/iris/v12"
)

type PatientController struct{}

func (controller PatientController) RegisterEndpoints(app *iris.Application, service *PatientService) {
	patientEndpoint := app.Party("/patients")

	patientEndpoint.Use(iris.Compression)
	patientEndpoint.Get("/", service.getPatients)
}
