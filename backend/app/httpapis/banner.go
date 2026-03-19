package httpapis

import (
	"backend/app/server/dto"
	"backend/pkg/api"
	"backend/pkg/logger"
	services "backend/app/server"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type BannerApi struct {
	api.Api
}

func (a BannerApi) GetAll(c *gin.Context) {
	a.MakeContext(c)
	result, err := services.BannerService.GetAll()
	if err != nil {
		logger.Sugar.Errorf("get banners error: %s", err.Error())
		a.ErrorApi(err)
		return
	}
	a.OK(result, "ok")
}

func (a BannerApi) Create(c *gin.Context) {
	var req dto.BannerCreateReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}
	if err := services.BannerService.Create(&req); err != nil {
		logger.Sugar.Errorf("create banner error: %s", err.Error())
		a.ErrorApi(err)
		return
	}
	a.OK(nil, "创建成功")
}

func (a BannerApi) Delete(c *gin.Context) {
	a.MakeContext(c)
	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		a.ErrorApi(err)
		return
	}
	if err := services.BannerService.Delete(uint(id)); err != nil {
		logger.Sugar.Errorf("delete banner error: %s", err.Error())
		a.ErrorApi(err)
		return
	}
	a.OK(nil, "删除成功")
}
