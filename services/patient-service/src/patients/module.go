package patients

import "github.com/kataras/iris/v12"

type PatientModule struct {
	App *iris.Application
}

func (module PatientModule) RegisterModule() {
	controller := PatientController{}
	service := PatientService{}

	controller.RegisterEndpoints(module.App, service)
}
