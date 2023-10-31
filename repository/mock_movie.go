package repository

import (
	"miniproject/model"

	"github.com/stretchr/testify/mock"
)

type mockMovieRepository struct {
	mock.Mock
}

func NewMockMovieRepository() *mockMovieRepository {
	return &mockMovieRepository{}
}

func (m *mockMovieRepository) GetAll() ([]model.Movie, error) {
	ret := m.Called()
	return ret.Get(0).([]model.Movie), ret.Error(1)
}

func (m *mockMovieRepository) Find(id int) (model.Movie, error) {
	ret := m.Called()
	return ret.Get(0).(model.Movie), ret.Error(1)
}

func (m *mockMovieRepository) Create(data model.Movie) error {
	ret := m.Called(data)
	return ret.Error(0)
}

func (m *mockMovieRepository) Update(id int, data model.Movie) error {
	ret := m.Called(id, data)
	return ret.Error(0)
}

func (m *mockMovieRepository) Delete(id int) error {
	ret := m.Called(id)
	return ret.Error(0)
}

func (m *mockMovieRepository) GetMovieByName(title string) (model.Movie, error) {
	ret := m.Called()
	return ret.Get(0).(model.Movie), ret.Error(1)
}
