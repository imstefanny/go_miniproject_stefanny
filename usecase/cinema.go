package usecase

import (
	"miniproject/dto"
	"miniproject/model"
	"miniproject/repository"
	// "miniproject/middlewares"
	"errors"
)

type CinemaUsecase interface {
	Create(cinema dto.CreateCinemaRequest) error
	GetAll() (interface{}, error)
	Find(id int) (interface{}, error)
	Delete(id int) error
	Update(id int, cinema dto.CreateCinemaRequest) (model.Cinema, error)
}

type cinemaUsecase struct {
	cinemaRepository		repository.CinemaRepository
}

func NewCinemaUsecase(cinemaRepo repository.CinemaRepository) *cinemaUsecase {
	return &cinemaUsecase{cinemaRepository: cinemaRepo}
}

func (s *cinemaUsecase) Create(cinema dto.CreateCinemaRequest) error {
	if cinema.Name == "" || cinema.Location == "" || cinema.Street == "" || cinema.City == "" || cinema.Contact == "" {
		return errors.New("Field can't be empty")
	}

	cinemaData := model.Cinema{
		Name: cinema.Name,
		Location: cinema.Location,
		Street: cinema.Street,
		City: cinema.City,
		Contact: cinema.Contact,
	}
	err := s.cinemaRepository.Create(cinemaData)

	if err != nil {
		return err
	}

	return nil
}

func (s *cinemaUsecase) GetAll() (interface{}, error) {
	cinemas, err := s.cinemaRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return cinemas, nil
}

func (s *cinemaUsecase) Find(id int) (interface{}, error) {
	cinema, err := s.cinemaRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return cinema, nil
}

func (s *cinemaUsecase) Delete(id int) error {
	err := s.cinemaRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *cinemaUsecase) Update(id int, cinema dto.CreateCinemaRequest) (model.Cinema, error) {
	cinemaData, err := s.cinemaRepository.Find(id)

	if err != nil {
		return model.Cinema{}, err
	}

	cinemaData.Name = cinema.Name
	cinemaData.Location = cinema.Location
	cinemaData.Street = cinema.Street
	cinemaData.City = cinema.City
	cinemaData.Contact = cinema.Contact
	
	e := s.cinemaRepository.Update(id, cinemaData)

	if e != nil {
		return model.Cinema{}, e
	}

	cinemaUpdated, err := s.cinemaRepository.Find(id)

	if err != nil {
		return model.Cinema{}, err
	}
	return cinemaUpdated, nil
}
