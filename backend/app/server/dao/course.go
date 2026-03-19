package dao

import (
	"backend/app/models"
	"backend/app/server/dto"
	"backend/pkg/db"
	"backend/pkg/db_query"
	"fmt"

	"gorm.io/gorm"
)

type courseDao struct{}

var CourseDao = new(courseDao)

func (d courseDao) GetAllPage(tx *gorm.DB, req *dto.CourseGetPageReq) (*db_query.PaginationQ, error) {
	if tx == nil {
		tx = db.Db
	}

	// 使用 UNION 合并两种课程
	query := tx.Debug().Table("(" +
		"SELECT id, ek, course_name AS title, course_time, price, original_price, image, course_category, parent_ek, course_tag_class, created_at " +
		"FROM courses WHERE course_category = ? AND deleted_at IS NULL " +
		"UNION ALL " +
		"SELECT id, ek, course_name AS title, course_time, price, original_price, image, course_category, parent_ek, course_tag_class, created_at " +
		"FROM courses WHERE course_category = ? AND deleted_at IS NULL" +
		") AS courses", models.CourseCategorySystem, models.CourseCategorySingle).
		Order("created_at DESC")

	if req.Search != "" {
		query = query.Where("title LIKE ?", "%"+req.Search+"%")
	}

	if req.CourseTagClass != 0 {
		query = query.Where("JSON_CONTAINS(course_tag_class, ?)", fmt.Sprintf("%d", req.CourseTagClass))
	}

	return db_query.GetPage(req, query, dto.CoursePageResp{}, false)
}

func (d courseDao) GetSystemPage(tx *gorm.DB, req *dto.CourseGetPageReq) (*db_query.PaginationQ, error) {
	if tx == nil {
		tx = db.Db
	}
	tx = tx.Debug().Model(&models.Course{}).
		Select(`
			courses.id,
			courses.ek,
			courses.course_name AS title,
			courses.course_time,
			courses.price,
			courses.original_price,
			courses.image,
			courses.course_category,
			courses.parent_ek,
			courses.course_tag_class`).
		Where("courses.course_category = ?", models.CourseCategorySystem).
		Order("courses.created_at DESC")

	if req.Search != "" {
		tx = tx.Where("courses.course_name LIKE ?", "%"+req.Search+"%")
	}

	if req.CourseTagClass != 0 {
		tx = tx.Where("JSON_CONTAINS(courses.course_tag_class, ?)", fmt.Sprintf("%d", req.CourseTagClass))
	}

	return db_query.GetPage(req, tx, dto.CoursePageResp{}, false)
}

func (d courseDao) GetSinglePage(tx *gorm.DB, req *dto.CourseGetPageReq) (*db_query.PaginationQ, error) {
	if tx == nil {
		tx = db.Db
	}
	tx = tx.Debug().Model(&models.Course{}).
		Select(`
			courses.id,
			courses.ek,
			courses.course_name AS title,
			courses.course_time,
			courses.price,
			courses.original_price,
			courses.image,
			courses.course_category,
			courses.parent_ek,
			courses.course_tag_class`).
		Where("courses.course_category = ?", models.CourseCategorySingle).
		Order("courses.created_at DESC")

	if req.Search != "" {
		tx = tx.Where("courses.course_name LIKE ?", "%"+req.Search+"%")
	}

	if req.CourseTagClass != 0 {
		tx = tx.Where("JSON_CONTAINS(courses.course_tag_class, ?)", fmt.Sprintf("%d", req.CourseTagClass))
	}

	return db_query.GetPage(req, tx, dto.CoursePageResp{}, false)
}

func (d courseDao) GetSystemByEk(tx *gorm.DB, ek int64) (models.Course, error) {
	if tx == nil {
		tx = db.Db
	}
	var data models.Course
	err := tx.Debug().
		Preload("Children", "course_category = ?", models.CourseCategorySingle).
		Where("ek = ? AND course_category = ?", ek, models.CourseCategorySystem).
		First(&data).Error
	return data, err
}

func (d courseDao) GetSingleByEk(tx *gorm.DB, ek int64) (models.Course, error) {
	if tx == nil {
		tx = db.Db
	}
	var data models.Course
	err := tx.Debug().
		Where("ek = ? AND course_category = ?", ek, models.CourseCategorySingle).
		First(&data).Error
	return data, err
}

func (d courseDao) CreateSystem(tx *gorm.DB, data *models.Course) error {
	if tx == nil {
		tx = db.Db
	}
	return tx.Debug().Create(data).Error
}

func (d courseDao) CreateSingle(tx *gorm.DB, data *models.Course) error {
	if tx == nil {
		tx = db.Db
	}
	return tx.Debug().Create(data).Error
}

func (d courseDao) GetByEk(tx *gorm.DB, ek int64) (*models.Course, error) {
	if tx == nil {
		tx = db.Db
	}
	var course models.Course
	err := tx.Where("ek = ?", ek).First(&course).Error
	return &course, err
}

func (d courseDao) GetSystemOptions(tx *gorm.DB) ([]dto.CourseSystemOption, error) {
	if tx == nil {
		tx = db.Db
	}
	var result []dto.CourseSystemOption
	err := tx.Debug().Model(&models.Course{}).
		Select("ek, course_name").
		Where("course_category = ?", models.CourseCategorySystem).
		Order("created_at DESC").
		Find(&result).Error
	return result, err
}
