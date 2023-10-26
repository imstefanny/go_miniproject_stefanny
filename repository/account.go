package repository

import (
	"miniproject/model"

	"gorm.io/gorm"
)

type AccountRepository interface {
	GetAll() ([]model.Account, error)
	Find(id int) (model.Account, error)
	Create(data model.Account) error
	Delete(id int) error
	Update(id int, account model.Account) error
	Login(data model.Account) (model.Account, error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *accountRepository {
	return &accountRepository{db}
}

func (r *accountRepository) Create(data model.Account) error {
	return r.db.Create(&data).Error
}

func (r *accountRepository) GetAll() ([]model.Account, error) {
	accounts := []model.Account{}
	if err := r.db.Preload("UserID").Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r *accountRepository) Find(id int) (model.Account, error) {
	account := model.Account{}
	if err := r.db.Preload("UserID").First(&account, id).Error; err != nil {
		return account, err
	}
	return account, nil
}

func (r *accountRepository) Delete(id int) error {
	account := model.Account{}
	if err := r.db.Delete(&account, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *accountRepository) Update(id int, account model.Account) error {
	if err := r.db.Model(&account).Updates(account).Error; err != nil {
		return err
	}
	return nil
}

func (r *accountRepository) Login(data model.Account) (model.Account, error) {
	account := model.Account{}
	err := r.db.Where("username = ? AND password = ?", data.Username, data.Password).First(&account).Error
	if err != nil {
		return model.Account{}, err
	}
	return account, nil
}
