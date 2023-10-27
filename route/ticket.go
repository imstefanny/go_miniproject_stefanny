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

func TicketRoute(e *echo.Echo, db *gorm.DB) {
	ticketRepository := repository.NewTicketRepository(db)

	ticketService := usecase.NewTicketUsecase(ticketRepository)

	ticketController := controller.NewTicketController(ticketService)

	eTicket := e.Group("/tickets")
	eTicket.GET("", ticketController.GetAll)
}