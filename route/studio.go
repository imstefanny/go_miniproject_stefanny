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

func StudioRoute(e *echo.Echo, db *gorm.DB) {
	studioRepository := repository.NewStudioRepository(db)
	seatRepository := repository.NewSeatRepository(db)
	ticketRepository := repository.NewTicketRepository(db)

	seatUsecase := usecase.NewSeatUsecase(seatRepository, ticketRepository)

	studioService := usecase.NewStudioUsecase(studioRepository, seatUsecase)

	studioController := controller.NewStudioController(studioService)

	eStudio := e.Group("/studios")
	eStudio.GET("", studioController.GetAll)
	eStudio.GET("/:id", studioController.Find)
	eStudio.POST("", studioController.Create)
	eStudio.PUT("/:id", studioController.Update)
	eStudio.DELETE("/:id", studioController.Delete)
}
