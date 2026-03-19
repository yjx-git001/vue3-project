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

type FlashSaleApi struct {
	api.Api
}

func (a FlashSaleApi) GetAll(c *gin.Context) {
	a.MakeContext(c)

	result, err := services.FlashSaleService.GetAll()
	if err != nil {
		logger.Sugar.Errorf("get flash sales error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(result, "ok")
}

func (a FlashSaleApi) GetActive(c *gin.Context) {
	a.MakeContext(c)

	result, err := services.FlashSaleService.GetActive()
	if err != nil {
		a.OK(nil, "ok")
		return
	}

	a.OK(result, "ok")
}

func (a FlashSaleApi) Create(c *gin.Context) {
	var req dto.FlashSaleCreateReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}

	if err := services.FlashSaleService.Create(&req); err != nil {
		logger.Sugar.Errorf("create flash sale error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(nil, "创建成功")
}

func (a FlashSaleApi) SetEnable(c *gin.Context) {
	a.MakeContext(c)

	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		a.ErrorApi(err)
		return
	}

	var req dto.FlashSaleSetEnableReq
	if err := c.ShouldBindJSON(&req); err != nil {
		a.ErrorApi(err)
		return
	}

	if err := services.FlashSaleService.SetEnable(uint(id), req.Enable); err != nil {
		logger.Sugar.Errorf("set flash sale enable error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(nil, "操作成功")
}

func (a FlashSaleApi) Delete(c *gin.Context) {
	a.MakeContext(c)

	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		a.ErrorApi(err)
		return
	}

	if err := services.FlashSaleService.Delete(uint(id)); err != nil {
		logger.Sugar.Errorf("delete flash sale error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(nil, "删除成功")
}
