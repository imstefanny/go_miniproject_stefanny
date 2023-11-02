package controller

import (
	"net/http"
	"strconv"

	"miniproject/dto"
	"miniproject/usecase"

	"github.com/labstack/echo/v4"
)

type UserController interface{
}

type userController struct {
	useCase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *userController {
	return &userController{userUsecase}
}

func (u *userController) Create(c echo.Context) error {
	user := dto.CreateUserRequest{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}

	err := u.useCase.Create(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": user,
	})
}

func (u *userController) GetAll(c echo.Context) error {
	users, err := u.useCase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": users,
	})
}

func (u *userController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": user,
	})
}

func (u *userController) Delete(c echo.Context) error {
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

func (u *userController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := dto.CreateUserRequest{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}

	userUpdated, err := u.useCase.Update(id, user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": userUpdated,
		"message": "success update data",
	})
}
