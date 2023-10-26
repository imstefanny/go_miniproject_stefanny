package controller

import (
	"net/http"
	
	"miniproject/usecase"

	"github.com/labstack/echo/v4"
)

type TicketController interface{
}

type ticketController struct {
	useCase usecase.TicketUsecase
}

func NewTicketController(ticketUsecase usecase.TicketUsecase) *ticketController {
	return &ticketController{ticketUsecase}
}

func (u *ticketController) GetAll(c echo.Context) error {
	tickets, err := u.useCase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": tickets,
	})
}