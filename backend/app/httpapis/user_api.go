package httpapis

import (
	"backend/app/pkg/api"
	"backend/app/pkg/logger"
	"backend/app/server/dto"
	"backend/app/server/services"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserApi struct {
	api.Api
}

func (a UserApi) Login(c *gin.Context) {
	var req dto.LoginReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}

	// TODO: 验证短信验证码
	user, err := services.UserService.GetByPhone(req.Phone)
	if err != nil {
		logger.Sugar.Errorf("UserApi login error: %s", err)
		a.Error(401, err, "用户不存在")
		return
	}

	a.OK(user, "登录成功")
}

func (a UserApi) GetUserInfo(c *gin.Context) {
	a.MakeContext(c)
	// TODO: 从 JWT token 中解析用户 ID
	a.OK(gin.H{}, "ok")
}
