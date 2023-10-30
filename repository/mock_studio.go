package repository

import (
	"miniproject/model"
	"github.com/stretchr/testify/mock"
)

type mockStudioRepository struct {
	mock.Mock
}

func NewMockStudioRepository() *mockStudioRepository {
	return &mockStudioRepository{}
}

func (m *mockStudioRepository) GetAll() ([]model.Studio, error) {
	ret := m.Called()
	return ret.Get(0).([]model.Studio), ret.Error(1)
}

func (m *mockStudioRepository) Find(id int) (model.Studio, error) {
	ret := m.Called()
	return ret.Get(0).(model.Studio), ret.Error(1)
}

func (m *mockStudioRepository) Create(data model.Studio) (uint, error) {
	ret := m.Called(data)
	return ret.Get(0).(uint), ret.Error(1)
}

func (m *mockStudioRepository) Delete(id int) error {
	ret := m.Called(id)
	return ret.Error(0)
}

func (m *mockStudioRepository) Update(id int, studio model.Studio) error {
	ret := m.Called(id, studio)
	return ret.Error(0)
}
