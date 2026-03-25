package routers

import (
	"backend/app/httpapis"
	"backend/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func init() {
	apiRouters = append(apiRouters, registerCouponRouter)
}

func registerCouponRouter(apiGroup *gin.RouterGroup) {
	couponApi := httpapis.CouponApi{}

	admin := apiGroup.Group("/api/admin/coupons")
	{
		admin.POST("/issue", couponApi.Issue)      // 发放卡券
		admin.GET("/list", couponApi.GetAdminList) // 后台卡券列表
	}

	auth := apiGroup.Group("/api/coupons").Use(jwt.AuthMiddleware())
	{
		auth.GET("", couponApi.GetMyCoupons) // 我的卡券
	}
}
