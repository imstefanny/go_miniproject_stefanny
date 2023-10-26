package controller

import (
	"net/http"
	"strconv"

	"miniproject/dto"
	"miniproject/usecase"

	"github.com/labstack/echo/v4"
)

type StudioController interface{
}

type studioController struct {
	useCase usecase.StudioUsecase
}

func NewStudioController(studioUsecase usecase.StudioUsecase) *studioController {
	return &studioController{studioUsecase}
}

func (u *studioController) Create(c echo.Context) error {
	studio := dto.CreateStudioRequest{}

	if err := c.Bind(&studio); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}

	err := u.useCase.Create(studio)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": studio,
	})
}

func (u *studioController) GetAll(c echo.Context) error {
	studios, err := u.useCase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": studios,
	})
}

func (u *studioController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	studio, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": studio,
	})
}

func (u *studioController) Delete(c echo.Context) error {
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

func (u *studioController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	studio := dto.CreateStudioRequest{}

	if err := c.Bind(&studio); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}

	studioUpdated, err := u.useCase.Update(id, studio)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": studioUpdated,
		"message": "success update data",
	})
}