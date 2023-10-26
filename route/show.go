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

func ShowRoute(e *echo.Echo, db *gorm.DB) {
	showRepository := repository.NewShowRepository(db)
	seatRepository := repository.NewSeatRepository(db)
	ticketRepository := repository.NewTicketRepository(db)

	ticketUsecase := usecase.NewTicketUsecase(ticketRepository)

	showService := usecase.NewShowUsecase(showRepository, seatRepository, ticketUsecase)

	showController := controller.NewShowController(showService)

	eShow := e.Group("/shows")
	eShow.GET("", showController.GetAll)
	eShow.GET("/:id", showController.Find)
	eShow.POST("", showController.Create)
	eShow.PUT("/:id", showController.Update)
	eShow.DELETE("/:id", showController.Delete)
}