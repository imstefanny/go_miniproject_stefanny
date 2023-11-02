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
	
	// basic authentication route
	e.POST("/register", accountController.Create)
	e.POST("/login", accountController.Login)

	// account route
	eAccount := e.Group("/accounts")
	eAccount.Use(middleware.JWT([]byte(constants.SECRET_KEY)))
	eAccount.Use(m.IsAdmin)
	eAccount.GET("", accountController.GetAll)
	eAccount.GET("/:id", accountController.Find)
	eAccount.PUT("/:id", accountController.Update)
	eAccount.DELETE("/:id", accountController.Delete)

	// user route
	eUser := e.Group("/users")
	eUser.Use(middleware.JWT([]byte(constants.SECRET_KEY)))
	eUser.GET("", userController.GetAll)
	eUser.GET("/:id", userController.Find)
	eUser.POST("", userController.Create)
	eUser.PUT("/:id", userController.Update)
	eUser.DELETE("/:id", userController.Delete, m.IsAdmin)

	// cinema route
	eCinema := e.Group("/cinemas")
	eCinema.Use(middleware.JWT([]byte(constants.SECRET_KEY)))
	eCinema.GET("", cinemaController.GetAll)
	eCinema.GET("/:id", cinemaController.Find)
	eCinema.POST("", cinemaController.Create, m.IsAdmin)
	eCinema.PUT("/:id", cinemaController.Update, m.IsAdmin)
	eCinema.DELETE("/:id", cinemaController.Delete, m.IsAdmin)

	// movie route
	eMovie := e.Group("/movies")
	eMovie.Use(middleware.JWT([]byte(constants.SECRET_KEY)))
	eMovie.GET("", movieController.GetAll)
	eMovie.GET("/recommend", movieController.GetMovieRecommendations)
	eMovie.GET("/:id", movieController.Find)
	eMovie.POST("", movieController.Create, m.IsAdmin)
	eMovie.PUT("/:id", movieController.Update, m.IsAdmin)
	eMovie.DELETE("/:id", movieController.Delete, m.IsAdmin)

	// ticket route
	eTicket := e.Group("/tickets")
	eTicket.Use(middleware.JWT([]byte(constants.SECRET_KEY)))
	eTicket.GET("", ticketController.GetAll)

	// seat route
	eSeat := e.Group("/seats")
	eSeat.Use(middleware.JWT([]byte(constants.SECRET_KEY)))
	eSeat.GET("", seatController.GetAll)
	eSeat.GET("/:show_id", seatController.GetAvailableSeats)

	// show route
	eShow := e.Group("/shows")
	eShow.Use(middleware.JWT([]byte(constants.SECRET_KEY)))
	eShow.GET("", showController.GetAll)
	eShow.GET("/:id", showController.Find)
	eShow.POST("", showController.Create, m.IsAdmin)
	eShow.PUT("/:id", showController.Update, m.IsAdmin)
	eShow.DELETE("/:id", showController.Delete, m.IsAdmin)

	// studio route
	eStudio := e.Group("/studios")
	eStudio.Use(middleware.JWT([]byte(constants.SECRET_KEY)))
	eStudio.GET("", studioController.GetAll)
	eStudio.GET("/:id", studioController.Find)
	eStudio.POST("", studioController.Create, m.IsAdmin)
	eStudio.PUT("/:id", studioController.Update, m.IsAdmin)
	eStudio.DELETE("/:id", studioController.Delete, m.IsAdmin)

	// transaction route
	eTrans := e.Group("/transactions")
	eTrans.Use(middleware.JWT([]byte(constants.SECRET_KEY)))
	eTrans.GET("", transactionController.GetAll, m.IsAdmin)
	eTrans.GET("/:invoice", transactionController.GetByInvoice)
	eTrans.POST("", transactionController.Create)
}
