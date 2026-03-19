package services

import (
	"backend/app/models"
	"backend/app/server/dao"
	"backend/app/server/dto"
)

type bannerService struct{}

var BannerService = new(bannerService)

func (s bannerService) GetAll() ([]models.Banner, error) {
	return dao.BannerDao.GetAll(nil)
}

func (s bannerService) Create(req *dto.BannerCreateReq) error {
	data := &models.Banner{
		Image:  req.Image,
		Sort:   req.Sort,
		Enable: true,
	}
	return dao.BannerDao.Create(nil, data)
}

func (s bannerService) Delete(id uint) error {
	return dao.BannerDao.Delete(nil, id)
}
