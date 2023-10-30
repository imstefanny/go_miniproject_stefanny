package repository

import (
	"miniproject/model"
	"github.com/stretchr/testify/mock"
)

type mockCinemaRepository struct {
	mock.Mock
}

func NewMockCinemaRepository() *mockCinemaRepository {
	return &mockCinemaRepository{}
}

func (m *mockCinemaRepository) GetAll() ([]model.Cinema, error) {
	ret := m.Called()
	return ret.Get(0).([]model.Cinema), ret.Error(1)
}

func (m *mockCinemaRepository) Find(id int) (model.Cinema, error) {
	ret := m.Called()
	return ret.Get(0).(model.Cinema), ret.Error(1)
}

func (m *mockCinemaRepository) Create(data model.Cinema) error {
	ret := m.Called(data)
	return ret.Error(0)
}

func (m *mockCinemaRepository) Delete(id int) error {
	ret := m.Called(id)
	return ret.Error(0)
}

func (m *mockCinemaRepository) Update(id int, cinema model.Cinema) error {
	ret := m.Called(id, cinema)
	return ret.Error(0)
}
