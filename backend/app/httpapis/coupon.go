package httpapis

import (
	services "backend/app/server"
	"backend/app/server/dto"
	"backend/pkg/api"
	"backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type CouponApi struct {
	api.Api
}

func (a CouponApi) Issue(c *gin.Context) {
	var req dto.IssueCouponReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}

	codes, err := services.CouponService.Issue(&req)
	if err != nil {
		logger.Sugar.Errorf("issue coupons error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(codes, "发放成功")
}

func (a CouponApi) GetAdminList(c *gin.Context) {
	a.MakeContext(c)

	list, err := services.CouponService.GetAdminList()
	if err != nil {
		logger.Sugar.Errorf("get coupons error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(list, "ok")
}

func (a CouponApi) GetMyCoupons(c *gin.Context) {
	a.MakeContext(c)

	userID := c.GetUint("userID")
	list, err := services.CouponService.GetByUserID(userID)
	if err != nil {
		logger.Sugar.Errorf("get my coupons error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(list, "ok")
}
