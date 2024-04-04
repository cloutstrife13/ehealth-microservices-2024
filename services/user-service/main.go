package main

import (
	"github.com/cloutstrife13/ehealth-microservices-2024/services/user-service/users"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	users.UserController(app)

	app.Listen(":8081")
}
