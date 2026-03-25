package routers

import (
	"backend/app/httpapis"
	"backend/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func init() {
	apiRouters = append(apiRouters, registerOrderRouter)
}

func registerOrderRouter(apiGroup *gin.RouterGroup) {
	orderApi := httpapis.OrderApi{}

	auth := apiGroup.Group("/api/order").Use(jwt.AuthMiddleware())
	{
		auth.POST("/create", orderApi.Create)         //创建订单（直接支付）
		auth.POST("/pending", orderApi.CreatePending) //创建待支付订单
		auth.POST("/pay", orderApi.Pay)               //支付订单
		auth.GET("/my", orderApi.GetMyOrders)         //获取我的订单
		auth.GET("/purchased", orderApi.HasPurchased) //查询是否购买过某课程
	}

	admin := apiGroup.Group("/api/admin/cardkey")
	{
		admin.POST("/generate", orderApi.GenerateCardKeys) //生成卡密
		admin.GET("/list", orderApi.GetCardKeyList)        //获取卡密列表
	}
}
