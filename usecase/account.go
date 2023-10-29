package usecase

import (
	"errors"
	"miniproject/dto"
	"miniproject/middlewares"
	"miniproject/model"
	"miniproject/repository"
	"reflect"
)

type AccountUsecase interface {
	GetAll() (interface{}, error)
	Find(id int) (interface{}, error)
	Create(account dto.CreateAccountRequest) error
	Delete(id int) error
	Update(id int, user dto.CreateAccountRequest) (model.Account, error)
	Login(data dto.LoginAccountRequest) (interface{}, error)
}

type accountUsecase struct {
	accountRepository		repository.AccountRepository
}

func NewAccountUsecase(accountRepo repository.AccountRepository) *accountUsecase {
	return &accountUsecase{accountRepository: accountRepo}
}

func validateCreateAccountRequest(req dto.CreateAccountRequest) error {
	val := reflect.ValueOf(req)
	for i := 0; i < val.NumField(); i++ {
			if isEmptyField(val.Field(i)) {
					return errors.New("Field can't be empty")
			}
	}
	return nil
}

func validateLoginAccountRequest(req dto.LoginAccountRequest) error {
	val := reflect.ValueOf(req)
	for i := 0; i < val.NumField(); i++ {
			if isEmptyField(val.Field(i)) {
					return errors.New("Field can't be empty")
			}
	}
	return nil
}

func (s *accountUsecase) Create(account dto.CreateAccountRequest) error {
	e := validateCreateAccountRequest(account)
	
	if e!= nil {
		return e
	}
	
	accountData := model.Account{
		Username: account.Username,
		Email: account.Email,
		Password: account.Password,
		Role: account.Role,
	}
	err := s.accountRepository.Create(accountData)

	if err != nil {
		return err
	}

	return nil
}

func (s *accountUsecase) GetAll() (interface{}, error) {
	accounts, err := s.accountRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s *accountUsecase) Find(id int) (interface{}, error) {
	account, err := s.accountRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *accountUsecase) Delete(id int) error {
	err := s.accountRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *accountUsecase) Update(id int, account dto.CreateAccountRequest) (model.Account, error) {
	accountData, err := s.accountRepository.Find(id)

	if err != nil {
		return model.Account{}, err
	}

	accountData.Username = account.Username
	accountData.Email = account.Email
	accountData.Password = account.Password
	accountData.Role = account.Role
	
	e := s.accountRepository.Update(id, accountData)

	if e != nil {
		return model.Account{}, e
	}

	accountUpdated, err := s.accountRepository.Find(id)

	if err != nil {
		return model.Account{}, err
	}
	return accountUpdated, nil
}


func (s *accountUsecase) Login(data dto.LoginAccountRequest) (interface{}, error) {
	e := validateLoginAccountRequest(data)
	
	if e!= nil {
		return nil, e
	}

	accountData := model.Account{
		Username: data.Username,
		Email: data.Email,
		Password: data.Password,
	}
	account, err := s.accountRepository.Login(accountData)

	if err != nil {
		return nil, err
	}

	token, err := middlewares.CreateToken(account.Username, account.Password, account.Role)

	if err != nil {
		return nil, err
	}

	accountResponse := model.AccountResponse{
		ID			: account.ID,
		Username: account.Username,
		Email		: account.Email,
		Token 	: token,
	}

	return accountResponse, nil
}
