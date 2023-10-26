package usecase

import (
	"miniproject/dto"
	"miniproject/model"
	"miniproject/repository"
	"reflect"

	// "miniproject/middlewares"
	"errors"
)

type UserUsecase interface {
	Create(user dto.CreateUserRequest) error
	GetAll() (interface{}, error)
	Find(id int) (interface{}, error)
	Delete(id int) error
	Update(id int, user dto.CreateUserRequest) (model.User, error)
}

type userUsecase struct {
	userRepository		repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func validateCreateUserRequest(req dto.CreateUserRequest) error {
	val := reflect.ValueOf(req)
	for i := 0; i < val.NumField(); i++ {
			if isEmptyField(val.Field(i)) {
					return errors.New("Field can't be empty")
			}
	}
	return nil
}

func (s *userUsecase) Create(user dto.CreateUserRequest) error {
	e := validateCreateUserRequest(user)
	
	if e!= nil {
		return e
	}

	userData := model.User{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Phone: user.Phone,
		Address: user.Address,
		Gender: user.Gender,
		DateOfBirth: user.DateOfBirth,
		Pin: user.Pin,
		AccountID: user.AccountID,
	}
	err := s.userRepository.Create(userData)

	if err != nil {
		return err
	}

	return nil
}

func (s *userUsecase) GetAll() (interface{}, error) {
	users, err := s.userRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userUsecase) Find(id int) (interface{}, error) {
	user, err := s.userRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userUsecase) Delete(id int) error {
	err := s.userRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *userUsecase) Update(id int, user dto.CreateUserRequest) (model.User, error) {
	userData, err := s.userRepository.Find(id)

	if err != nil {
		return model.User{}, err
	}

	userData.FirstName = user.FirstName
	userData.LastName = user.LastName
	userData.Email = user.Email
	userData.Phone = user.Phone
	userData.Address = user.Address
	userData.Gender = user.Gender
	userData.DateOfBirth = user.DateOfBirth
	userData.Pin = user.Pin
	userData.AccountID = user.AccountID
	
	e := s.userRepository.Update(id, userData)

	if e != nil {
		return model.User{}, e
	}

	userUpdated, err := s.userRepository.Find(id)

	if err != nil {
		return model.User{}, err
	}
	return userUpdated, nil
}
