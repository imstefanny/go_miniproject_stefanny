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

func NewRoute(e *echo.Echo, db *gorm.DB) {
	cinemaRepository := repository.NewCinemaRepository(db)

	cinemaService := usecase.NewCinemaUsecase(cinemaRepository)

	cinemaController := controller.NewCinemaController(cinemaService)

	eCinema := e.Group("/cinemas")
	eCinema.GET("", cinemaController.GetAll)
	eCinema.GET("/:id", cinemaController.Find)
	eCinema.POST("", cinemaController.Create)
	eCinema.PUT("/:id", cinemaController.Update)
	eCinema.DELETE("/:id", cinemaController.Delete)
}