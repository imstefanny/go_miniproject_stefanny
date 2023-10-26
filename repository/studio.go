package repository

import (
	"miniproject/model"

	"gorm.io/gorm"
)

type StudioRepository interface {
	GetAll() ([]model.Studio, error)
	Find(id int) (model.Studio, error)
	Create(data model.Studio) error
	Delete(id int) error
	Update(id int, studio model.Studio) error
}

type studioRepository struct {
	db *gorm.DB
}

func NewStudioRepository(db *gorm.DB) *studioRepository {
	return &studioRepository{db}
}

func (r *studioRepository) GetAll() ([]model.Studio, error) {
	studios := []model.Studio{}
	if err := r.db.Find(&studios).Error; err != nil {
		return nil, err
	}
	return studios, nil
}

func (r *studioRepository) Find(id int) (model.Studio, error) {
	studio := model.Studio{}
	if err := r.db.First(&studio, id).Error; err != nil {
		return studio, err
	}
	return studio, nil
}

func (r *studioRepository) Create(data model.Studio) error {
	return r.db.Create(&data).Error
}

func (r *studioRepository) Delete(id int) error {
	studio := model.Studio{}
	if err := r.db.Delete(&studio, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *studioRepository) Update(id int, studio model.Studio) error {
	if err := r.db.Model(&studio).Updates(studio).Error; err != nil {
		return err
	}
	return nil
}
