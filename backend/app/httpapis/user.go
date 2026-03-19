package httpapis

import (
	services "backend/app/server"
	"backend/app/server/dto"
	"backend/pkg/api"
	"backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserApi struct {
	api.Api
}

func (a UserApi) Register(c *gin.Context) {
	var req dto.RegisterReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}

	err := services.UserService.Register(&req)
	if err != nil {
		logger.Sugar.Errorf("register error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(nil, "注册成功")
}

func (a UserApi) Login(c *gin.Context) {
	var req dto.LoginReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}

	result, err := services.UserService.Login(&req)
	if err != nil {
		logger.Sugar.Errorf("login error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(result, "登录成功")
}

func (a UserApi) GetUserInfo(c *gin.Context) {
	a.MakeContext(c)

	userID := c.GetUint("userID")
	user, err := services.UserService.GetByID(userID)
	if err != nil {
		logger.Sugar.Errorf("get user info error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(user, "ok")
}

func (a UserApi) UpdateProfile(c *gin.Context) {
	var req dto.UpdateProfileReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}

	userID := c.GetUint("userID")
	if err := services.UserService.UpdateProfile(userID, &req); err != nil {
		logger.Sugar.Errorf("update profile error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(nil, "更新成功")
}
