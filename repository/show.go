package repository

import (
	"miniproject/model"

	"gorm.io/gorm"
)

type ShowRepository interface {
	GetAll() ([]model.Show, error)
	Find(id int) (model.Show, error)
	Create(data model.Show) (uint, error)
	Delete(id int) error
	Update(id int, show model.Show) error
}

type showRepository struct {
	db *gorm.DB
}

func NewShowRepository(db *gorm.DB) *showRepository {
	return &showRepository{db}
}

func (r *showRepository) GetAll() ([]model.Show, error) {
	shows := []model.Show{}
	if err := r.db.Find(&shows).Error; err != nil {
		return nil, err
	}
	return shows, nil
}

func (r *showRepository) Find(id int) (model.Show, error) {
	show := model.Show{}
	if err := r.db.First(&show, id).Error; err != nil {
		return show, err
	}
	return show, nil
}

func (r *showRepository) Create(data model.Show) (uint, error) {
	err := r.db.Create(&data)
	if err.Error != nil {
		return 0, err.Error
	}
	return data.ID, nil
}

func (r *showRepository) Delete(id int) error {
	show := model.Show{}
	if err := r.db.Delete(&show, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *showRepository) Update(id int, show model.Show) error {
	if err := r.db.Model(&show).Updates(show).Error; err != nil {
		return err
	}
	return nil
}
