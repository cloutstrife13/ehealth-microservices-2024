package patients

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type PatientModule struct {
	App *echo.Echo
	Db  *gorm.DB
}

func (module PatientModule) RegisterModule() {
	service := PatientService{Db: module.Db}
	controller := PatientController{Service: &service}

	controller.RegisterEndpoints(module.App)
}
