package services

import (
	"backend/app/models"
	"backend/pkg/db"
	"backend/pkg/jwt"
	"backend/app/server/dao"
	"backend/app/server/dto"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct{}

var UserService = new(userService)

func (s userService) Register(req *dto.RegisterReq) error {
	// 检查手机号是否已存在
	_, err := dao.UserDao.GetByPhone(db.Db, req.Phone)
	if err == nil {
		return errors.New("手机号已被注册")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 创建用户
	user := &models.User{
		Phone:    req.Phone,
		Password: string(hashedPassword),
	}

	return dao.UserDao.Create(db.Db, user)
}

func (s userService) Login(req *dto.LoginReq) (*dto.LoginResp, error) {
	// 查询用户
	user, err := dao.UserDao.GetByPhone(db.Db, req.Phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("密码错误")
	}

	// 生成 token
	token, err := jwt.GenerateToken(user.ID, user.Phone)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResp{
		Token: token,
		User: dto.UserInfoResp{
			ID:     user.ID,
			Name:   user.Name,
			Phone:  user.Phone,
			Avatar: user.Avatar,
		},
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
	return dto.UserInfoResp{
		ID:     user.ID,
		Name:   user.Name,
		Phone:  user.Phone,
		Avatar: user.Avatar,
	}, nil
}

func (s userService) GetByPhone(phone string) (dto.UserInfoResp, error) {
	user, err := dao.UserDao.GetByPhone(nil, phone)
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
