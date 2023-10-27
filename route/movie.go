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

func MovieRoute(e *echo.Echo, db *gorm.DB) {
	movieRepository := repository.NewMovieRepository(db)

	movieService := usecase.NewMovieUsecase(movieRepository)

	movieController := controller.NewMovieController(movieService)

	eMovie := e.Group("/movies")
	eMovie.GET("", movieController.GetAll)
	eMovie.GET("/recommend", movieController.GetMovieRecommendations)
	eMovie.GET("/:id", movieController.Find)
	eMovie.POST("", movieController.Create)
	eMovie.PUT("/:id", movieController.Update)
	eMovie.DELETE("/:id", movieController.Delete)
}