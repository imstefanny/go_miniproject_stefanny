package usecase

import (
	"miniproject/dto"
	"miniproject/model"
	"miniproject/repository"
	"testing"
)

func TestCreateTicket(t *testing.T) {
	data := dto.CreateTicketRequest{
		ShowID: 1,
		SeatID: 1,
	}

	ticketData := model.Ticket{
		ShowID: 1,
		SeatID: 1,
	}
	
	mockTicketRepository := repository.NewMockTicketRepository()
	mockTicketRepository.On("Create", ticketData).Return(nil)

	service := NewTicketUsecase(mockTicketRepository)

	if err := service.Create(data); err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetAllTicket(t *testing.T) {
	mockTicket := []model.Ticket{
		{
			ID:	1,
			ShowID: 1,
			SeatID: 1,
		},
	}

	mockTicketRepository := repository.NewMockTicketRepository()
	mockTicketRepository.On("GetAll").Return(mockTicket, nil)

	service := NewTicketUsecase(mockTicketRepository)

	_, err := service.GetAll()
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetTicket(t *testing.T) {
	dataTicket := dto.CreateTicketRequest{
		ShowID: 1,
		SeatID: 1,
	}

	mockTicket := model.Ticket{
		ID:	1,
			ShowID: 1,
			SeatID: 1,
	}

	mockTicketRepository := repository.NewMockTicketRepository()
	mockTicketRepository.On("GetTicket").Return(mockTicket, nil)

	service := NewTicketUsecase(mockTicketRepository)

	_, err := service.GetTicket(dataTicket)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestCheckSoldOutTicket(t *testing.T) {
	mockTicket := model.Ticket{
		SeatID: 1,
		ShowID: 1,
		TransactionID: 1,
	}

	mockTicketRepository := repository.NewMockTicketRepository()
	mockTicketRepository.On("CheckSoldOut").Return(true, nil)
	mockTicketRepository.On("Find").Return(mockTicket, nil)

	service := NewTicketUsecase(mockTicketRepository)

	_, err := service.CheckSoldOut(uint(1))
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestCheckUnsoldOutTicket(t *testing.T) {
	mockTicket := model.Ticket{
		SeatID: 1,
		ShowID: 1,
	}

	mockTicketRepository := repository.NewMockTicketRepository()
	mockTicketRepository.On("CheckSoldOut").Return(false, nil)
	mockTicketRepository.On("Find").Return(mockTicket, nil)

	service := NewTicketUsecase(mockTicketRepository)

	_, err := service.CheckSoldOut(uint(1))
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}
