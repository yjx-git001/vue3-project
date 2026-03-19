package dao

import (
	"backend/app/models"

	"gorm.io/gorm"
)

type orderDao struct{}

var OrderDao = new(orderDao)

func (d orderDao) Create(db *gorm.DB, order *models.Order) error {
	return db.Create(order).Error
}

func (d orderDao) GetByUserID(db *gorm.DB, userID uint) ([]models.Order, error) {
	var list []models.Order
	err := db.Where("user_id = ?", userID).Order("created_at DESC").Find(&list).Error
	return list, err
}

func (d orderDao) GetByOrderNo(db *gorm.DB, orderNo string) (*models.Order, error) {
	var order models.Order
	err := db.Where("order_no = ?", orderNo).First(&order).Error
	return &order, err
}

func (d orderDao) UpdateStatus(db *gorm.DB, id uint, payType models.PayType, status models.OrderStatus) error {
	return db.Model(&models.Order{}).Where("id = ?", id).Updates(map[string]interface{}{
		"pay_type": payType,
		"status":   status,
	}).Error
}

func (d orderDao) CancelIfPending(db *gorm.DB, id uint) error {
	return db.Model(&models.Order{}).
		Where("id = ? AND status = ?", id, models.OrderStatusPending).
		Update("status", models.OrderStatusCancelled).Error
}

func (d orderDao) ExistsByCourseAndUser(db *gorm.DB, userID uint, courseEk int64) (bool, error) {
	var count int64
	err := db.Model(&models.Order{}).Where("user_id = ? AND course_ek = ? AND status = ?", userID, courseEk, models.OrderStatusPaid).Count(&count).Error
	return count > 0, err
}
