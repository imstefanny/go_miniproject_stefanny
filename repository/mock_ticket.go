package repository

import (
	"miniproject/model"
	"github.com/stretchr/testify/mock"
)

type mockTicketRepository struct {
	mock.Mock
}

func NewMockTicketRepository() *mockTicketRepository {
	return &mockTicketRepository{}
}

func (m *mockTicketRepository) GetAll() ([]model.Ticket, error) {
	ret := m.Called()
	return ret.Get(0).([]model.Ticket), ret.Error(1)
}

func (m *mockTicketRepository) Find(id int) (model.Ticket, error) {
	ret := m.Called()
	return ret.Get(0).(model.Ticket), ret.Error(1)
}

func (m *mockTicketRepository) Create(data model.Ticket) error {
	ret := m.Called(data)
	return ret.Error(0)
}

func (m *mockTicketRepository) GetAvailableTickets(showID int) ([]model.Ticket, error) {
	ret := m.Called()
	return ret.Get(0).([]model.Ticket), ret.Error(1)
}

func (m *mockTicketRepository) GetTicket(showID, seatID uint) (model.Ticket, error) {
	ret := m.Called()
	return ret.Get(0).(model.Ticket), ret.Error(1)
}
