package models

import "time"

type StudyRecord struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"userId" gorm:"index;comment:用户ID"`
	Duration  int       `json:"duration" gorm:"comment:学习时长(秒)"`
	Date      time.Time `json:"date" gorm:"type:date;index;comment:学习日期"`
	CreatedAt time.Time `json:"createdAt"`
}

func (StudyRecord) TableName() string {
	return "study_records"
}
