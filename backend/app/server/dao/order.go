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

func (d orderDao) GetByUserID(db *gorm.DB, userID uint, status *models.OrderStatus) ([]models.Order, error) {
	var list []models.Order
	query := db.Where("user_id = ?", userID)
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	err := query.Order("created_at DESC").Find(&list).Error
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

func (d orderDao) GetPendingByCourseAndUser(db *gorm.DB, userID uint, courseEk int64) (*models.Order, error) {
	var order models.Order
	err := db.Where("user_id = ? AND course_ek = ? AND status = ?", userID, courseEk, models.OrderStatusPending).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (d orderDao) ExistsByCourseAndUser(db *gorm.DB, userID uint, courseEk int64) (bool, error) {
	// 检查是否直接购买了该课程
	var count int64
	err := db.Model(&models.Order{}).Where("user_id = ? AND course_ek = ? AND status = ?", userID, courseEk, models.OrderStatusPaid).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}

	// 检查该课程是否属于某个体系课程
	var course models.Course
	if err := db.Select("parent_ek").Where("ek = ?", courseEk).First(&course).Error; err != nil {
		return false, nil
	}

	// 如果有父课程，检查是否购买了父课程
	if course.ParentEk > 0 {
		err = db.Model(&models.Order{}).Where("user_id = ? AND course_ek = ? AND status = ?", userID, course.ParentEk, models.OrderStatusPaid).Count(&count).Error
		if err != nil {
			return false, err
		}
		return count > 0, nil
	}

	return false, nil
}
