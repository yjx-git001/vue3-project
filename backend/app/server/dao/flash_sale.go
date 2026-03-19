package dao

import (
	"backend/app/models"
	"backend/pkg/db"

	"gorm.io/gorm"
)

type flashSaleDao struct{}

var FlashSaleDao = new(flashSaleDao)

func (d flashSaleDao) GetAll(tx *gorm.DB) ([]models.FlashSale, error) {
	if tx == nil {
		tx = db.Db
	}
	var list []models.FlashSale
	err := tx.Preload("Course").Order("created_at DESC").Find(&list).Error
	return list, err
}

func (d flashSaleDao) GetActive(tx *gorm.DB) (*models.FlashSale, error) {
	if tx == nil {
		tx = db.Db
	}
	var item models.FlashSale
	err := tx.Preload("Course").Where("enable = 1").Order("created_at DESC").First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (d flashSaleDao) Create(tx *gorm.DB, data *models.FlashSale) error {
	if tx == nil {
		tx = db.Db
	}
	return tx.Create(data).Error
}

func (d flashSaleDao) SetEnable(tx *gorm.DB, id uint, enable bool) error {
	if tx == nil {
		tx = db.Db
	}
	return tx.Model(&models.FlashSale{}).Where("id = ?", id).Update("enable", enable).Error
}

func (d flashSaleDao) Delete(tx *gorm.DB, id uint) error {
	if tx == nil {
		tx = db.Db
	}
	return tx.Delete(&models.FlashSale{}, id).Error
}

