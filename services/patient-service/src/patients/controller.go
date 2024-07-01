package patients

import "github.com/kataras/iris/v12"

func PatientController(app *iris.Application) {
	service := PatientsService{}

	patientsAPI := app.Party("/patients")
	{
		patientsAPI.Use(iris.Compression)
		patientsAPI.Get("/", service.getPatients)
	}
}
