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

func AccountRoute(e *echo.Echo, db *gorm.DB) {
	accountRepository := repository.NewAccountRepository(db)

	accountService := usecase.NewAccountUsecase(accountRepository)

	accountController := controller.NewAccountController(accountService)
	
	e.POST("/login", accountController.Login)
	e.POST("/register", accountController.Create)

	eAccount := e.Group("/accounts")
	eAccount.GET("", accountController.GetAll)
	eAccount.GET("/:id", accountController.Find)
	eAccount.POST("", accountController.Create)
	eAccount.PUT("/:id", accountController.Update)
	eAccount.DELETE("/:id", accountController.Delete)
}