package users

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

var users = []User{
	{ID: "1", FirstName: "Ada", LastName: "Lovelace", DateOfBirth: "10/12/1815"},
	{ID: "2", FirstName: "Lightning", LastName: "Quirks", DateOfBirth: "02/06/1955"},
}

func getUsers(ctx echo.Context) error {
	return ctx.JSON(http.StatusFound, users)
}

func postUser(ctx echo.Context) error {
	newUser := new(User)

	if err := ctx.Bind(newUser); err != nil {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	// Load into separate struct for security
	user := User{
		FirstName:   newUser.FirstName,
		LastName:    newUser.LastName,
		DateOfBirth: newUser.DateOfBirth,
	}

	fmt.Println(user)

	return ctx.JSON(http.StatusCreated, user)
}
