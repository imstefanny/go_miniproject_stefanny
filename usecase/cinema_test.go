package usecase

import (
	"miniproject/dto"
	"miniproject/model"
	"miniproject/repository"
	"testing"
)

func TestCreateCinema(t *testing.T) {
	data := dto.CreateCinemaRequest{
		Name: "Cinema XII",
    Location: "Thamrin Plaza",
    Street: "Jalan Thamrin No 123",
    City: "Medan",
    Contact: "123456",
	}

	cinemaData := model.Cinema{
		Name: "Cinema XII",
    Location: "Thamrin Plaza",
    Street: "Jalan Thamrin No 123",
    City: "Medan",
    Contact: "123456",
	}
	
	mockCinemaRepository := repository.NewMockCinemaRepository()
	mockCinemaRepository.On("Create", cinemaData).Return(nil)

	service := NewCinemaUsecase(mockCinemaRepository)

	if err := service.Create(data); err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestCreateCinemaWithEmptyValue(t *testing.T) {
	data := dto.CreateCinemaRequest{}

	cinemaData := model.Cinema{}
	
	mockCinemaRepository := repository.NewMockCinemaRepository()
	mockCinemaRepository.On("Create", cinemaData).Return(nil)

	service := NewCinemaUsecase(mockCinemaRepository)

	if err := service.Create(data); err.Error() != "Field can't be empty" {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetAllCinema(t *testing.T) {
	mockCinema := []model.Cinema{
		{
			ID:	1,
			Name: "Cinema XII",
			Location: "Thamrin Plaza",
			Street: "Jalan Thamrin No 123",
			City: "Medan",
			Contact: "123456",
		},
	}

	mockCinemaRepository := repository.NewMockCinemaRepository()
	mockCinemaRepository.On("GetAll").Return(mockCinema, nil)

	service := NewCinemaUsecase(mockCinemaRepository)

	_, err := service.GetAll()
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetCinema(t *testing.T) {
	mockCinema := model.Cinema{
		ID:	1,
		Name: "Cinema XII",
    Location: "Thamrin Plaza",
    Street: "Jalan Thamrin No 123",
    City: "Medan",
    Contact: "123456",
	}

	mockCinemaRepository := repository.NewMockCinemaRepository()
	mockCinemaRepository.On("Find").Return(mockCinema, nil)

	service := NewCinemaUsecase(mockCinemaRepository)

	_, err := service.Find(1)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestUpdateCinema(t *testing.T) {
	data := dto.CreateCinemaRequest{
		Name: "Cinema XII",
    Location: "Thamrin Plaza",
    Street: "Jalan Thamrin No 123",
    City: "Medan",
    Contact: "123456",
	}

	mockCinema := model.Cinema{
		Name: "Cinema XII",
    Location: "Thamrin Plaza",
    Street: "Jalan Thamrin No 123",
    City: "Medan",
    Contact: "123456",
	}

	mockCinemaRepository := repository.NewMockCinemaRepository()
	mockCinemaRepository.On("Update", 1, mockCinema).Return(nil)
	mockCinemaRepository.On("Find").Return(mockCinema, nil)

	service := NewCinemaUsecase(mockCinemaRepository)

	_, err := service.Update(1, data)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestDeleteCinema(t *testing.T) {
	mockCinemaRepository := repository.NewMockCinemaRepository()
	mockCinemaRepository.On("Delete", 1).Return(nil)

	service := NewCinemaUsecase(mockCinemaRepository)

	err := service.Delete(1)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}
