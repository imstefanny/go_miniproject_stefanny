package usecase

import (
	"miniproject/dto"
	"miniproject/model"
	"miniproject/repository"

	// "miniproject/middlewares"
)

type SeatUsecase interface {
	Create(seat dto.CreateSeatRequest) error
	GetAll() (interface{}, error)
	Find(id int) (interface{}, error)
}

type seatUsecase struct {
	seatRepository		repository.SeatRepository
}

func NewSeatUsecase(seatRepo repository.SeatRepository) *seatUsecase {
	return &seatUsecase{seatRepository: seatRepo}
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
