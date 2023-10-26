package repository

import (
	"miniproject/model"

	"gorm.io/gorm"
)

type SeatRepository interface {
	GetAll() ([]model.Seat, error)
	Create(data model.Seat) error
}

type seatRepository struct {
	db *gorm.DB
}

func NewSeatRepository(db *gorm.DB) *seatRepository {
	return &seatRepository{db}
}

func (r *seatRepository) GetAll() ([]model.Seat, error) {
	seats := []model.Seat{}
	if err := r.db.Find(&seats).Error; err != nil {
		return nil, err
	}
	return seats, nil
}

func (r *seatRepository) Create(data model.Seat) error {
	return r.db.Create(&data).Error
}