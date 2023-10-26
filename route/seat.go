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

	seatService := usecase.NewSeatUsecase(seatRepository)

	seatController := controller.NewSeatController(seatService)

	eSeat := e.Group("/seats")
	eSeat.GET("", seatController.GetAll)
}