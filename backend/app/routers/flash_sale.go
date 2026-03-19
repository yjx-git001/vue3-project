package routers

import (
	"backend/app/httpapis"

	"github.com/gin-gonic/gin"
)

func init() {
	apiRouters = append(apiRouters, registerFlashSaleRouter)
}

func registerFlashSaleRouter(apiGroup *gin.RouterGroup) {
	flashSaleApi := httpapis.FlashSaleApi{}

	fs := apiGroup.Group("/api/flash_sales")
	{
		fs.GET("/list", flashSaleApi.GetAll)      //秒杀活动列表
		fs.GET("/active", flashSaleApi.GetActive) //获取当前秒杀活动
		fs.POST("/create", flashSaleApi.Create)   //创建秒杀活动
		fs.PUT("/enable", flashSaleApi.SetEnable) //启用/禁用秒杀活动
		fs.DELETE("/delete", flashSaleApi.Delete) //删除秒杀活动
	}
}
