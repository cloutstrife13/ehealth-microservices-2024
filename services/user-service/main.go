package main

import (
	"cloutstrife13/user-service/src/users"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	users.UserController(app)

	app.Listen(":8081")
}
