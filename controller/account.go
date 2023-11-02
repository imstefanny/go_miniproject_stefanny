package controller

import (
	"net/http"
	"strconv"

	"miniproject/dto"
	"miniproject/usecase"

	"github.com/labstack/echo/v4"
)

type AccountController interface {
}

type accountController struct {
	useCase usecase.AccountUsecase
}

func NewAccountController(accountUsecase usecase.AccountUsecase) *accountController {
	return &accountController{accountUsecase}
}

func (u *accountController) Create(c echo.Context) error {
	account := dto.CreateAccountRequest{}

	if err := c.Bind(&account); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}

	err := u.useCase.Create(account)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": account,
	})
}

func (u *accountController) GetAll(c echo.Context) error {
	accounts, err := u.useCase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": accounts,
	})
}

func (u *accountController) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	account, err := u.useCase.Find(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": account,
	})
}

func (u *accountController) Delete(c echo.Context) error {
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

func (u *accountController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	account := dto.CreateAccountRequest{}

	if err := c.Bind(&account); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}

	accountUpdated, err := u.useCase.Update(id, account)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": accountUpdated,
		"message": "success update data",
	})
}

func (u *accountController) Login(c echo.Context) error {
	account := dto.LoginAccountRequest{}

	if err := c.Bind(&account); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}

	accountResponse, err := u.useCase.Login(account)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"account": accountResponse,
	})
}
