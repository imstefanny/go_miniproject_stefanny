package repository

import (
	"miniproject/model"
	"github.com/stretchr/testify/mock"
)

type mockSeatRepository struct {
	mock.Mock
}

func NewMockSeatRepository() *mockSeatRepository {
	return &mockSeatRepository{}
}

func (m *mockSeatRepository) GetAll() ([]model.Seat, error) {
	ret := m.Called()
	return ret.Get(0).([]model.Seat), ret.Error(1)
}

func (m *mockSeatRepository) Find(studio_id int) ([]model.Seat, error) {
	ret := m.Called()
	return ret.Get(0).([]model.Seat), ret.Error(1)
}

func (m *mockSeatRepository) Create(data model.Seat) error {
	ret := m.Called(data)
	return ret.Error(0)
}

func (m *mockSeatRepository) GetAvailableSeats(seat_id int) (model.Seat, error) {
	ret := m.Called()
	return ret.Get(0).(model.Seat), ret.Error(1)
}
