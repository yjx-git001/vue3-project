package dao

import (
	"backend/app/models"

	"gorm.io/gorm"
)

type couponDao struct{}

var CouponDao = new(couponDao)

func (d couponDao) BatchCreate(db *gorm.DB, coupons []models.Coupon) error {
	return db.Create(&coupons).Error
}

func (d couponDao) GetList(db *gorm.DB) ([]models.Coupon, error) {
	var list []models.Coupon
	err := db.Order("created_at DESC").Find(&list).Error
	return list, err
}

func (d couponDao) GetByUserID(db *gorm.DB, userID uint) ([]models.Coupon, error) {
	var list []models.Coupon
	err := db.Where("user_id = ?", userID).Order("created_at DESC").Find(&list).Error
	return list, err
}
