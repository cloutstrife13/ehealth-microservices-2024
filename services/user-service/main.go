package main

import (
	"github.com/cloutstrife13/ehealth-microservices-2024/services/user-service/users"
	"github.com/labstack/echo"
)

func main() {
	app := echo.New()

	users.UserController(app)

	app.Start(":8081")
}
