package usecase

import (
	"miniproject/dto"
	"miniproject/model"
	"miniproject/repository"
	"reflect"

	// "miniproject/middlewares"
	"errors"
)

type StudioUsecase interface {
	Create(studio dto.CreateStudioRequest) error
	GetAll() (interface{}, error)
	Find(id int) (interface{}, error)
	Delete(id int) error
	Update(id int, studio dto.CreateStudioRequest) (model.Studio, error)
}

type studioUsecase struct {
	studioRepository		repository.StudioRepository
}

func NewStudioUsecase(studioRepo repository.StudioRepository) *studioUsecase {
	return &studioUsecase{studioRepository: studioRepo}
}

func validateCreateStudioRequest(req dto.CreateStudioRequest) error {
	val := reflect.ValueOf(req)
	for i := 0; i < val.NumField(); i++ {
			if isEmptyField(val.Field(i)) {
					return errors.New("Field can't be empty")
			}
	}
	return nil
}

func (s *studioUsecase) Create(studio dto.CreateStudioRequest) error {
	e := validateCreateStudioRequest(studio)
	
	if e!= nil {
		return e
	}

	studioData := model.Studio{
		Name: studio.Name,
		Capacity: studio.Capacity,
		CinemaID: studio.CinemaID,
	}
	err := s.studioRepository.Create(studioData)

	if err != nil {
		return err
	}

	return nil
}

func (s *studioUsecase) GetAll() (interface{}, error) {
	studios, err := s.studioRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return studios, nil
}

func (s *studioUsecase) Find(id int) (interface{}, error) {
	studio, err := s.studioRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return studio, nil
}

func (s *studioUsecase) Delete(id int) error {
	err := s.studioRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *studioUsecase) Update(id int, studio dto.CreateStudioRequest) (model.Studio, error) {
	studioData, err := s.studioRepository.Find(id)

	if err != nil {
		return model.Studio{}, err
	}

	studioData.Name = studio.Name
	studioData.Capacity = studio.Capacity
	studioData.CinemaID = studio.CinemaID
	
	e := s.studioRepository.Update(id, studioData)

	if e != nil {
		return model.Studio{}, e
	}

	studioUpdated, err := s.studioRepository.Find(id)

	if err != nil {
		return model.Studio{}, err
	}
	return studioUpdated, nil
}
