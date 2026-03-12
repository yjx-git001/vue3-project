package dao

import (
	"backend/app/models"
	"backend/pkg/db"
	"backend/pkg/db_query"
	"backend/app/server/dto"

	"gorm.io/gorm"
)

type courseDao struct{}

var CourseDao = new(courseDao)

func (d courseDao) GetAll(tx *gorm.DB) ([]models.Course, error) {
	if tx == nil {
		tx = db.Db
	}
	var list []models.Course
	err := tx.Debug().Find(&list).Error
	return list, err
}

func (d courseDao) GetByID(tx *gorm.DB, id uint) (models.Course, error) {
	if tx == nil {
		tx = db.Db
	}
	var data models.Course
	err := tx.Debug().First(&data, id).Error
	return data, err
}

func (d courseDao) GetPage(tx *gorm.DB, req *dto.CourseGetPageReq) (*db_query.PaginationQ, error) {
	if tx == nil {
		tx = db.Db
	}
	tx = tx.Debug().Model(&models.Course{}).Order("created_at desc")
	return db_query.GetPage(req, tx, models.Course{}, false)
}

func (d courseDao) Save(tx *gorm.DB, data models.Course) error {
	if tx == nil {
		tx = db.Db
	}
	return tx.Debug().Save(&data).Error
}
