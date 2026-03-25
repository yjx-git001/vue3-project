package httpapis

import (
	"backend/app/models"
	services "backend/app/server"
	"backend/app/server/dto"
	"backend/pkg/api"
	"backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type OrderApi struct {
	api.Api
}

func (a OrderApi) Create(c *gin.Context) {
	var req dto.CreateOrderReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}

	userID := c.GetUint("userID")
	if err := services.OrderService.Create(userID, &req); err != nil {
		logger.Sugar.Errorf("create order error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(nil, "购买成功")
}

func (a OrderApi) GetMyOrders(c *gin.Context) {
	a.MakeContext(c)
	var req dto.GetMyOrdersReq
	if err := c.ShouldBindQuery(&req); err != nil {
		a.Error(400, err, "参数错误")
		return
	}
	if req.Status != nil {
		if *req.Status != models.OrderStatusPending &&
			*req.Status != models.OrderStatusPaid &&
			*req.Status != models.OrderStatusCancelled {
			a.Error(400, nil, "状态参数错误")
			return
		}
	}

	userID := c.GetUint("userID")
	list, err := services.OrderService.GetByUserID(userID, req.Status)
	if err != nil {
		logger.Sugar.Errorf("get orders error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(list, "ok")
}

func (a OrderApi) HasPurchased(c *gin.Context) {
	a.MakeContext(c)

	var req struct {
		CourseEk int64 `form:"courseEk" binding:"required"`
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		a.Error(400, err, "参数错误")
		return
	}

	userID := c.GetUint("userID")
	ok, err := services.OrderService.HasPurchased(userID, req.CourseEk)
	if err != nil {
		logger.Sugar.Errorf("has purchased error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(ok, "ok")
}

func (a OrderApi) CreatePending(c *gin.Context) {
	var req dto.CreatePendingReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}

	userID := c.GetUint("userID")
	order, err := services.OrderService.CreatePending(userID, req.CourseEk)
	if err != nil {
		logger.Sugar.Errorf("create pending order error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(gin.H{"orderNo": order.OrderNo}, "ok")
}

func (a OrderApi) Pay(c *gin.Context) {
	var req dto.PayOrderReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}

	userID := c.GetUint("userID")
	payReq := &dto.CreateOrderReq{PayType: req.PayType, CardCode: req.CardCode}
	if err := services.OrderService.Pay(userID, req.OrderNo, payReq); err != nil {
		logger.Sugar.Errorf("pay order error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(nil, "支付成功")
}

func (a OrderApi) GenerateCardKeys(c *gin.Context) {
	var req dto.GenerateCardKeyReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}

	codes, err := services.OrderService.GenerateCardKeys(&req)
	if err != nil {
		logger.Sugar.Errorf("generate card keys error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(codes, "生成成功")
}

func (a OrderApi) GetCardKeyList(c *gin.Context) {
	a.MakeContext(c)

	list, err := services.OrderService.GetCardKeyList()
	if err != nil {
		logger.Sugar.Errorf("get card key list error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(list, "ok")
}
