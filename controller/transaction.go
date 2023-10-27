package controller

import (
	"net/http"

	"miniproject/dto"
	"miniproject/usecase"

	"github.com/labstack/echo/v4"
)

type TransactionController interface{
}

type transactionController struct {
	useCase usecase.TransactionUsecase
}

func NewTransactionController(transactionUsecase usecase.TransactionUsecase) *transactionController {
	return &transactionController{transactionUsecase}
}


func (u *transactionController) Create(c echo.Context) error {
	transaction := dto.CreateTransactionRequest{}

	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}

	inv, err := u.useCase.Create(transaction)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"ticket_code": inv,
		"data": transaction,
	})
}

func (u *transactionController) GetAll(c echo.Context) error {
	transactions, err := u.useCase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": transactions,
	})
}

func (u *transactionController) GetByInvoice(c echo.Context) error {
	invoice := c.Param("invoice")
	transaction, err := u.useCase.GetByInvoice(invoice)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": transaction,
	})
}

