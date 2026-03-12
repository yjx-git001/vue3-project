package dao

import (
	"backend/app/models"
	"backend/pkg/db"

	"gorm.io/gorm"
)

type userDao struct{}

var UserDao = new(userDao)

func (d userDao) GetByPhone(tx *gorm.DB, phone string) (models.User, error) {
	if tx == nil {
		tx = db.Db
	}
	var data models.User
	err := tx.Debug().Where("phone = ?", phone).First(&data).Error
	return data, err
}

func (d userDao) GetByID(tx *gorm.DB, id uint) (models.User, error) {
	if tx == nil {
		tx = db.Db
	}
	var data models.User
	err := tx.Debug().First(&data, id).Error
	return data, err
}

func (d userDao) Create(tx *gorm.DB, data *models.User) error {
	if tx == nil {
		tx = db.Db
	}
	return tx.Debug().Create(data).Error
}
