package repository

import (
	"miniproject/model"

	"github.com/stretchr/testify/mock"
)

type mockUserRepository struct {
	mock.Mock
}

func NewMockUserRepository() *mockUserRepository {
	return &mockUserRepository{}
}

func (m *mockUserRepository) GetAll() ([]model.User, error) {
	ret := m.Called()
	return ret.Get(0).([]model.User), ret.Error(1)
}

func (m *mockUserRepository) Find(id int) (model.User, error) {
	ret := m.Called()
	return ret.Get(0).(model.User), ret.Error(1)
}

func (m *mockUserRepository) Create(data model.User) error {
	ret := m.Called(data)
	return ret.Error(0)
}

func (m *mockUserRepository) Update(id int, data model.User) error {
	ret := m.Called(id, data)
	return ret.Error(0)
}

func (m *mockUserRepository) Delete(id int) error {
	ret := m.Called(id)
	return ret.Error(0)
}
