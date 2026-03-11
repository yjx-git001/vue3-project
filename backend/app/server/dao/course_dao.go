package dao

import (
	"backend/app/models"
	"backend/app/pkg/db_query"
	"backend/app/server/dto"
	"backend/config"
	"context"

	"gorm.io/gorm"
)

type courseDao struct{}

var CourseDao = new(courseDao)

func (d courseDao) GetAll() ([]models.Course, error) {
	var list []models.Course
	err := config.DB.Find(&list).Error
	return list, err
}

func (d courseDao) GetByID(id uint) (models.Course, error) {
	var data models.Course
	err := config.DB.First(&data, id).Error
	return data, err
}

func (d courseDao) GetPage(ctx context.Context, req *dto.CourseGetPageReq) (*db_query.PaginationQ, error) {
	tx := config.DB.Model(&models.Course{}).Order("created_at desc")
	return db_query.GetPage[models.Course](req, tx, models.Course{}, false)
}

func (d courseDao) Save(tx *gorm.DB, data models.Course) error {
	return tx.Save(&data).Error
}
