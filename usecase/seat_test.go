package usecase

import (
	"miniproject/dto"
	"miniproject/model"
	"miniproject/repository"
	"testing"
)

func TestCreateSeat(t *testing.T) {
	data := dto.CreateSeatRequest{
		SeatNo: "1",
		StudioID: 1,
	}

	seatData := model.Seat{
		SeatNo: "1",
		StudioID: 1,
	}
	
	mockTicketRepository := repository.NewMockTicketRepository()

	mockSeatRepository := repository.NewMockSeatRepository()
	mockSeatRepository.On("Create", seatData).Return(nil)

	service := NewSeatUsecase(mockSeatRepository, mockTicketRepository)

	if err := service.Create(data); err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetAllSeat(t *testing.T) {
	mockSeat := []model.Seat{
		{
			ID:	1,
			SeatNo: "1",
			StudioID: 1,
		},
	}

	mockTicketRepository := repository.NewMockTicketRepository()

	mockSeatRepository := repository.NewMockSeatRepository()
	mockSeatRepository.On("GetAll").Return(mockSeat, nil)

	service := NewSeatUsecase(mockSeatRepository, mockTicketRepository)

	_, err := service.GetAll()
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetSeat(t *testing.T) {
	mockSeat := []model.Seat{
		{
			ID:	1,
			SeatNo: "1",
			StudioID: 1,
		},
	}

	mockTicketRepository := repository.NewMockTicketRepository()

	mockSeatRepository := repository.NewMockSeatRepository()
	mockSeatRepository.On("Find").Return(mockSeat, nil)

	service := NewSeatUsecase(mockSeatRepository, mockTicketRepository)

	_, err := service.Find(1)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetAvailableSeats(t *testing.T) {
	mockSeat := []model.Seat{
		{
			ID:	1,
			SeatNo: "1",
			StudioID: 1,
		},
	}

	mockTicket := []model.Ticket{}

	mockTicketRepository := repository.NewMockTicketRepository()
	mockTicketRepository.On("GetAvailableTickets").Return(mockTicket, nil)

	mockSeatRepository := repository.NewMockSeatRepository()
	mockSeatRepository.On("GetAvailableSeats", 1).Return(mockSeat, nil)

	service := NewSeatUsecase(mockSeatRepository, mockTicketRepository)

	_, err := service.GetAvailableSeats(1)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}
