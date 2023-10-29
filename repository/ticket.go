package repository

import (
	"miniproject/model"

	"gorm.io/gorm"
)

type TicketRepository interface {
	GetAll() ([]model.Ticket, error)
	Create(data model.Ticket) error
	GetTicket(showID, seatID uint) (model.Ticket, error)
	Find(id int) (model.Ticket, error)
	GetAvailableTickets(showID int) ([]model.Ticket, error)
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

func (r *ticketRepository) Find(id int) (model.Ticket, error) {
	ticket := model.Ticket{}
	if err := r.db.First(&ticket, id).Error; err != nil {
		return ticket, err
	}
	return ticket, nil
}

func (r *ticketRepository) Create(data model.Ticket) error {
	return r.db.Create(&data).Error
}

func (r *ticketRepository) GetTicket(showID, seatID uint) (model.Ticket, error) {
	ticket := model.Ticket{}
	if err := r.db.Where("show_id = ? AND seat_id = ?", showID, seatID).First(&ticket).Error; err != nil {
		return ticket, err
	}
	return ticket, nil
}

func (r *ticketRepository) GetAvailableTickets(showID int) ([]model.Ticket, error) {
	tickets := []model.Ticket{}
	if err := r.db.Where("show_id = ? AND transaction_id is null", showID).Find(&tickets).Error; err != nil {
		return tickets, err
	}
	return tickets, nil
}
