package dao

import (
	"backend/app/models"
	"backend/pkg/db"
	"strings"

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

func (d userDao) ListOptions(tx *gorm.DB, keyword string, limit int) ([]models.User, error) {
	if tx == nil {
		tx = db.Db
	}
	if limit <= 0 || limit > 200 {
		limit = 100
	}

	query := tx.Debug().Select("id, name, nickname, phone").Order("id DESC").Limit(limit)
	kw := strings.TrimSpace(keyword)
	if kw != "" {
		like := "%" + kw + "%"
		query = query.Where("name LIKE ? OR nickname LIKE ? OR phone LIKE ?", like, like, like)
	}

	var list []models.User
	err := query.Find(&list).Error
	return list, err
}
