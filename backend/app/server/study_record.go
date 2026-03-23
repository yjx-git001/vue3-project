package services

import (
	"backend/app/models"
	"backend/app/server/dao"
	"backend/app/server/dto"
	"backend/pkg/db"
	"time"
)

type studyRecordService struct{}

var StudyRecordService = new(studyRecordService)

func (s studyRecordService) Add(userID uint, req *dto.AddStudyRecordReq) error {
	record := &models.StudyRecord{
		UserID:   userID,
		CourseEk: req.CourseEk,
		Duration: req.Duration,
		Date:     time.Now(),
	}
	return dao.StudyRecordDao.Create(db.Db, record)
}

func (s studyRecordService) GetStats(userID uint) (*dto.StudyStatsResp, error) {
	today, err := dao.StudyRecordDao.GetTodayDuration(db.Db, userID)
	if err != nil {
		return nil, err
	}
	total, err := dao.StudyRecordDao.GetTotalDuration(db.Db, userID)
	if err != nil {
		return nil, err
	}
	return &dto.StudyStatsResp{
		TodayDuration: today,
		TotalDuration: total,
	}, nil
}

func (s studyRecordService) GetCourseDuration(userID uint, courseEk int64) (*dto.CourseDurationResp, error) {
	total, err := dao.StudyRecordDao.GetCourseDuration(db.Db, userID, courseEk)
	if err != nil {
		return nil, err
	}
	return &dto.CourseDurationResp{
		CourseEk:      courseEk,
		TotalDuration: total,
	}, nil
}
