package usecase

import (
	"miniproject/dto"
	"miniproject/model"
	"miniproject/repository"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	data := dto.CreateAccountRequest{
		Username: "Alta",
		Email: "Alta@gmail.com",
		Password: "123456",
		Role: "admin",
	}

	accountData := model.Account{
		Username: "Alta",
		Email: "Alta@gmail.com",
		Password: "123456",
		Role: "admin",
	}
	
	mockAccountRepository := repository.NewMockAccountRepository()
	mockAccountRepository.On("Create", accountData).Return(nil)

	service := NewAccountUsecase(mockAccountRepository)

	if err := service.Create(data); err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestCreateAccountWithEmptyValue(t *testing.T) {
	data := dto.CreateAccountRequest{}

	accountData := model.Account{}
	
	mockAccountRepository := repository.NewMockAccountRepository()
	mockAccountRepository.On("Create", accountData).Return(nil)

	service := NewAccountUsecase(mockAccountRepository)

	if err := service.Create(data); err.Error() != "Field can't be empty" {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetAllAccount(t *testing.T) {
	mockAccount := []model.Account{
		{
			ID:				1,
			Username: "Alta",
			Email:		"Alta@gmail.com",
			Password:	"123456",
			Role:			"admin",
		},
	}

	mockAccountRepository := repository.NewMockAccountRepository()
	mockAccountRepository.On("GetAll").Return(mockAccount, nil)

	service := NewAccountUsecase(mockAccountRepository)

	_, err := service.GetAll()
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetAccount(t *testing.T) {
	mockAccount := model.Account{
		ID:				1,
		Username: "Alta",
		Email:		"Alta@gmail.com",
		Password:	"123456",
		Role:			"admin",
	}

	mockAccountRepository := repository.NewMockAccountRepository()
	mockAccountRepository.On("Find").Return(mockAccount, nil)

	service := NewAccountUsecase(mockAccountRepository)

	_, err := service.Find(1)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestUpdateAccount(t *testing.T) {
	data := dto.CreateAccountRequest{
		Username: "AltaONE",
		Email: "Alta@gmail.com",
		Password: "123456",
		Role: "admin",
	}

	mockAccount := model.Account{
		Username: "AltaONE",
		Email:		"Alta@gmail.com",
		Password:	"123456",
		Role:			"admin",
	}

	mockAccountRepository := repository.NewMockAccountRepository()
	mockAccountRepository.On("Update", 1, mockAccount).Return(nil)
	mockAccountRepository.On("Find").Return(mockAccount, nil)

	service := NewAccountUsecase(mockAccountRepository)

	updated, err := service.Update(1, data)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
	if updated != mockAccount {
		t.Error("Not updated")
	}
}

func TestDeleteAccount(t *testing.T) {
	mockAccountRepository := repository.NewMockAccountRepository()
	mockAccountRepository.On("Delete", 1).Return(nil)

	service := NewAccountUsecase(mockAccountRepository)

	err := service.Delete(1)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestLogin(t *testing.T) {
	data := dto.LoginAccountRequest{
		Username: "Alta",
		Email:		"Alta@gmail.com",
		Password:	"123456",
	}
	
	userData := model.Account{
		Username: "Alta",
		Email:		"Alta@gmail.com",
		Password:	"123456",
	}

	mockAccount := model.Account{
		ID:				1,
		Username: "Alta",
		Email:		"Alta@gmail.com",
		Password:	"123456",
		Role:			"admin",
	}

	mockResponse := model.AccountResponse{
		ID:				1,
		Username: "Alta",
		Email:	  "Alta@gmail.com",
		Token:		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6IjEyMzQ1NiIsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoiQWx0YSJ9.zoL_No_MNpjVcH0aoVjAL9jkYLeqh4Sv1h_GTV55duI",
	}

	mockAccountRepository := repository.NewMockAccountRepository()
	mockAccountRepository.On("Login", userData).Return(mockAccount, nil)

	service := NewAccountUsecase(mockAccountRepository)

	res, err := service.Login(data)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
	if res != mockResponse {
		t.Errorf("Cannot login")
	}
}

func TestLoginWithEmptyValue(t *testing.T) {
	data := dto.LoginAccountRequest{}
	
	userData := model.Account{}

	mockAccount := model.Account{}

	mockAccountRepository := repository.NewMockAccountRepository()
	mockAccountRepository.On("Login", userData).Return(mockAccount, nil)

	service := NewAccountUsecase(mockAccountRepository)

	_, err := service.Login(data)
	if err.Error() != "Field can't be empty" {
		t.Errorf("Got Error %v", err)
	}
}
