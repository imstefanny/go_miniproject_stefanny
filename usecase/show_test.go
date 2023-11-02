package usecase

import (
	"miniproject/dto"
	"miniproject/model"
	"miniproject/repository"
	"testing"
	"time"
)

func TestCreateShow(t *testing.T) {
	showDateStr := "2023-10-28T00:00:00Z"
	showStartStr := "2023-10-28T14:30:00Z"
	showEndStr := "2023-10-28T16:00:00Z"
	
	showDate, _ := time.Parse(time.RFC3339, showDateStr)
	showStart, _ := time.Parse(time.RFC3339, showStartStr)
	showEnd, _ := time.Parse(time.RFC3339, showEndStr)
	
	data := dto.CreateShowRequest{
		MovieID: 1,
		StudioID: 1,
		ShowDate: showDate,
		ShowStart: showStart,
		ShowEnd: showEnd,
		Price: 65000,
	}

	showData := model.Show{
		MovieID: 1,
		StudioID: 1,
		ShowDate: showDate,
		ShowStart: showStart,
		ShowEnd: showEnd,
		Price: 65000,
	}

	seatData := []model.Seat{
		{
			SeatNo: "1",
			StudioID: 1,
		},
		{
			SeatNo: "2",
			StudioID: 1,
		},
		{
			SeatNo: "3",
			StudioID: 1,
		},
		{
			SeatNo: "4",
			StudioID: 1,
		},
		{
			SeatNo: "5",
			StudioID: 1,
		},
	}
	
	mockShowRepository := repository.NewMockShowRepository()
	mockShowRepository.On("Create", showData).Return(uint(1), nil)

	mockSeatRepository := repository.NewMockSeatRepository()
	mockSeatRepository.On("Find").Return(seatData, nil)

	mockTicketRepository := repository.NewMockTicketRepository()
	ticketUsecase := NewTicketUsecase(mockTicketRepository)
	ticketData1 := model.Ticket{
		SeatID: 0,
		ShowID: 1,
	}
	mockTicketRepository.On("Create", ticketData1).Return(nil)
	ticketData2 := model.Ticket{
		SeatID: 2,
		ShowID: 1,
	}
	mockTicketRepository.On("Create", ticketData2).Return(nil)
	ticketData3 := model.Ticket{
		SeatID: 3,
		ShowID: 1,
	}
	mockTicketRepository.On("Create", ticketData3).Return(nil)
	ticketData4 := model.Ticket{
		SeatID: 4,
		ShowID: 1,
	}
	mockTicketRepository.On("Create", ticketData4).Return(nil)
	ticketData5 := model.Ticket{
		SeatID: 5,
		ShowID: 1,
	}
	mockTicketRepository.On("Create", ticketData5).Return(nil)

	service := NewShowUsecase(mockShowRepository, mockSeatRepository, ticketUsecase)

	if err := service.Create(data); err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestCreateShowWithEmptyValue(t *testing.T) {
	data := dto.CreateShowRequest{}

	showData := model.Show{}
	
	mockShowRepository := repository.NewMockShowRepository()
	mockShowRepository.On("Create", showData).Return(nil)

	mockSeatRepository := repository.NewMockSeatRepository()

	mockTicketRepository := repository.NewMockTicketRepository()
	ticketUsecase := NewTicketUsecase(mockTicketRepository)

	service := NewShowUsecase(mockShowRepository, mockSeatRepository, ticketUsecase)

	if err := service.Create(data); err.Error() != "Field can't be empty" {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetAllShow(t *testing.T) {
	showDateStr := "2023-10-28T00:00:00Z"
	showStartStr := "2023-10-28T14:30:00Z"
	showEndStr := "2023-10-28T16:00:00Z"
	
	showDate, _ := time.Parse(time.RFC3339, showDateStr)
	showStart, _ := time.Parse(time.RFC3339, showStartStr)
	showEnd, _ := time.Parse(time.RFC3339, showEndStr)
	
	mockShow := []model.Show{
		{
			ID:	1,
			MovieID: 1,
			StudioID: 1,
			ShowDate: showDate,
			ShowStart: showStart,
			ShowEnd: showEnd,
			Price: 65000,
		},
	}

	mockShowRepository := repository.NewMockShowRepository()
	mockShowRepository.On("GetAll").Return(mockShow, nil)

	mockSeatRepository := repository.NewMockSeatRepository()

	mockTicketRepository := repository.NewMockTicketRepository()
	ticketUsecase := NewTicketUsecase(mockTicketRepository)

	service := NewShowUsecase(mockShowRepository, mockSeatRepository, ticketUsecase)

	_, err := service.GetAll()
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetShow(t *testing.T) {
	showDateStr := "2023-10-28T00:00:00Z"
	showStartStr := "2023-10-28T14:30:00Z"
	showEndStr := "2023-10-28T16:00:00Z"
	
	showDate, _ := time.Parse(time.RFC3339, showDateStr)
	showStart, _ := time.Parse(time.RFC3339, showStartStr)
	showEnd, _ := time.Parse(time.RFC3339, showEndStr)

	mockShow := model.Show{
		ID:	1,
		MovieID: 1,
		StudioID: 1,
		ShowDate: showDate,
		ShowStart: showStart,
		ShowEnd: showEnd,
		Price: 65000,
	}

	mockShowRepository := repository.NewMockShowRepository()
	mockShowRepository.On("Find").Return(mockShow, nil)

	mockSeatRepository := repository.NewMockSeatRepository()

	mockTicketRepository := repository.NewMockTicketRepository()
	ticketUsecase := NewTicketUsecase(mockTicketRepository)

	service := NewShowUsecase(mockShowRepository, mockSeatRepository, ticketUsecase)

	_, err := service.Find(1)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestUpdateShow(t *testing.T) {
	showDateStr := "2023-10-28T00:00:00Z"
	showStartStr := "2023-10-28T14:30:00Z"
	showEndStr := "2023-10-28T16:00:00Z"
	
	showDate, _ := time.Parse(time.RFC3339, showDateStr)
	showStart, _ := time.Parse(time.RFC3339, showStartStr)
	showEnd, _ := time.Parse(time.RFC3339, showEndStr)

	data := dto.CreateShowRequest{
		MovieID: 1,
		StudioID: 1,
		ShowDate: showDate,
		ShowStart: showStart,
		ShowEnd: showEnd,
		Price: 65000,
	}

	mockShow := model.Show{
		MovieID: 1,
		StudioID: 1,
		ShowDate: showDate,
		ShowStart: showStart,
		ShowEnd: showEnd,
		Price: 65000,
	}

	mockShowRepository := repository.NewMockShowRepository()
	mockShowRepository.On("Update", 1, mockShow).Return(nil)
	mockShowRepository.On("Find").Return(mockShow, nil)

	mockSeatRepository := repository.NewMockSeatRepository()

	mockTicketRepository := repository.NewMockTicketRepository()
	ticketUsecase := NewTicketUsecase(mockTicketRepository)

	service := NewShowUsecase(mockShowRepository, mockSeatRepository, ticketUsecase)

	_, err := service.Update(1, data)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestDeleteShow(t *testing.T) {
	mockShowRepository := repository.NewMockShowRepository()
	mockShowRepository.On("Delete", 1).Return(nil)

	mockSeatRepository := repository.NewMockSeatRepository()

	mockTicketRepository := repository.NewMockTicketRepository()
	ticketUsecase := NewTicketUsecase(mockTicketRepository)

	service := NewShowUsecase(mockShowRepository, mockSeatRepository, ticketUsecase)

	err := service.Delete(1)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}
