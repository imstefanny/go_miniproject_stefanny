package route

import (
	"miniproject/controller"
	"miniproject/repository"
	"miniproject/usecase"
	"miniproject/constants"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	m "miniproject/middlewares"
)

func Route(e *echo.Echo, db *gorm.DB) {
	accountRepository := repository.NewAccountRepository(db)
	accountService := usecase.NewAccountUsecase(accountRepository)
	accountController := controller.NewAccountController(accountService)

	userRepository := repository.NewUserRepository(db)
	userService := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userService)

	cinemaRepository := repository.NewCinemaRepository(db)
	cinemaService := usecase.NewCinemaUsecase(cinemaRepository)
	cinemaController := controller.NewCinemaController(cinemaService)
	
	movieRepository := repository.NewMovieRepository(db)
	movieService := usecase.NewMovieUsecase(movieRepository)
	movieController := controller.NewMovieController(movieService)

	ticketRepository := repository.NewTicketRepository(db)
	ticketService := usecase.NewTicketUsecase(ticketRepository)
	ticketController := controller.NewTicketController(ticketService)

	seatRepository := repository.NewSeatRepository(db)
	seatService := usecase.NewSeatUsecase(seatRepository, ticketRepository)
	seatController := controller.NewSeatController(seatService)

	showRepository := repository.NewShowRepository(db)
	showService := usecase.NewShowUsecase(showRepository, seatRepository, ticketService)
	showController := controller.NewShowController(showService)

	studioRepository := repository.NewStudioRepository(db)
	studioService := usecase.NewStudioUsecase(studioRepository, seatService)
	studioController := controller.NewStudioController(studioService)

	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := usecase.NewTransactionUsecase(transactionRepository, ticketService, showService)
	transactionController := controller.NewTransactionController(transactionService)

	e.Pre(middleware.RemoveTrailingSlash())
	
	eAdmin := e.Group("")
	eAdmin.Use(m.IsAdmin)

	eUser := e.Group("")
	eUser.Use(middleware.JWT([]byte(constants.SECRET_KEY)))

	// basic authentication route
	e.POST("/register", accountController.Create)
	e.POST("/login", accountController.Login)

	// account route
	eAdmin.GET("/accounts", accountController.GetAll)
	eAdmin.GET("/accounts/:id", accountController.Find)
	eAdmin.PUT("/accounts/:id", accountController.Update)
	eAdmin.DELETE("/accounts/:id", accountController.Delete)

	// user route
	eUser.GET("/users", userController.GetAll)
	eUser.GET("/users/:id", userController.Find)
	eUser.POST("/users", userController.Create)
	eUser.PUT("/users/:id", userController.Update)
	eAdmin.DELETE("/users/:id", userController.Delete)

	// cinema route
	eUser.GET("/cinemas", cinemaController.GetAll)
	eUser.GET("/cinemas/:id", cinemaController.Find)
	eAdmin.POST("/cinemas", cinemaController.Create)
	eAdmin.PUT("/cinemas/:id", cinemaController.Update)
	eAdmin.DELETE("/cinemas/:id", cinemaController.Delete)

	// movie route
	eUser.GET("/movies", movieController.GetAll)
	eUser.GET("/movies/recommend", movieController.GetMovieRecommendations)
	eUser.GET("/movies/:id", movieController.Find)
	eAdmin.POST("/movies", movieController.Create)
	eAdmin.PUT("/movies/:id", movieController.Update)
	eAdmin.DELETE("/movies/:id", movieController.Delete)

	// ticket route
	eUser.GET("/tickets", ticketController.GetAll)

	// seat route
	eUser.GET("/seats", seatController.GetAll)
	eUser.GET("/seats/:show_id", seatController.GetAvailableSeats)

	// show route
	eUser.GET("/shows", showController.GetAll)
	eUser.GET("/shows/:id", showController.Find)
	eAdmin.POST("/shows", showController.Create)
	eAdmin.PUT("/shows/:id", showController.Update)
	eAdmin.DELETE("/shows/:id", showController.Delete)

	// studio route
	eUser.GET("/studios", studioController.GetAll)
	eUser.GET("/studios/:id", studioController.Find)
	eAdmin.POST("/studios", studioController.Create)
	eAdmin.PUT("/studios/:id", studioController.Update)
	eAdmin.DELETE("/studios/:id", studioController.Delete)

	// transaction route
	eAdmin.GET("/transactions", transactionController.GetAll)
	eUser.GET("/transactions/:invoice", transactionController.GetByInvoice)
	eUser.POST("/transactions", transactionController.Create)
}
