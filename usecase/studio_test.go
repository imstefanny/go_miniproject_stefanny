package usecase

import (
	"miniproject/dto"
	"miniproject/model"
	"miniproject/repository"
	"testing"
)

func TestCreateStudio(t *testing.T) {
	data := dto.CreateStudioRequest{
		Name: "studio 1",
    Capacity: 5,
    CinemaID: 1,
	}

	studioData := model.Studio{
		Name: "studio 1",
    Capacity: 5,
    CinemaID: 1,
	}
	
	mockStudioRepository := repository.NewMockStudioRepository()
	mockStudioRepository.On("Create", studioData).Return(uint(1), nil)

	mockTicketRepository := repository.NewMockTicketRepository()
	
	mockSeatRepository := repository.NewMockSeatRepository()
	seatData1 := model.Seat{
		StudioID: 1,
		SeatNo: "1",
	}
	mockSeatRepository.On("Create", seatData1).Return(nil)
	seatData2 := model.Seat{
		StudioID: 1,
		SeatNo: "2",
	}
	mockSeatRepository.On("Create", seatData2).Return(nil)
	seatData3 := model.Seat{
		StudioID: 1,
		SeatNo: "3",
	}
	mockSeatRepository.On("Create", seatData3).Return(nil)
	seatData4 := model.Seat{
		StudioID: 1,
		SeatNo: "4",
	}
	mockSeatRepository.On("Create", seatData4).Return(nil)
	seatData5 := model.Seat{
		StudioID: 1,
		SeatNo: "5",
	}
	mockSeatRepository.On("Create", seatData5).Return(nil)

	seatService := NewSeatUsecase(mockSeatRepository, mockTicketRepository)

	service := NewStudioUsecase(mockStudioRepository, seatService)

	if err := service.Create(data); err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestCreateStudioWithEmptyValue(t *testing.T) {
	data := dto.CreateStudioRequest{}

	studioData := model.Studio{}
	
	mockStudioRepository := repository.NewMockStudioRepository()
	mockStudioRepository.On("Create", studioData).Return(nil)

	mockTicketRepository := repository.NewMockTicketRepository()
	
	mockSeatRepository := repository.NewMockSeatRepository()
	seatService := NewSeatUsecase(mockSeatRepository, mockTicketRepository)

	service := NewStudioUsecase(mockStudioRepository, seatService)

	if err := service.Create(data); err.Error() != "Field can't be empty" {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetAllStudio(t *testing.T) {
	mockStudio := []model.Studio{
		{
			ID:	1,
			Name: "studio 1",
			Capacity: 5,
			CinemaID: 1,
		},
	}

	mockStudioRepository := repository.NewMockStudioRepository()
	mockStudioRepository.On("GetAll").Return(mockStudio, nil)

	mockTicketRepository := repository.NewMockTicketRepository()
	
	mockSeatRepository := repository.NewMockSeatRepository()
	seatService := NewSeatUsecase(mockSeatRepository, mockTicketRepository)

	service := NewStudioUsecase(mockStudioRepository, seatService)
	_, err := service.GetAll()
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetStudio(t *testing.T) {
	mockStudio := model.Studio{
		ID:	1,
		Name: "studio 1",
    Capacity: 5,
    CinemaID: 1,
	}

	mockStudioRepository := repository.NewMockStudioRepository()
	mockStudioRepository.On("Find").Return(mockStudio, nil)

	mockTicketRepository := repository.NewMockTicketRepository()
	
	mockSeatRepository := repository.NewMockSeatRepository()
	seatService := NewSeatUsecase(mockSeatRepository, mockTicketRepository)

	service := NewStudioUsecase(mockStudioRepository, seatService)

	_, err := service.Find(1)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestUpdateStudio(t *testing.T) {
	data := dto.CreateStudioRequest{
		Name: "studio 1",
    Capacity: 5,
    CinemaID: 1,
	}

	mockStudio := model.Studio{
		Name: "studio 1",
    Capacity: 5,
    CinemaID: 1,
	}

	mockStudioRepository := repository.NewMockStudioRepository()
	mockStudioRepository.On("Update", 1, mockStudio).Return(nil)
	mockStudioRepository.On("Find").Return(mockStudio, nil)

	mockTicketRepository := repository.NewMockTicketRepository()
	
	mockSeatRepository := repository.NewMockSeatRepository()
	seatService := NewSeatUsecase(mockSeatRepository, mockTicketRepository)

	service := NewStudioUsecase(mockStudioRepository, seatService)

	_, err := service.Update(1, data)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestDeleteStudio(t *testing.T) {
	mockStudioRepository := repository.NewMockStudioRepository()
	mockStudioRepository.On("Delete", 1).Return(nil)

	mockTicketRepository := repository.NewMockTicketRepository()
	
	mockSeatRepository := repository.NewMockSeatRepository()
	seatService := NewSeatUsecase(mockSeatRepository, mockTicketRepository)

	service := NewStudioUsecase(mockStudioRepository, seatService)

	err := service.Delete(1)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}
