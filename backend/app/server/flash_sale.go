package services

import (
	"backend/app/models"
	"backend/app/server/dao"
	"backend/app/server/dto"
	"time"
)

type flashSaleService struct{}

var FlashSaleService = new(flashSaleService)

func (s flashSaleService) GetAll() ([]models.FlashSale, error) {
	return dao.FlashSaleDao.GetAll(nil)
}

func (s flashSaleService) GetActive() (*models.FlashSale, error) {
	return dao.FlashSaleDao.GetActive(nil)
}

func (s flashSaleService) Create(req *dto.FlashSaleCreateReq) error {
	startTime, err := time.ParseInLocation("2006-01-02 15:04:05", req.StartTime, time.Local)
	if err != nil {
		return err
	}
	endTime, err := time.ParseInLocation("2006-01-02 15:04:05", req.EndTime, time.Local)
	if err != nil {
		return err
	}
	data := &models.FlashSale{
		CourseEk:  req.CourseEk,
		Price:     req.Price,
		StartTime: &startTime,
		EndTime:   &endTime,
		Enable:    false,
	}
	return dao.FlashSaleDao.Create(nil, data)
}

func (s flashSaleService) SetEnable(id uint, enable bool) error {
	return dao.FlashSaleDao.SetEnable(nil, id, enable)
}

func (s flashSaleService) Delete(id uint) error {
	return dao.FlashSaleDao.Delete(nil, id)
}
