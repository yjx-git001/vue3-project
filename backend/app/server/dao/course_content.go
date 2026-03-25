package dao

import (
	"backend/app/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

// ===== 题目 DAO =====

type questionDao struct{}

var QuestionDao = new(questionDao)

func (d questionDao) Create(tx *gorm.DB, q *models.Question) error {
	return tx.Create(q).Error
}

func (d questionDao) ListByCourseEk(tx *gorm.DB, courseEk int64) ([]models.Question, error) {
	var list []models.Question
	err := tx.Where("course_ek = ? AND deleted_at IS NULL", courseEk).
		Order("question_type ASC, id ASC").
		Find(&list).Error
	return list, err
}

func (d questionDao) Delete(tx *gorm.DB, id uint) error {
	return tx.Delete(&models.Question{}, id).Error
}

// ===== 视频 DAO =====

type courseVideoDao struct{}

var CourseVideoDao = new(courseVideoDao)

func (d courseVideoDao) ListByCourseEk(tx *gorm.DB, courseEk int64) ([]models.CourseVideo, error) {
	var list []models.CourseVideo
	err := tx.Where("course_ek = ? AND deleted_at IS NULL", courseEk).
		Order("sort ASC").
		Find(&list).Error
	return list, err
}

// SaveAll 全量覆盖：先删除旧数据再批量插入
func (d courseVideoDao) SaveAll(tx *gorm.DB, courseEk int64, items []models.CourseVideo) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		// 软删除旧数据
		if err := tx.Where("course_ek = ?", courseEk).Delete(&models.CourseVideo{}).Error; err != nil {
			return err
		}
		if len(items) == 0 {
			return nil
		}
		return tx.Create(&items).Error
	})
}

// ===== 附件 DAO =====

type courseAttachmentDao struct{}

var CourseAttachmentDao = new(courseAttachmentDao)

func (d courseAttachmentDao) ListByCourseEk(tx *gorm.DB, courseEk int64) ([]models.CourseAttachment, error) {
	var list []models.CourseAttachment
	err := tx.Where("course_ek = ? AND deleted_at IS NULL", courseEk).
		Order("sort ASC").
		Find(&list).Error
	return list, err
}

func (d courseAttachmentDao) SaveAll(tx *gorm.DB, courseEk int64, items []models.CourseAttachment) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("course_ek = ?", courseEk).Delete(&models.CourseAttachment{}).Error; err != nil {
			return err
		}
		if len(items) == 0 {
			return nil
		}
		return tx.Create(&items).Error
	})
}

// ===== 笔记 DAO =====

type courseNotesDao struct{}

var CourseNotesDao = new(courseNotesDao)

func (d courseNotesDao) GetByCourseEk(tx *gorm.DB, courseEk int64) (*models.CourseNotes, error) {
	var notes models.CourseNotes
	err := tx.Where("course_ek = ?", courseEk).First(&notes).Error
	if err != nil {
		return nil, err
	}
	return &notes, nil
}

func (d courseNotesDao) Save(tx *gorm.DB, courseEk int64, content string) error {
	var notes models.CourseNotes
	result := tx.Where("course_ek = ?", courseEk).First(&notes)
	if result.Error != nil {
		// 新建
		return tx.Create(&models.CourseNotes{CourseEk: courseEk, Content: content}).Error
	}
	// 更新
	return tx.Model(&notes).Update("content", content).Error
}

// ===== 模拟考试配置 DAO =====

type mockExamConfigDao struct{}

var MockExamConfigDao = new(mockExamConfigDao)

func (d mockExamConfigDao) GetByCourseEk(tx *gorm.DB, courseEk int64) (*models.MockExamConfig, error) {
	var cfg models.MockExamConfig
	err := tx.Where("course_ek = ?", courseEk).First(&cfg).Error
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (d mockExamConfigDao) Save(tx *gorm.DB, courseEk int64, single, multiple, judge int) error {
	var cfg models.MockExamConfig
	result := tx.Where("course_ek = ?", courseEk).First(&cfg)
	if result.Error != nil {
		// 新建
		return tx.Create(&models.MockExamConfig{
			CourseEk:      courseEk,
			SingleCount:   single,
			MultipleCount: multiple,
			JudgeCount:    judge,
		}).Error
	}
	// 更新
	return tx.Model(&cfg).Updates(map[string]interface{}{
		"single_count":   single,
		"multiple_count": multiple,
		"judge_count":    judge,
	}).Error
}

// ===== 模拟考试记录 DAO =====

type mockExamRecordDao struct{}

var MockExamRecordDao = new(mockExamRecordDao)

func (d mockExamRecordDao) Save(tx *gorm.DB, record *models.MockExamRecord) error {
	return tx.Create(record).Error
}

func (d mockExamRecordDao) GetStats(tx *gorm.DB, userID uint, courseEk int64) (mockCount int, highestScore int, err error) {
	var count int64
	if err = tx.Model(&models.MockExamRecord{}).
		Where("user_id = ? AND course_ek = ? AND deleted_at IS NULL", userID, courseEk).
		Count(&count).Error; err != nil {
		return
	}
	mockCount = int(count)

	var maxScore struct{ MaxScore int }
	if err = tx.Model(&models.MockExamRecord{}).
		Select("COALESCE(MAX(score), 0) as max_score").
		Where("user_id = ? AND course_ek = ? AND deleted_at IS NULL", userID, courseEk).
		Scan(&maxScore).Error; err != nil {
		return
	}
	highestScore = maxScore.MaxScore
	return
}

func (d mockExamRecordDao) GetList(tx *gorm.DB, userID uint, courseEk int64) ([]models.MockExamRecord, error) {
	var records []models.MockExamRecord
	err := tx.Model(&models.MockExamRecord{}).
		Where("user_id = ? AND course_ek = ? AND deleted_at IS NULL", userID, courseEk).
		Order("created_at DESC").
		Find(&records).Error
	return records, err
}

func (d mockExamRecordDao) GetFirstPassDate(tx *gorm.DB, userID uint, courseEk int64, minScore int) (*time.Time, error) {
	var record models.MockExamRecord
	err := tx.Model(&models.MockExamRecord{}).
		Where("user_id = ? AND course_ek = ? AND score >= ? AND deleted_at IS NULL", userID, courseEk, minScore).
		Order("created_at ASC").
		First(&record).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	if record.CreatedAt == nil {
		return nil, nil
	}
	return record.CreatedAt, nil
}

// ===== 错题记录 DAO =====

type wrongQuestionRecordDao struct{}

var WrongQuestionRecordDao = new(wrongQuestionRecordDao)

// Save 保存错题（去重：同一用户同一课程同一题只保留一条）
func (d wrongQuestionRecordDao) Save(tx *gorm.DB, userID uint, courseEk int64, questionIDs []uint) error {
	if len(questionIDs) == 0 {
		return nil
	}
	return tx.Transaction(func(tx *gorm.DB) error {
		for _, qid := range questionIDs {
			var count int64
			tx.Model(&models.WrongQuestionRecord{}).
				Where("user_id = ? AND course_ek = ? AND question_id = ? AND deleted_at IS NULL", userID, courseEk, qid).
				Count(&count)
			if count == 0 {
				if err := tx.Create(&models.WrongQuestionRecord{
					UserID:     userID,
					CourseEk:   courseEk,
					QuestionID: qid,
				}).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
}

// ListCourseStats 按课程聚合：返回该用户所有有错题的课程及各题型数量
func (d wrongQuestionRecordDao) ListCourseStats(tx *gorm.DB, userID uint) ([]struct {
	CourseEk      int64
	QuestionType  int8
	Cnt           int
}, error) {
	var rows []struct {
		CourseEk     int64
		QuestionType int8
		Cnt          int
	}
	err := tx.Table("wrong_question_records wqr").
		Select("wqr.course_ek, q.question_type, COUNT(*) as cnt").
		Joins("JOIN questions q ON q.id = wqr.question_id AND q.deleted_at IS NULL").
		Where("wqr.user_id = ? AND wqr.deleted_at IS NULL", userID).
		Group("wqr.course_ek, q.question_type").
		Scan(&rows).Error
	return rows, err
}

// ListByCourse 获取某课程下所有错题详情
func (d wrongQuestionRecordDao) ListByCourse(tx *gorm.DB, userID uint, courseEk int64) ([]models.Question, error) {
	var questions []models.Question
	err := tx.Table("questions q").
		Joins("JOIN wrong_question_records wqr ON wqr.question_id = q.id AND wqr.user_id = ? AND wqr.course_ek = ? AND wqr.deleted_at IS NULL", userID, courseEk).
		Where("q.deleted_at IS NULL").
		Order("q.question_type ASC, q.id ASC").
		Find(&questions).Error
	return questions, err
}
