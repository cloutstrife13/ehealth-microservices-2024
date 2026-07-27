package auth

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type AuthModule struct {
	App *echo.Echo
	Db  *gorm.DB
}

func (module AuthModule) RegisterModule() {
	service := AuthService{Db: module.Db}
	controller := AuthController{Service: &service}

	controller.RegisterEndpoints(module.App)
}
