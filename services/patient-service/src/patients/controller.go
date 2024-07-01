package patients

import (
	"github.com/kataras/iris/v12"
)

type PatientController struct {
	Service *PatientService
}

func (controller PatientController) RegisterEndpoints(app *iris.Application) {
	patientEndpoint := app.Party("/patients")

	patientEndpoint.Use(iris.Compression)
	patientEndpoint.Get("/", controller.FindPatients)
}

func (controller PatientController) FindPatients(ctx iris.Context) {
	patients := []Patient{}
	controller.Service.FindPatients(&patients)
	ctx.JSON(patients)
}
