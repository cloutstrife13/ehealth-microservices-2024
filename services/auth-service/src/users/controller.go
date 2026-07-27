package users

import (
	"github.com/labstack/echo"
)

func UserController(app *echo.Echo) {
	usersAPI := app.Group("/users")
	usersAPI.GET("", getUsers)
	usersAPI.POST("", postUser)
}
