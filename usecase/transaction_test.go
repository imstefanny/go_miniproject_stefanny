package usecase

import (
	"miniproject/dto"
	"miniproject/model"
	"miniproject/repository"
	"testing"
	"time"
)

func TestCreateTransaction(t *testing.T) {
	transactionData := dto.CreateTransactionRequest{
		AccountID: 1,
		ShowID: 1,
		SeatID: []uint{1},
	}
	
	mockTransaction := model.Transaction{
    AccountID: 1,
		Tickets: []model.Ticket{
			{
				ID:	1,
				ShowID: 1,
				SeatID: 1,
			},
		},
		Date: time.Now(),
		TotalPrice: 65000,
		Status: "paid",
		TicketCode: generateInvoiceNumber(),
	}

	mockTicket := model.Ticket{
		ID:	1,
		ShowID: 1,
		SeatID: 1,
	}

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
	
	mockTransactionRepository := repository.NewMockTransactionRepository()
	mockTransactionRepository.On("Create", mockTransaction).Return(nil)

	mockTicketRepository := repository.NewMockTicketRepository()
	mockTicketRepository.On("GetTicket").Return(mockTicket, nil)
	mockTicketRepository.On("CheckSoldOut").Return(false, nil)
	mockTicketRepository.On("Find").Return(mockTicket, nil)

	mockShowRepository := repository.NewMockShowRepository()
	mockShowRepository.On("Find").Return(mockShow, nil)

	mockSeatRepository := repository.NewMockSeatRepository()

	ticketUsecase := NewTicketUsecase(mockTicketRepository)
	showUsecase := NewShowUsecase(mockShowRepository, mockSeatRepository, ticketUsecase)

	service := NewTransactionUsecase(mockTransactionRepository, ticketUsecase, showUsecase)

	invoice, err := service.Create(transactionData)
	
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
	if invoice != mockTransaction.TicketCode{
		t.Errorf("Invoice is not the same")
	}
}

func TestCreateTransactionWithEmptyValue(t *testing.T) {
	transactionData := dto.CreateTransactionRequest{}
	
	mockTransaction := model.Transaction{}

	mockTicket := model.Ticket{}

	mockShow := model.Show{}
	
	mockTransactionRepository := repository.NewMockTransactionRepository()
	mockTransactionRepository.On("Create", mockTransaction).Return(nil)

	mockTicketRepository := repository.NewMockTicketRepository()
	mockTicketRepository.On("GetTicket").Return(mockTicket, nil)
	mockTicketRepository.On("CheckSoldOut").Return(false, nil)
	mockTicketRepository.On("Find").Return(mockTicket, nil)

	mockShowRepository := repository.NewMockShowRepository()
	mockShowRepository.On("Find").Return(mockShow, nil)

	mockSeatRepository := repository.NewMockSeatRepository()

	ticketUsecase := NewTicketUsecase(mockTicketRepository)
	showUsecase := NewShowUsecase(mockShowRepository, mockSeatRepository, ticketUsecase)

	service := NewTransactionUsecase(mockTransactionRepository, ticketUsecase, showUsecase)

	if _, err := service.Create(transactionData); err.Error() != "Field can't be empty" {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetAllTransaction(t *testing.T) {
	mockTransaction := []model.Transaction{
		{
			AccountID: 1,
			Tickets: []model.Ticket{
				{
					ID:	1,
					ShowID: 1,
					SeatID: 1,
				},
			},
			Date: time.Now(),
			TotalPrice: 65000,
			Status: "paid",
			TicketCode: generateInvoiceNumber(),
		},
	}

	mockTransactionRepository := repository.NewMockTransactionRepository()
	mockTransactionRepository.On("GetAll").Return(mockTransaction, nil)

	mockTicketRepository := repository.NewMockTicketRepository()

	mockShowRepository := repository.NewMockShowRepository()

	mockSeatRepository := repository.NewMockSeatRepository()

	ticketUsecase := NewTicketUsecase(mockTicketRepository)
	showUsecase := NewShowUsecase(mockShowRepository, mockSeatRepository, ticketUsecase)

	service := NewTransactionUsecase(mockTransactionRepository, ticketUsecase, showUsecase)

	_, err := service.GetAll()
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}

func TestGetTransaction(t *testing.T) {
	mockTransaction := model.Transaction{
		AccountID: 1,
		Tickets: []model.Ticket{
			{
				ID:	1,
				ShowID: 1,
				SeatID: 1,
			},
		},
		Date: time.Now(),
		TotalPrice: 65000,
		Status: "paid",
		TicketCode: "1698771538-4788",
	}

	mockTransactionRepository := repository.NewMockTransactionRepository()
	mockTransactionRepository.On("GetByInvoice").Return(mockTransaction, nil)

	mockTicketRepository := repository.NewMockTicketRepository()

	mockShowRepository := repository.NewMockShowRepository()

	mockSeatRepository := repository.NewMockSeatRepository()

	ticketUsecase := NewTicketUsecase(mockTicketRepository)
	showUsecase := NewShowUsecase(mockShowRepository, mockSeatRepository, ticketUsecase)

	service := NewTransactionUsecase(mockTransactionRepository, ticketUsecase, showUsecase)

	_, err := service.GetByInvoice("1698771538-4788")
	if err != nil {
		t.Errorf("Got Error %v", err)
	}
}
