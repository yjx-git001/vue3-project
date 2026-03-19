package dao

import (
	"backend/app/models"
	"time"

	"gorm.io/gorm"
)

type studyRecordDao struct{}

var StudyRecordDao = new(studyRecordDao)

func (d studyRecordDao) Create(db *gorm.DB, record *models.StudyRecord) error {
	return db.Create(record).Error
}

func (d studyRecordDao) GetTodayDuration(db *gorm.DB, userID uint) (int, error) {
	var total int
	today := time.Now().Format("2006-01-02")
	err := db.Model(&models.StudyRecord{}).
		Select("COALESCE(SUM(duration), 0)").
		Where("user_id = ? AND date = ?", userID, today).
		Scan(&total).Error
	return total, err
}

func (d studyRecordDao) GetTotalDuration(db *gorm.DB, userID uint) (int, error) {
	var total int
	err := db.Model(&models.StudyRecord{}).
		Select("COALESCE(SUM(duration), 0)").
		Where("user_id = ?", userID).
		Scan(&total).Error
	return total, err
}
