package services

import (
	"backend/app/server/dao"
	"backend/app/server/dto"
	"errors"

	"gorm.io/gorm"
)

type userService struct{}

var UserService = new(userService)

func (s userService) GetByPhone(phone string) (dto.UserInfoResp, error) {
	user, err := dao.UserDao.GetByPhone(phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.UserInfoResp{}, errors.New("用户不存在")
		}
		return dto.UserInfoResp{}, err
	}
	return dto.UserInfoResp{
		ID:     user.ID,
		Name:   user.Name,
		Phone:  user.Phone,
		Avatar: user.Avatar,
	}, nil
}
