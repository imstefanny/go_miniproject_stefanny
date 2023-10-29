package usecase

import (
	"miniproject/dto"
	"miniproject/model"
	"miniproject/repository"
)

type TicketUsecase interface {
	Create(ticket dto.CreateTicketRequest) error
	GetAll() (interface{}, error)
	GetTicket(ticket dto.CreateTicketRequest) (model.Ticket, error)
	CheckSoldOut(ticketID uint) (bool, error)
}

type ticketUsecase struct {
	ticketRepository		repository.TicketRepository
}

func NewTicketUsecase(ticketRepo repository.TicketRepository) *ticketUsecase {
	return &ticketUsecase{ticketRepository: ticketRepo}
}

func (s *ticketUsecase) Create(ticket dto.CreateTicketRequest) error {
	ticketData := model.Ticket{
		ShowID: ticket.ShowID,
		SeatID: ticket.SeatID,
	}
	err := s.ticketRepository.Create(ticketData)

	if err != nil {
		return err
	}

	return nil
}

func (s *ticketUsecase) GetAll() (interface{}, error) {
	tickets, err := s.ticketRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (s *ticketUsecase) GetTicket(ticket dto.CreateTicketRequest) (model.Ticket, error) {
	ticketID, err := s.ticketRepository.GetTicket(ticket.ShowID, ticket.SeatID)

	if err != nil {
		return model.Ticket{}, err
	}

	return ticketID, nil
}

func (s *ticketUsecase) CheckSoldOut(ticketID uint) (bool, error) {
	sold, err := s.ticketRepository.Find(int(ticketID))

	if err != nil {
		return false, err
	}

	var sold_out bool
	if sold.TransactionID != 0 {
		sold_out = true
	} else {
		sold_out = false
	}

	return sold_out, err
}
