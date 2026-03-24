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

func (d studyRecordDao) GetCourseDuration(db *gorm.DB, userID uint, courseEk int64) (int, error) {
	var total int
	err := db.Model(&models.StudyRecord{}).
		Select("COALESCE(SUM(duration), 0)").
		Where("user_id = ? AND course_ek = ?", userID, courseEk).
		Scan(&total).Error
	return total, err
}

type courseDailyDurationRow struct {
	Date  time.Time `gorm:"column:date"`
	Total int       `gorm:"column:total"`
}

func (d studyRecordDao) GetCourseDailyDurations(db *gorm.DB, userID uint, courseEk int64) ([]courseDailyDurationRow, error) {
	var rows []courseDailyDurationRow
	err := db.Model(&models.StudyRecord{}).
		Select("date, COALESCE(SUM(duration), 0) AS total").
		Where("user_id = ? AND course_ek = ?", userID, courseEk).
		Group("date").
		Order("date ASC").
		Scan(&rows).Error
	return rows, err
}

func (d studyRecordDao) GetCourseQualifiedDate(db *gorm.DB, userID uint, courseEk int64, threshold int) (*time.Time, error) {
	rows, err := d.GetCourseDailyDurations(db, userID, courseEk)
	if err != nil {
		return nil, err
	}
	total := 0
	for _, r := range rows {
		total += r.Total
		if total >= threshold {
			date := r.Date
			return &date, nil
		}
	}
	return nil, nil
}
