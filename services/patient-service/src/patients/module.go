package patients

import (
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type PatientModule struct {
	App *iris.Application
	Db  *gorm.DB
}

func (module PatientModule) RegisterModule() {
	controller := PatientController{}
	service := PatientService{Db: module.Db}

	controller.RegisterEndpoints(module.App, service)
}
