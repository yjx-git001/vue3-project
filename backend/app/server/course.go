package services

import (
	"backend/app/models"
	"backend/pkg/db"
	"backend/pkg/db_query"
	"backend/app/server/dao"
	"backend/app/server/dto"
	"errors"
)

type courseService struct{}

var CourseService = new(courseService)

func (s courseService) GetPage(req *dto.CourseGetPageReq) (*db_query.PaginationQ, error) {
	if req.CourseCategory == models.CourseCategorySingle {
		result, err := dao.CourseDao.GetSinglePage(db.Db, req)
		if err != nil {
			return nil, err
		}
		return fillCategoryStr(result), nil
	} else if req.CourseCategory == models.CourseCategorySystem {
		result, err := dao.CourseDao.GetSystemPage(db.Db, req)
		if err != nil {
			return nil, err
		}
		return fillCategoryStr(result), nil
	}

	// 未指定类型，合并两张表结果
	systemResult, err := dao.CourseDao.GetSystemPage(db.Db, req)
	if err != nil {
		return nil, err
	}
	singleResult, err := dao.CourseDao.GetSinglePage(db.Db, req)
	if err != nil {
		return nil, err
	}

	systemData := systemResult.Data.([]dto.CoursePageResp)
	singleData := singleResult.Data.([]dto.CoursePageResp)
	merged := append(systemData, singleData...)
	for i := range merged {
		merged[i].CourseCategoryStr = models.CourseCategoryMap[merged[i].CourseCategory]
	}

	return &db_query.PaginationQ{
		Page:  req.Page,
		Size:  req.PageSize,
		Total: systemResult.Total + singleResult.Total,
		Data:  merged,
	}, nil
}

func fillCategoryStr(result *db_query.PaginationQ) *db_query.PaginationQ {
	list := result.Data.([]dto.CoursePageResp)
	for i := range list {
		list[i].CourseCategoryStr = models.CourseCategoryMap[list[i].CourseCategory]
	}
	result.Data = list
	return result
}

func (s courseService) GetDetail(req *dto.CourseDetailReq) (interface{}, error) {
	if req.CourseCategory == models.CourseCategorySystem {
		return s.getSystemDetail(req.Ek)
	} else if req.CourseCategory == models.CourseCategorySingle {
		return s.getSingleDetail(req.Ek)
	}

	// 不传 courseCategory，先查单门再查体系
	if result, err := s.getSingleDetail(req.Ek); err == nil {
		return result, nil
	}
	return s.getSystemDetail(req.Ek)
}

func (s courseService) getSystemDetail(ek int64) (interface{}, error) {
	course, err := dao.CourseDao.GetSystemByEk(db.Db, ek)
	if err != nil {
		return nil, errors.New("课程不存在")
	}
	return dto.CourseSystemDetailResp{
		ID:                 course.ID,
		Ek:                 course.Ek,
		Code:               course.Code,
		CourseSystemName:   course.CourseSystemName,
		CourseTime:         course.CourseTime,
		Price:              course.Price,
		OriginalPrice:      course.OriginalPrice,
		Image:              course.Image,
		CourseCategory:     course.CourseClass,
		CourseTag:          course.CourseTag,
		CourseDetail:       course.CourseDetail,
		CourseContent:      course.CourseContent.Slice(),
		CourseIntroduction: course.CourseIntroduction,
	}, nil
}

func (s courseService) getSingleDetail(ek int64) (interface{}, error) {
	course, err := dao.CourseDao.GetSingleByEk(db.Db, ek)
	if err != nil {
		return nil, errors.New("课程不存在")
	}
	var theoreticalQuestions models.Num
	if len(course.TheoreticalQuestions) > 0 {
		theoreticalQuestions = course.TheoreticalQuestions[0]
	}
	return dto.CourseSingleDetailResp{
		ID:                   course.ID,
		Ek:                   course.Ek,
		Code:                 course.Code,
		SingleCourseName:     course.SingleCourseName,
		CourseTime:           course.CourseTime,
		Price:                course.Price,
		OriginalPrice:        course.OriginalPrice,
		Image:                course.Image,
		CourseTag:            course.CourseTag,
		CourseDetail:         course.CourseDetail,
		CourseCategory:       course.CourseClass,
		TheoreticalQuestions: theoreticalQuestions,
		VideoQuestions:       course.VideoQuestions,
		Attachment:           course.Attachment,
	}, nil
}

func (s courseService) CreateSystem(req *dto.CourseSystemCreateReq) error {
	course := &models.CourseSystem{
		CourseSystemName:   req.CourseSystemName,
		CourseTime:         req.CourseTime,
		Price:              req.Price,
		OriginalPrice:      req.OriginalPrice,
		Image:              req.Image,
		CourseClass:        req.CourseCategory,
		CourseTag:          req.CourseTag,
		CourseDetail:       req.CourseDetail,
		CourseContent:      models.Array[models.SingleCourseTagList](req.CourseContent),
		CourseIntroduction: req.CourseIntroduction,
	}
	return dao.CourseDao.CreateSystem(db.Db, course)
}

func (s courseService) CreateSingle(req *dto.CourseSingleCreateReq) error {
	course := &models.CourseSingle{
		SingleCourseName:     req.SingleCourseName,
		CourseTime:           req.CourseTime,
		Price:                req.Price,
		OriginalPrice:        req.OriginalPrice,
		Image:                req.Image,
		CourseTag:            req.CourseTag,
		CourseDetail:         req.CourseDetail,
		CourseClass:          req.CourseCategory,
		TheoreticalQuestions: models.Array[models.Num]{req.TheoreticalQuestions},
		VideoQuestions:       req.VideoQuestions,
		Attachment:           req.Attachment,
	}
	return dao.CourseDao.CreateSingle(db.Db, course)
}
