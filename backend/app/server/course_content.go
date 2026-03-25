package services

import (
	"backend/app/models"
	"backend/app/server/dao"
	"backend/app/server/dto"
	"backend/pkg/db"
)

type courseContentService struct{}

var CourseContentService = new(courseContentService)

// ===== 题目 =====

func (s courseContentService) CreateQuestion(req *dto.QuestionCreateReq) error {
	q := &models.Question{
		CourseEk:     req.CourseEk,
		QuestionType: req.QuestionType,
		Content:      req.Content,
		OptionA:      req.OptionA,
		OptionB:      req.OptionB,
		OptionC:      req.OptionC,
		OptionD:      req.OptionD,
		Answer:       req.Answer,
		Analysis:     req.Analysis,
	}
	return dao.QuestionDao.Create(db.Db, q)
}

func (s courseContentService) ListQuestions(courseEk int64) ([]models.Question, error) {
	return dao.QuestionDao.ListByCourseEk(db.Db, courseEk)
}

func (s courseContentService) DeleteQuestion(id uint) error {
	return dao.QuestionDao.Delete(db.Db, id)
}

// ===== 视频 =====

func (s courseContentService) GetVideos(courseEk int64) ([]models.CourseVideo, error) {
	return dao.CourseVideoDao.ListByCourseEk(db.Db, courseEk)
}

func (s courseContentService) SaveVideos(req *dto.VideoSaveReq) error {
	items := make([]models.CourseVideo, 0, len(req.Videos))
	for _, v := range req.Videos {
		items = append(items, models.CourseVideo{
			CourseEk: req.CourseEk,
			Sort:     v.Sort,
			Title:    v.Title,
			Duration: v.Duration,
		})
	}
	return dao.CourseVideoDao.SaveAll(db.Db, req.CourseEk, items)
}

// ===== 附件 =====

func (s courseContentService) GetAttachments(courseEk int64) ([]models.CourseAttachment, error) {
	return dao.CourseAttachmentDao.ListByCourseEk(db.Db, courseEk)
}

func (s courseContentService) SaveAttachments(req *dto.AttachmentSaveReq) error {
	items := make([]models.CourseAttachment, 0, len(req.Attachments))
	for _, a := range req.Attachments {
		items = append(items, models.CourseAttachment{
			CourseEk: req.CourseEk,
			Sort:     a.Sort,
			Name:     a.Name,
		})
	}
	return dao.CourseAttachmentDao.SaveAll(db.Db, req.CourseEk, items)
}

// ===== 笔记 =====

func (s courseContentService) GetNotes(courseEk int64) (*dto.NotesResp, error) {
	notes, err := dao.CourseNotesDao.GetByCourseEk(db.Db, courseEk)
	if err != nil {
		// 未找到时返回空内容
		return &dto.NotesResp{CourseEk: courseEk, Content: ""}, nil
	}
	return &dto.NotesResp{CourseEk: notes.CourseEk, Content: notes.Content}, nil
}

func (s courseContentService) SaveNotes(req *dto.NotesSaveReq) error {
	return dao.CourseNotesDao.Save(db.Db, req.CourseEk, req.Content)
}

// ===== 模拟考试配置 =====

func (s courseContentService) GetMockExamConfig(courseEk int64) (*dto.MockExamConfigResp, error) {
	cfg, err := dao.MockExamConfigDao.GetByCourseEk(db.Db, courseEk)
	if err != nil {
		// 未配置时返回全0默认值
		return &dto.MockExamConfigResp{CourseEk: courseEk}, nil
	}
	return &dto.MockExamConfigResp{
		CourseEk:      cfg.CourseEk,
		SingleCount:   cfg.SingleCount,
		MultipleCount: cfg.MultipleCount,
		JudgeCount:    cfg.JudgeCount,
	}, nil
}

func (s courseContentService) SaveMockExamConfig(req *dto.MockExamConfigSaveReq) error {
	return dao.MockExamConfigDao.Save(db.Db, req.CourseEk, req.SingleCount, req.MultipleCount, req.JudgeCount)
}

// ===== 模拟考试记录 =====

func (s courseContentService) SaveMockExamRecord(userID uint, req *dto.MockExamRecordSaveReq) error {
	return dao.MockExamRecordDao.Save(db.Db, &models.MockExamRecord{
		UserID:   userID,
		CourseEk: req.CourseEk,
		Score:    req.Score,
		Total:    req.Total,
		Correct:  req.Correct,
		Duration: req.Duration,
	})
}

func (s courseContentService) GetMockExamStats(userID uint, courseEk int64) (*dto.MockExamStatsResp, error) {
	mockCount, highestScore, err := dao.MockExamRecordDao.GetStats(db.Db, userID, courseEk)
	if err != nil {
		return nil, err
	}
	return &dto.MockExamStatsResp{
		MockCount:    mockCount,
		HighestScore: highestScore,
	}, nil
}

func (s courseContentService) GetMockExamHistory(userID uint, courseEk int64) ([]dto.MockExamRecordItem, error) {
	records, err := dao.MockExamRecordDao.GetList(db.Db, userID, courseEk)
	if err != nil {
		return nil, err
	}
	result := make([]dto.MockExamRecordItem, 0, len(records))
	for _, r := range records {
		createdAt := ""
		if r.CreatedAt != nil {
			createdAt = r.CreatedAt.Format("2006-01-02 15:04:05")
		}
		result = append(result, dto.MockExamRecordItem{
			ID:        r.ID,
			Score:     r.Score,
			Total:     r.Total,
			Correct:   r.Correct,
			Duration:  r.Duration,
			CreatedAt: createdAt,
		})
	}
	return result, nil
}

// ===== 错题记录 =====

func (s courseContentService) SaveWrongQuestions(userID uint, req *dto.WrongQuestionSaveReq) error {
	return dao.WrongQuestionRecordDao.Save(db.Db, userID, req.CourseEk, req.QuestionIDs)
}

func (s courseContentService) GetWrongQuestionCourseList(userID uint) ([]dto.WrongQuestionCourseResp, error) {
	rows, err := dao.WrongQuestionRecordDao.ListCourseStats(db.Db, userID)
	if err != nil {
		return nil, err
	}

	// 聚合每个课程的各题型数量
	type stat struct {
		single, multiple, judge int
	}
	statMap := map[int64]*stat{}
	courseEks := []int64{}
	for _, r := range rows {
		if _, ok := statMap[r.CourseEk]; !ok {
			statMap[r.CourseEk] = &stat{}
			courseEks = append(courseEks, r.CourseEk)
		}
		switch r.QuestionType {
		case 1:
			statMap[r.CourseEk].single = r.Cnt
		case 2:
			statMap[r.CourseEk].multiple = r.Cnt
		case 3:
			statMap[r.CourseEk].judge = r.Cnt
		}
	}

	result := make([]dto.WrongQuestionCourseResp, 0, len(courseEks))
	for _, ek := range courseEks {
		course, err := dao.CourseDao.GetByEk(db.Db, ek)
		if err != nil {
			continue
		}
		s := statMap[ek]
		result = append(result, dto.WrongQuestionCourseResp{
			CourseEk:      ek,
			CourseName:    course.CourseName,
			CourseImage:   course.Image,
			SingleCount:   s.single,
			MultipleCount: s.multiple,
			JudgeCount:    s.judge,
		})
	}
	return result, nil
}

func (s courseContentService) GetWrongQuestionList(userID uint, courseEk int64) ([]dto.WrongQuestionItem, error) {
	questions, err := dao.WrongQuestionRecordDao.ListByCourse(db.Db, userID, courseEk)
	if err != nil {
		return nil, err
	}
	items := make([]dto.WrongQuestionItem, 0, len(questions))
	for _, q := range questions {
		items = append(items, dto.WrongQuestionItem{
			ID:           q.ID,
			QuestionType: q.QuestionType,
			Content:      q.Content,
			OptionA:      q.OptionA,
			OptionB:      q.OptionB,
			OptionC:      q.OptionC,
			OptionD:      q.OptionD,
			Answer:       q.Answer,
			Analysis:     q.Analysis,
		})
	}
	return items, nil
}
