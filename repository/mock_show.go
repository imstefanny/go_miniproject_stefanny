package repository

import (
	"miniproject/model"
	"github.com/stretchr/testify/mock"
)

type mockShowRepository struct {
	mock.Mock
}

func NewMockShowRepository() *mockShowRepository {
	return &mockShowRepository{}
}

func (m *mockShowRepository) GetAll() ([]model.Show, error) {
	ret := m.Called()
	return ret.Get(0).([]model.Show), ret.Error(1)
}

func (m *mockShowRepository) Find(id int) (model.Show, error) {
	ret := m.Called()
	return ret.Get(0).(model.Show), ret.Error(1)
}

func (m *mockShowRepository) Create(data model.Show) (uint, error) {
	ret := m.Called(data)
	return ret.Get(0).(uint), ret.Error(1)
}

func (m *mockShowRepository) Delete(id int) error {
	ret := m.Called(id)
	return ret.Error(0)
}

func (m *mockShowRepository) Update(id int, show model.Show) error {
	ret := m.Called(id, show)
	return ret.Error(0)
}
