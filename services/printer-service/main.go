package main

import "github.com/labstack/echo"

func main() {
	app := echo.New()

	app.Start(":8082")
}
