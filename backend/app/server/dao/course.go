package dao

import (
	"backend/app/models"
	"backend/app/server/dto"
	"backend/pkg/db"
	"backend/pkg/db_query"

	"gorm.io/gorm"
)

type courseDao struct{}

var CourseDao = new(courseDao)

func (d courseDao) GetSystemPage(tx *gorm.DB, req *dto.CourseGetPageReq) (*db_query.PaginationQ, error) {
	if tx == nil {
		tx = db.Db
	}
	tx = tx.Debug().Model(&models.CourseSystem{}).
		Select(`
			courses_system.id,
			courses_system.ek,
			courses_system.course_system_name AS title,
			courses_system.course_time,
			courses_system.price,
			courses_system.original_price,
			courses_system.image,
			courses_system.course_class AS course_category,
			courses_system.course_tag`).
		Order("courses_system.created_at DESC")

	if req.Search != "" {
		tx = tx.Where("courses_system.course_system_name LIKE ?", "%"+req.Search+"%")
	}

	return db_query.GetPage(req, tx, dto.CoursePageResp{}, false)
}

func (d courseDao) GetSinglePage(tx *gorm.DB, req *dto.CourseGetPageReq) (*db_query.PaginationQ, error) {
	if tx == nil {
		tx = db.Db
	}
	tx = tx.Debug().Model(&models.CourseSingle{}).
		Select(`
			courses_single.id,
			courses_single.ek,
			courses_single.single_course_name AS title,
			courses_single.course_time,
			courses_single.price,
			courses_single.original_price,
			courses_single.image,
			courses_single.course_class AS course_category,
			courses_single.course_tag`).
		Order("courses_single.created_at DESC")

	if req.Search != "" {
		tx = tx.Where("courses_single.single_course_name LIKE ?", "%"+req.Search+"%")
	}

	return db_query.GetPage(req, tx, dto.CoursePageResp{}, false)
}

func (d courseDao) GetSystemByEk(tx *gorm.DB, ek int64) (models.CourseSystem, error) {
	if tx == nil {
		tx = db.Db
	}
	var data models.CourseSystem
	err := tx.Debug().Where("ek = ?", ek).First(&data).Error
	return data, err
}

func (d courseDao) GetSingleByEk(tx *gorm.DB, ek int64) (models.CourseSingle, error) {
	if tx == nil {
		tx = db.Db
	}
	var data models.CourseSingle
	err := tx.Debug().Where("ek = ?", ek).First(&data).Error
	return data, err
}

func (d courseDao) CreateSystem(tx *gorm.DB, data *models.CourseSystem) error {
	if tx == nil {
		tx = db.Db
	}
	return tx.Debug().Create(data).Error
}

func (d courseDao) CreateSingle(tx *gorm.DB, data *models.CourseSingle) error {
	if tx == nil {
		tx = db.Db
	}
	return tx.Debug().Create(data).Error
}
