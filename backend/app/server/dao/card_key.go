package dao

import (
	"backend/app/models"

	"gorm.io/gorm"
)

type cardKeyDao struct{}

var CardKeyDao = new(cardKeyDao)

func (d cardKeyDao) BatchCreate(db *gorm.DB, keys []models.CardKey) error {
	return db.Create(&keys).Error
}

func (d cardKeyDao) GetByCode(db *gorm.DB, code string) (*models.CardKey, error) {
	var key models.CardKey
	err := db.Where("code = ?", code).First(&key).Error
	return &key, err
}

func (d cardKeyDao) MarkUsed(db *gorm.DB, id uint, userID uint) error {
	return db.Model(&models.CardKey{}).Where("id = ?", id).Updates(map[string]interface{}{
		"used":    true,
		"user_id": userID,
	}).Error
}

func (d cardKeyDao) GetList(db *gorm.DB) ([]models.CardKey, error) {
	var list []models.CardKey
	err := db.Order("created_at DESC").Find(&list).Error
	return list, err
}
