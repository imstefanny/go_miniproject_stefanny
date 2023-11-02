package controller

import (
	"net/http"
	"strconv"

	"miniproject/usecase"

	"github.com/labstack/echo/v4"
)

type SeatController interface{
}

type seatController struct {
	useCase usecase.SeatUsecase
}

func NewSeatController(seatUsecase usecase.SeatUsecase) *seatController {
	return &seatController{seatUsecase}
}

func (u *seatController) GetAll(c echo.Context) error {
	seats, err := u.useCase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": seats,
	})
}

func (u *seatController) GetAvailableSeats(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("show_id"))

	seats, err := u.useCase.GetAvailableSeats(id)	

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": seats,
	})
}
