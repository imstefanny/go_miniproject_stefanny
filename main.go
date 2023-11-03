package main

import (
	"miniproject/database"
	"miniproject/route"

	"github.com/labstack/echo/v4"
)

func main() {
	db := database.InitDB()

	e := echo.New()

	route.Route(e, db)
	
	e.Logger.Fatal(e.Start(":8080"))
}
