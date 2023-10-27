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
	route.UserRoute(e, db)
	route.AccountRoute(e, db)
	route.ShowRoute(e, db)
	route.StudioRoute(e, db)
	route.SeatRoute(e, db)
	route.TicketRoute(e, db)
	route.TransactionRoute(e, db)
	
	e.Logger.Fatal(e.Start(":8000"))
}