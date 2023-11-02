package usecase

import (
	"miniproject/dto"
	"miniproject/model"
	"miniproject/repository"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	dateStr := "2003-12-21T00:00:00Z"
	date, _ := time.Parse(time.RFC3339, dateStr)
	data := dto.CreateUserRequest{
		FirstName: "Stefanny",
    LastName: "Stefanny",
    Email: "stefanny@mikroskil.ac.id",
    Phone: "08123456789",
    Address: "Jalan Sendirian No 123",
    Gender: "Perempuan",
    DateOfBirth: date,
    Pin: "123456",
    AccountID: 1,
	}

	userData := model.User{
		FirstName: "Stefanny",
    LastName: "Stefanny",
    Email: "stefanny@mikroskil.ac.id",
    Phone: "08123456789",
    Address: "Jalan Sendirian No 123",
    Gender: "Perempuan",
    DateOfBirth: date,
    Pin: "123456",
    AccountID: 1,
	}
	
	mockUserRepository := repository.NewMockUserRepository()
	mockUserRepository.On("Create", userData).Return(nil)

	service := NewUserUsecase(mockUserRepository)

	if err := service.Create(data); err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestCreateUserWithEmptyValue(t *testing.T) {
	data := dto.CreateUserRequest{}

	userData := model.User{}
	
	mockUserRepository := repository.NewMockUserRepository()
	mockUserRepository.On("Create", userData).Return(nil)

	service := NewUserUsecase(mockUserRepository)

	if err := service.Create(data); err.Error() != "Field can't be empty" {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetAllUser(t *testing.T) {
	dateStr := "2003-12-21T00:00:00Z"
	date, _ := time.Parse(time.RFC3339, dateStr)
	mockUser := []model.User{
		{
			ID:	1,
			FirstName: "Stefanny",
			LastName: "Stefanny",
			Email: "stefanny@mikroskil.ac.id",
			Phone: "08123456789",
			Address: "Jalan Sendirian No 123",
			Gender: "Perempuan",
			DateOfBirth: date,
			Pin: "123456",
			AccountID: 1,
		},
	}

	mockUserRepository := repository.NewMockUserRepository()
	mockUserRepository.On("GetAll").Return(mockUser, nil)

	service := NewUserUsecase(mockUserRepository)

	_, err := service.GetAll()
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetUser(t *testing.T) {
	dateStr := "2003-12-21T00:00:00Z"
	date, _ := time.Parse(time.RFC3339, dateStr)
	mockUser := model.User{
		ID:	1,
		FirstName: "Stefanny",
    LastName: "Stefanny",
    Email: "stefanny@mikroskil.ac.id",
    Phone: "08123456789",
    Address: "Jalan Sendirian No 123",
    Gender: "Perempuan",
    DateOfBirth: date,
    Pin: "123456",
    AccountID: 1,
	}

	mockUserRepository := repository.NewMockUserRepository()
	mockUserRepository.On("Find").Return(mockUser, nil)

	service := NewUserUsecase(mockUserRepository)

	_, err := service.Find(1)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestUpdateUser(t *testing.T) {
	dateStr := "2003-12-21T00:00:00Z"
	date, _ := time.Parse(time.RFC3339, dateStr)
	data := dto.CreateUserRequest{
		FirstName: "Stefanny",
    LastName: "Stefanny",
    Email: "stefanny@mikroskil.ac.id",
    Phone: "08123456789",
    Address: "Jalan Sendirian No 123",
    Gender: "Perempuan",
    DateOfBirth: date,
    Pin: "123456",
    AccountID: 1,
	}

	mockUser := model.User{
		FirstName: "Stefanny",
    LastName: "Stefanny",
    Email: "stefanny@mikroskil.ac.id",
    Phone: "08123456789",
    Address: "Jalan Sendirian No 123",
    Gender: "Perempuan",
    DateOfBirth: date,
    Pin: "123456",
    AccountID: 1,
	}

	mockUserRepository := repository.NewMockUserRepository()
	mockUserRepository.On("Update", 1, mockUser).Return(nil)
	mockUserRepository.On("Find").Return(mockUser, nil)

	service := NewUserUsecase(mockUserRepository)

	_, err := service.Update(1, data)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestDeleteUser(t *testing.T) {
	mockUserRepository := repository.NewMockUserRepository()
	mockUserRepository.On("Delete", 1).Return(nil)

	service := NewUserUsecase(mockUserRepository)

	err := service.Delete(1)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}
