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
	service := PatientService{Db: module.Db}
	controller := PatientController{Service: &service}

	controller.RegisterEndpoints(module.App)
}
