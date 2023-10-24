package main

import (
	"miniproject/config"
	"miniproject/route"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()

	e := echo.New()

	route.CinemaRoute(e, db)
	route.MovieRoute(e, db)
	
	e.Logger.Fatal(e.Start(":8000"))
}