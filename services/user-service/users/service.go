package users

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

var users = []User{
	{ID: "1", FirstName: "Ada", LastName: "Lovelace", DateOfBirth: "10/12/1815"},
	{ID: "2", FirstName: "Lightning", LastName: "Quirks", DateOfBirth: "02/06/1955"},
}

func getUsers(ctx iris.Context) {
	ctx.JSON(users)
}

func postUser(ctx iris.Context) {
	var user User

	err := ctx.ReadJSON(&user)

	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("User creation failure").DetailErr(err))
		return
	}

	fmt.Println(user)

	ctx.StatusCode(iris.StatusCreated)
}
