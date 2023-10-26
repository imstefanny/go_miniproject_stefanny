package repository

import (
	"miniproject/model"

	"gorm.io/gorm"
)

type TicketRepository interface {
	GetAll() ([]model.Ticket, error)
	Create(data model.Ticket) error
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *ticketRepository {
	return &ticketRepository{db}
}

func (r *ticketRepository) GetAll() ([]model.Ticket, error) {
	tickets := []model.Ticket{}
	if err := r.db.Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *ticketRepository) Create(data model.Ticket) error {
	return r.db.Create(&data).Error
}
