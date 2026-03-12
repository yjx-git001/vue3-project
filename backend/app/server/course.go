package services

import (
	"backend/app/models"
	"backend/pkg/db"
	"backend/pkg/db_query"
	"backend/pkg/logger"
	"backend/app/server/dao"
	"backend/app/server/dto"
	"errors"
)

type courseService struct{}

var CourseService = new(courseService)

func (s courseService) GetPage(req *dto.CourseGetPageReq) (*db_query.PaginationQ, error) {
	result, err := dao.CourseDao.GetPage(db.Db, req)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	for i, c := range result.Data.([]models.Course) {
		result.Data.([]models.Course)[i].Status = c.Status
	}

	// 转换为 DTO
	var respList []dto.CoursePageResp
	for _, c := range result.Data.([]models.Course) {
		respList = append(respList, dto.CoursePageResp{
			ID:            c.ID,
			Title:         c.Title,
			Description:   c.Description,
			Price:         c.Price,
			OriginalPrice: c.OriginalPrice,
			Image:         c.Image,
			Tag:           c.Tag,
			Category:      c.Category,
			Status:        c.Status,
			StatusName:    models.CourseStatusMap[c.Status],
		})
	}
	result.Data = respList

	return result, err
}

func (s courseService) GetDetail(id uint) (dto.CourseDetailResp, error) {
	course, err := dao.CourseDao.GetByID(db.Db, id)
	if err != nil {
		return dto.CourseDetailResp{}, errors.New("课程不存在")
	}
	return dto.CourseDetailResp{
		CoursePageResp: dto.CoursePageResp{
			ID:            course.ID,
			Title:         course.Title,
			Description:   course.Description,
			Price:         course.Price,
			OriginalPrice: course.OriginalPrice,
			Image:         course.Image,
			Tag:           course.Tag,
			Category:      course.Category,
			Status:        course.Status,
			StatusName:    models.CourseStatusMap[course.Status],
		},
	}, nil
}
