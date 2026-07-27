package auth

import (
	"net/http"

	"github.com/cloutstrife13/ehealth-microservices-2024/packages/models"
	"github.com/labstack/echo"
)

type AuthController struct {
	Service *AuthService
}

func (controller AuthController) RegisterEndpoints(app *echo.Echo) {
	authAPI := app.Group("/auth")
	authAPI.POST("", controller.Register)
}

func (controller AuthController) Register(c echo.Context) error {
	newUser := new(models.User)
	if err := c.Bind(newUser); err != nil {
		return err
	}
	if err := controller.Service.Register(newUser); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, newUser)
}
