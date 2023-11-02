package repository

import (
	"miniproject/model"

	"gorm.io/gorm"
)

type MovieRepository interface {
	GetAll() ([]model.Movie, error)
	Find(id int) (model.Movie, error)
	Create(data model.Movie) error
	Delete(id int) error
	Update(id int, movie model.Movie) error
	GetMovieByName(title string) (model.Movie, error)
}

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *movieRepository {
	return &movieRepository{db}
}

func (r *movieRepository) GetAll() ([]model.Movie, error) {
	movies := []model.Movie{}
	if err := r.db.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *movieRepository) Find(id int) (model.Movie, error) {
	movie := model.Movie{}
	if err := r.db.First(&movie, id).Error; err != nil {
		return movie, err
	}
	return movie, nil
}

func (r *movieRepository) Create(data model.Movie) error {
	return r.db.Create(&data).Error
}

func (r *movieRepository) Delete(id int) error {
	movie := model.Movie{}
	if err := r.db.Delete(&movie, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *movieRepository) Update(id int, movie model.Movie) error {
	if err := r.db.Model(&movie).Updates(movie).Error; err != nil {
		return err
	}
	return nil
}

func (r *movieRepository) GetMovieByName(title string) (model.Movie, error) {
	movie := model.Movie{}
	if err := r.db.Where("title = ?", title).Find(&movie).Error; err != nil {
		return movie, err
	}
	return movie, nil
}