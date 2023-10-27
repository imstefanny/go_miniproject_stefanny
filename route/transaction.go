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

func TransactionRoute(e *echo.Echo, db *gorm.DB) {
	transactionRepository := repository.NewTransactionRepository(db)

	seatRepository := repository.NewSeatRepository(db)

	ticketRepository := repository.NewTicketRepository(db)
	ticketUsecase := usecase.NewTicketUsecase(ticketRepository)

	showRepository := repository.NewShowRepository(db)
	showUsecase := usecase.NewShowUsecase(showRepository, seatRepository, ticketUsecase)

	transactionService := usecase.NewTransactionUsecase(transactionRepository, ticketUsecase, showUsecase)

	transactionController := controller.NewTransactionController(transactionService)

	eTransaction := e.Group("/transactions")
	eTransaction.GET("", transactionController.GetAll)
	eTransaction.POST("", transactionController.Create)
	eTransaction.GET("/:invoice", transactionController.GetByInvoice)
}