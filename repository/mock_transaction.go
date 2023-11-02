package repository

import (
	"miniproject/model"

	"github.com/stretchr/testify/mock"
)

type mockTransactionRepository struct {
	mock.Mock
}

func NewMockTransactionRepository() *mockTransactionRepository {
	return &mockTransactionRepository{}
}

func (m *mockTransactionRepository) GetAll() ([]model.Transaction, error) {
	ret := m.Called()
	return ret.Get(0).([]model.Transaction), ret.Error(1)
}

func (m *mockTransactionRepository) Create(data model.Transaction) error {
	ret := m.Called(data)
	return ret.Error(0)
}

func (m *mockTransactionRepository) GetByInvoice(invoice string) (model.Transaction, error) {
	ret := m.Called()
	return ret.Get(0).(model.Transaction), ret.Error(1)
}

func (m *mockTransactionRepository) CheckSeatValidity(ticketID uint) (bool, error) {
	ret := m.Called()
	return ret.Get(0).(bool), ret.Error(1)
}
