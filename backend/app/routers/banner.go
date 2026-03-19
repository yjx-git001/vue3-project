package routers

import (
	"backend/app/httpapis"

	"github.com/gin-gonic/gin"
)

func init() {
	apiRouters = append(apiRouters, registerBannerRouter)
}

func registerBannerRouter(apiGroup *gin.RouterGroup) {
	bannerApi := httpapis.BannerApi{}

	banner := apiGroup.Group("/api/banners")
	{
		banner.GET("/list", bannerApi.GetAll)      //轮播图列表
		banner.POST("/create", bannerApi.Create)   //创建轮播图
		banner.DELETE("/delete", bannerApi.Delete) //删除轮播图
	}
}
