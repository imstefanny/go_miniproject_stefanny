package usecase

import (
	"miniproject/dto"
	"miniproject/model"
	"miniproject/repository"

	// "miniproject/middlewares"
)

type TicketUsecase interface {
	Create(ticket dto.CreateTicketRequest) error
	GetAll() (interface{}, error)
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

