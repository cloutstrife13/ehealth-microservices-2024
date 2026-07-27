package users

import (
	"fmt"
	"net/http"

	"github.com/cloutstrife13/ehealth-microservices-2024/packages/helpers"
	"github.com/cloutstrife13/ehealth-microservices-2024/packages/models"
	"github.com/labstack/echo"
)

var users = []models.User{
	{ID: "1", Name: "Ada", Surname: "Lovelace", DateOfBirth: helpers.ParseDate("10/12/1815")},
	{ID: "2", Name: "Lightning", Surname: "Quirks", DateOfBirth: helpers.ParseDate("02/06/1955")},
}

func getUsers(ctx echo.Context) error {
	return ctx.JSON(http.StatusFound, users)
}

func postUser(ctx echo.Context) error {
	newUser := new(models.User)

	if err := ctx.Bind(newUser); err != nil {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	// Load into separate struct for security
	user := models.User{
		Name:        newUser.Name,
		Surname:     newUser.Surname,
		DateOfBirth: newUser.DateOfBirth,
	}

	fmt.Println(user)

	return ctx.JSON(http.StatusCreated, user)
}
