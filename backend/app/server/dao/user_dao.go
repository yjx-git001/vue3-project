package dao

import (
	"backend/app/models"
	"backend/config"

	"gorm.io/gorm"
)

type userDao struct{}

var UserDao = new(userDao)

func (d userDao) GetByPhone(phone string) (models.User, error) {
	var data models.User
	err := config.DB.Where("phone = ?", phone).First(&data).Error
	return data, err
}

func (d userDao) GetByID(id uint) (models.User, error) {
	var data models.User
	err := config.DB.First(&data, id).Error
	return data, err
}

func (d userDao) Create(tx *gorm.DB, data *models.User) error {
	return tx.Create(data).Error
}
