package main

import (
	"miniproject/config"
	"miniproject/route"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()

	e := echo.New()

	route.Route(e, db)
	
	e.Logger.Fatal(e.Start(":8080"))
}
