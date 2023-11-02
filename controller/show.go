package controller

import (
	"net/http"
	"strconv"

	"miniproject/dto"
	"miniproject/usecase"

	"github.com/labstack/echo/v4"
)

type ShowController interface{
}

type showController struct {
	useCase usecase.ShowUsecase
}

func NewShowController(showUsecase usecase.ShowUsecase) *showController {
	return &showController{showUsecase}
}

func (u *showController) Create(c echo.Context) error {
	show := dto.CreateShowRequest{}

	if err := c.Bind(&show); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}

	err := u.useCase.Create(show)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": show,
	})
}

func (u *showController) GetAll(c echo.Context) error {
	shows, err := u.useCase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": shows,
	})
}

func (u *showController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	show, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": show,
	})
}

func (u *showController) Delete(c echo.Context) error {
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

func (u *showController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	show := dto.CreateShowRequest{}

	if err := c.Bind(&show); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}

	showUpdated, err := u.useCase.Update(id, show)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": showUpdated,
		"message": "success update data",
	})
}
