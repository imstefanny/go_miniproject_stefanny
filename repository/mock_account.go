package repository

import (
	"miniproject/model"

	"github.com/stretchr/testify/mock"
)

type mockAccountRepository struct {
	mock.Mock
}

func NewMockAccountRepository() *mockAccountRepository {
	return &mockAccountRepository{}
}

func (m *mockAccountRepository) GetAll() ([]model.Account, error) {
	ret := m.Called()
	return ret.Get(0).([]model.Account), ret.Error(1)
}

func (m *mockAccountRepository) Find(id int) (model.Account, error) {
	ret := m.Called()
	return ret.Get(0).(model.Account), ret.Error(1)
}

func (m *mockAccountRepository) Create(data model.Account) error {
	ret := m.Called(data)
	return ret.Error(0)
}

func (m *mockAccountRepository) Update(id int, data model.Account) error {
	ret := m.Called(id, data)
	return ret.Error(0)
}

func (m *mockAccountRepository) Delete(id int) error {
	ret := m.Called(id)
	return ret.Error(0)
}

func (m *mockAccountRepository) Login(data model.Account) (model.Account, error) {
	ret := m.Called(data)
	return ret.Get(0).(model.Account), ret.Error(1)
}
