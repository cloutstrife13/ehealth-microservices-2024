package main

import (
	"github.com/kataras/iris/v12"
)

type User struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
	// InsuranceCode string `json:"insuranceCode"`
}

var users = []User{
	{ID: "1", FirstName: "Ada", LastName: "Lovelace", DateOfBirth: "10/12/1815"},
	{ID: "2", FirstName: "Lightning", LastName: "Quirks", DateOfBirth: "02/06/1955"},
}

func getUsers(ctx iris.Context) {
	ctx.JSON(users)
}

func main() {
	app := iris.New()

	usersAPI := app.Party("/users")
	{
		usersAPI.Use(iris.Compression)

		// GET: http://localhost:8080/books
		usersAPI.Get("/", getUsers)
	}

	app.Listen(":8181")
}
