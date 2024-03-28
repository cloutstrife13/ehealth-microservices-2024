package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID            string `json:"id"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	DateOfBirth   string `json:"dateOfBirth"`
	InsuranceCode string `json:"insuranceCode"`
}

var users = []user{
	{ID: "1", FirstName: "Ada", LastName: "Lovelace", DateOfBirth: "10/12/1815", InsuranceCode: "K435834234"},
	{ID: "2", FirstName: "Lightning", LastName: "Quirks", DateOfBirth: "02/06/1955", InsuranceCode: "K453656452"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.Run(":8181")
}
