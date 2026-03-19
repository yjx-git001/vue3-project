package dao

import (
	"backend/app/models"
	"backend/pkg/db"

	"gorm.io/gorm"
)

type bannerDao struct{}

var BannerDao = new(bannerDao)

func (d bannerDao) GetAll(tx *gorm.DB) ([]models.Banner, error) {
	if tx == nil {
		tx = db.Db
	}
	var list []models.Banner
	err := tx.Where("enable = ?", true).Order("sort ASC, created_at DESC").Find(&list).Error
	return list, err
}

func (d bannerDao) Create(tx *gorm.DB, data *models.Banner) error {
	if tx == nil {
		tx = db.Db
	}
	return tx.Create(data).Error
}

func (d bannerDao) Delete(tx *gorm.DB, id uint) error {
	if tx == nil {
		tx = db.Db
	}
	return tx.Delete(&models.Banner{}, id).Error
}
