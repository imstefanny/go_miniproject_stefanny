package usecase

import (
	"miniproject/dto"
	"miniproject/model"
	"miniproject/repository"
)

type SeatUsecase interface {
	Create(seat dto.CreateSeatRequest) error
	GetAll() (interface{}, error)
	Find(id int) (interface{}, error)
	GetAvailableSeats(showID int) ([]model.Seat, error)
}

type seatUsecase struct {
	seatRepository		repository.SeatRepository
	ticketRepository  repository.TicketRepository
}

func NewSeatUsecase(seatRepo repository.SeatRepository, ticketRepo repository.TicketRepository) *seatUsecase {
	return &seatUsecase{
		seatRepository: seatRepo,
		ticketRepository: ticketRepo,
	}
}

func (s *seatUsecase) Create(seat dto.CreateSeatRequest) error {
	seatData := model.Seat{
		StudioID: seat.StudioID,
		SeatNo: seat.SeatNo,
	}
	err := s.seatRepository.Create(seatData)

	if err != nil {
		return err
	}

	return nil
}

func (s *seatUsecase) GetAll() (interface{}, error) {
	seats, err := s.seatRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return seats, nil
}

func (s *seatUsecase) Find(id int) (interface{}, error) {
	seat, err := s.seatRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return seat, nil
}

func (s *seatUsecase) GetAvailableSeats(showID int) ([]model.Seat, error) {
	tickets, err := s.ticketRepository.GetAvailableTickets(showID)

	if err != nil {
		return nil, err
	}

	var available_seat []model.Seat
	for _, ticket := range tickets {
		seat, e := s.seatRepository.GetAvailableSeats(int(ticket.ID))
		if e != nil {
			return nil, e
		}
		available_seat = append(available_seat, seat)
	}

	return available_seat, nil
}
