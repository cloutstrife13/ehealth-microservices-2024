package users

import "github.com/kataras/iris/v12"

func UserController(app *iris.Application) {
	usersAPI := app.Party("/users")
	{
		usersAPI.Use(iris.Compression)
		usersAPI.Get("/", getUsers)
		usersAPI.Post("/", postUser)
	}
}
