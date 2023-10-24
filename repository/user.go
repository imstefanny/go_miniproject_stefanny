package repository

import (
	"miniproject/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	Find(id int) (model.User, error)
	Create(data model.User) error
	Delete(id int) error
	Update(id int, user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAll() ([]model.User, error) {
	users := []model.User{}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Find(id int) (model.User, error) {
	user := model.User{}
	if err := r.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Create(data model.User) error {
	return r.db.Create(&data).Error
}

func (r *userRepository) Delete(id int) error {
	user := model.User{}
	if err := r.db.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(id int, user model.User) error {
	if err := r.db.Model(&user).Updates(user).Error; err != nil {
		return err
	}
	return nil
}
