package controller

import (
	"net/http"
	"strconv"

	"miniproject/dto"
	"miniproject/usecase"

	"github.com/labstack/echo/v4"
)

type CinemaController interface{
}

type cinemaController struct {
	useCase usecase.CinemaUsecase
}

func NewCinemaController(cinemaUsecase usecase.CinemaUsecase) *cinemaController {
	return &cinemaController{cinemaUsecase}
}

func (u *cinemaController) Create(c echo.Context) error {
	cinema := dto.CreateCinemaRequest{}

	if err := c.Bind(&cinema); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}

	err := u.useCase.Create(cinema)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": cinema,
	})
}

func (u *cinemaController) GetAll(c echo.Context) error {
	cinemas, err := u.useCase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": cinemas,
	})
}

func (u *cinemaController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cinema, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": cinema,
	})
}

func (u *cinemaController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := u.useCase.Delete(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted succesfully",
	})
}

func (u *cinemaController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cinema := dto.CreateCinemaRequest{}

	if err := c.Bind(&cinema); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}

	cinemaUpdated, err := u.useCase.Update(id, cinema)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": cinemaUpdated,
		"message": "success update data",
	})
}
