package repository

import (
	"miniproject/model"

	"gorm.io/gorm"
)

type CinemaRepository interface {
	GetAll() ([]model.Cinema, error)
	Find(id int) (model.Cinema, error)
	Create(data model.Cinema) error
	Delete(id int) error
	Update(id int, cinema model.Cinema) error
}

type cinemaRepository struct {
	db *gorm.DB
}

func NewCinemaRepository(db *gorm.DB) *cinemaRepository {
	return &cinemaRepository{db}
}

func (r *cinemaRepository) GetAll() ([]model.Cinema, error) {
	cinemas := []model.Cinema{}
	if err := r.db.Find(&cinemas).Error; err != nil {
		return nil, err
	}
	return cinemas, nil
}

func (r *cinemaRepository) Find(id int) (model.Cinema, error) {
	cinema := model.Cinema{}
	if err := r.db.First(&cinema, id).Error; err != nil {
		return cinema, err
	}
	return cinema, nil
}

func (r *cinemaRepository) Create(data model.Cinema) error {
	return r.db.Create(&data).Error
}

func (r *cinemaRepository) Delete(id int) error {
	cinema := model.Cinema{}
	if err := r.db.Delete(&cinema, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *cinemaRepository) Update(id int, cinema model.Cinema) error {
	if err := r.db.Model(&cinema).Updates(cinema).Error; err != nil {
		return err
	}
	return nil
}
