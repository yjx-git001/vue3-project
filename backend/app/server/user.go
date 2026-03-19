package services

import (
	"backend/app/models"
	"backend/app/server/dao"
	"backend/app/server/dto"
	"backend/pkg/db"
	"backend/pkg/jwt"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct{}

var UserService = new(userService)

func (s userService) Register(req *dto.RegisterReq) error {
	_, err := dao.UserDao.GetByPhone(db.Db, req.Phone)
	if err == nil {
		return errors.New("手机号已被注册")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Password: string(hashedPassword),
	}

	return dao.UserDao.Create(db.Db, user)
}

func (s userService) Login(req *dto.LoginReq) (*dto.LoginResp, error) {
	user, err := dao.UserDao.GetByPhone(db.Db, req.Phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("密码错误")
	}

	token, err := jwt.GenerateToken(user.ID, user.Phone)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResp{
		Token: token,
		User:  toUserInfoResp(user),
	}, nil
}

func (s userService) GetByID(id uint) (dto.UserInfoResp, error) {
	user, err := dao.UserDao.GetByID(nil, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.UserInfoResp{}, errors.New("用户不存在")
		}
		return dto.UserInfoResp{}, err
	}
	return toUserInfoResp(user), nil
}

func (s userService) GetByPhone(phone string) (dto.UserInfoResp, error) {
	user, err := dao.UserDao.GetByPhone(nil, phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.UserInfoResp{}, errors.New("用户不存在")
		}
		return dto.UserInfoResp{}, err
	}
	return toUserInfoResp(user), nil
}

func (s userService) UpdateProfile(userID uint, req *dto.UpdateProfileReq) error {
	return db.Db.Model(&models.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"nickname":     req.Nickname,
		"name":         req.Name,
		"phone":        req.Phone,
		"city":         req.City,
		"organization": req.Organization,
		"company":      req.Company,
		"language":     req.Language,
		"avatar":       req.Avatar,
	}).Error
}

func toUserInfoResp(user models.User) dto.UserInfoResp {
	return dto.UserInfoResp{
		ID:           user.ID,
		Nickname:     user.Nickname,
		Name:         user.Name,
		Phone:        user.Phone,
		Avatar:       user.Avatar,
		City:         user.City,
		Organization: user.Organization,
		Company:      user.Company,
		Language:     user.Language,
		CreatedAt:    user.CreatedAt.Format("2006.01.02"),
	}
}
