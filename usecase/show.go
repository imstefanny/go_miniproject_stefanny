package usecase

import (
	"miniproject/dto"
	"miniproject/model"
	"miniproject/repository"
	"reflect"

	// "miniproject/middlewares"
	"errors"
)

type ShowUsecase interface {
	Create(show dto.CreateShowRequest) error
	GetAll() (interface{}, error)
	Find(id int) (model.Show, error)
	Delete(id int) error
	Update(id int, show dto.CreateShowRequest) (model.Show, error)
}

type showUsecase struct {
	showRepository		repository.ShowRepository
	seatRepository		repository.SeatRepository
	ticketUsecase			TicketUsecase
}

func NewShowUsecase(showRepo repository.ShowRepository, seatRepository repository.SeatRepository, ticketUsecase TicketUsecase) *showUsecase {
	return &showUsecase{
		showRepository: showRepo,
		seatRepository: seatRepository,
		ticketUsecase: ticketUsecase,
	}
}

func validateCreateShowRequest(req dto.CreateShowRequest) error {
	val := reflect.ValueOf(req)
	for i := 0; i < val.NumField(); i++ {
			if isEmptyField(val.Field(i)) {
					return errors.New("Field can't be empty")
			}
	}
	return nil
}

func (s *showUsecase) Create(show dto.CreateShowRequest) error {
	e := validateCreateShowRequest(show)
	
	if e!= nil {
		return e
	}

	showData := model.Show{
		MovieID: show.MovieID,
		StudioID: show.StudioID,
		ShowDate: show.ShowDate,
		ShowStart: show.ShowStart,
		ShowEnd: show.ShowEnd,
		Price: show.Price,
	}
	showID, err := s.showRepository.Create(showData)

	if err != nil {
		return err
	}

	// this is to retrieve seats associated with the specified studioID
	seats, err := s.seatRepository.Find(int(show.StudioID))
	if err != nil {
		return err
	}

	// this is for trigger the auto creation of ticket based on the show created
	for _, seat := range seats {
		ticket := dto.CreateTicketRequest{
			SeatID: seat.ID,
			ShowID: showID,
		}

		if err := s.ticketUsecase.Create(ticket); err != nil {
			return err
		}
	}

	return nil
}

func (s *showUsecase) GetAll() (interface{}, error) {
	shows, err := s.showRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return shows, nil
}

func (s *showUsecase) Find(id int) (model.Show, error) {
	show, err := s.showRepository.Find(id)

	if err != nil {
		return model.Show{}, err
	}

	return show, nil
}

func (s *showUsecase) Delete(id int) error {
	err := s.showRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *showUsecase) Update(id int, show dto.CreateShowRequest) (model.Show, error) {
	showData, err := s.showRepository.Find(id)

	if err != nil {
		return model.Show{}, err
	}

	showData.MovieID = show.MovieID
	showData.StudioID = show.StudioID
	showData.ShowDate = show.ShowDate
	showData.ShowStart = show.ShowStart
	showData.ShowEnd = show.ShowEnd
	showData.Price = show.Price
	
	e := s.showRepository.Update(id, showData)

	if e != nil {
		return model.Show{}, e
	}

	showUpdated, err := s.showRepository.Find(id)

	if err != nil {
		return model.Show{}, err
	}
	return showUpdated, nil
}
