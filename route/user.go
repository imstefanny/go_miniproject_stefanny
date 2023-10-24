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

func UserRoute(e *echo.Echo, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)

	userService := usecase.NewUserUsecase(userRepository)

	userController := controller.NewUserController(userService)

	eUser := e.Group("/users")
	eUser.GET("", userController.GetAll)
	eUser.GET("/:id", userController.Find)
	eUser.POST("", userController.Create)
	eUser.PUT("/:id", userController.Update)
	eUser.DELETE("/:id", userController.Delete)
}