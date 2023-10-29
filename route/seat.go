package route

import (
	"miniproject/controller"
	"miniproject/repository"
	"miniproject/usecase"
	// "miniproject/constants"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func SeatRoute(e *echo.Echo, db *gorm.DB) {
	seatRepository := repository.NewSeatRepository(db)
	ticketRepository := repository.NewTicketRepository(db)

	seatService := usecase.NewSeatUsecase(seatRepository, ticketRepository)

	seatController := controller.NewSeatController(seatService)

	eSeat := e.Group("/seats")
	eSeat.GET("", seatController.GetAll)
	eSeat.GET("/:show_id", seatController.GetAvailableSeats)
}
