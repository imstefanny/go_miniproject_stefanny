package repository

import (
	"miniproject/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetAll() ([]model.Transaction, error)
	Create(data model.Transaction) error
	GetByInvoice(invoice string) (model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) GetAll() ([]model.Transaction, error) {
	transactions := []model.Transaction{}
	if err := r.db.Preload("Tickets").Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *transactionRepository) Create(data model.Transaction) error {
	return r.db.Create(&data).Error
}

func(r *transactionRepository) GetByInvoice(invoice string) (model.Transaction, error) {
	transaction := model.Transaction{}
	if err := r.db.Preload("Tickets").Where("ticket_code = ?", invoice).First(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}
