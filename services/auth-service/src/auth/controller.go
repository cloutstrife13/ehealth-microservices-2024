package auth

import (
	"net/http"

	"github.com/cloutstrife13/ehealth-microservices-2024/packages/models"
	"github.com/labstack/echo"
)

type RegisterRequest struct {
	models.User
	Password string `json:"password"`
}

type AuthController struct {
	Service *AuthService
}

func (controller AuthController) RegisterEndpoints(app *echo.Echo) {
	authEndpoint := app.Group("/auth")
	authEndpoint.POST("", controller.Register)
}

func (controller *AuthController) Register(ctx echo.Context) error {
	var req RegisterRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload: " + err.Error()})
	}

	if req.Password == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Password is required"})
	}

	if err := controller.Service.Register(&req.User, req.Password); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, req.User)
}
