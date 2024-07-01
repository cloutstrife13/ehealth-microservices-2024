package patients

import (
	"net/http"

	"github.com/labstack/echo"
)

type PatientController struct {
	Service *PatientService
}

func (controller PatientController) RegisterEndpoints(app *echo.Echo) {
	patientEndpoint := app.Group("/patients")

	patientEndpoint.GET("", controller.FindPatients)
}

func (controller PatientController) FindPatients(ctx echo.Context) error {
	patients := []Patient{}
	controller.Service.FindPatients(&patients)
	return ctx.JSON(http.StatusFound, patients)
}
