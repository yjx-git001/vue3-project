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

type hotCourseRow struct {
	ID             uint                                `gorm:"column:id"`
	Ek             int64                               `gorm:"column:ek"`
	Title          string                              `gorm:"column:title"`
	CourseTime     int                                 `gorm:"column:course_time"`
	Price          int64                               `gorm:"column:price"`
	OriginalPrice  int64                               `gorm:"column:original_price"`
	Image          string                              `gorm:"column:image"`
	CourseCategory models.CourseCategory               `gorm:"column:course_category"`
	ParentEk       int64                               `gorm:"column:parent_ek"`
	CourseTagClass models.Array[models.CourseTagClass] `gorm:"column:course_tag_class"`
}

func (d courseDao) GetHotCourses(tx *gorm.DB) ([]dto.CoursePageResp, error) {
	if tx == nil {
		tx = db.Db
	}
	var rows []hotCourseRow
	err := tx.Debug().Table("courses").
		Select("courses.id, courses.ek, courses.course_name AS title, courses.course_time, courses.price, courses.original_price, courses.image, courses.course_category, courses.parent_ek, courses.course_tag_class").
		Where("courses.deleted_at IS NULL").
		Joins("LEFT JOIN (SELECT course_ek, COUNT(*) AS buy_count FROM orders WHERE status = 2 GROUP BY course_ek) AS o ON courses.ek = o.course_ek").
		Order("buy_count DESC").
		Limit(4).
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}
	result := make([]dto.CoursePageResp, len(rows))
	for i, r := range rows {
		result[i] = dto.CoursePageResp{
			ID:             r.ID,
			Ek:             r.Ek,
			Title:          r.Title,
			CourseTime:     r.CourseTime,
			Price:          r.Price,
			OriginalPrice:  r.OriginalPrice,
			Image:          r.Image,
			CourseCategory: r.CourseCategory,
			ParentEk:       r.ParentEk,
			CourseTagClass: r.CourseTagClass,
		}
	}
	return result, nil
}

func (d courseDao) GetMyCourses(tx *gorm.DB, userID uint) ([]dto.MyCourseResp, error) {
	if tx == nil {
		tx = db.Db
	}

	// 1. 取该用户所有已支付订单的 course_ek
	var paidEks []int64
	if err := tx.Model(&models.Order{}).
		Where("user_id = ? AND status = ?", userID, models.OrderStatusPaid).
		Pluck("course_ek", &paidEks).Error; err != nil {
		return nil, err
	}
	if len(paidEks) == 0 {
		return []dto.MyCourseResp{}, nil
	}

	// 2. 查这些课程的基本信息
	var courses []models.Course
	if err := tx.Select("ek, course_name, image, course_time, course_category, parent_ek").
		Where("ek IN ? AND deleted_at IS NULL", paidEks).
		Find(&courses).Error; err != nil {
		return nil, err
	}

	// 3. 收集其中体系课程的 ek，查它们的子单门课
	var systemEks []int64
	for _, c := range courses {
		if c.CourseCategory == models.CourseCategorySystem {
			systemEks = append(systemEks, c.Ek)
		}
	}

	var children []models.Course
	if len(systemEks) > 0 {
		if err := tx.Select("ek, course_name, image, course_time, course_category, parent_ek").
			Where("parent_ek IN ? AND course_category = ? AND deleted_at IS NULL",
				systemEks, models.CourseCategorySingle).
			Find(&children).Error; err != nil {
			return nil, err
		}
	}

	// 4. 合并去重（以 ek 为 key）
	seen := make(map[int64]struct{})
	result := make([]dto.MyCourseResp, 0, len(courses)+len(children))

	addCourse := func(ek int64, name, image string, courseTime int, category models.CourseCategory) {
		if _, ok := seen[ek]; ok {
			return
		}
		seen[ek] = struct{}{}
		result = append(result, dto.MyCourseResp{
			Ek:             ek,
			CourseName:     name,
			Image:          image,
			CourseTime:     courseTime,
			CourseCategory: category,
		})
	}

	for _, c := range courses {
		addCourse(c.Ek, c.CourseName, c.Image, c.CourseTime, c.CourseCategory)
	}
	for _, c := range children {
		addCourse(c.Ek, c.CourseName, c.Image, c.CourseTime, c.CourseCategory)
	}

	return result, nil
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
