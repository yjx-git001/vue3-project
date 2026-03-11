package routers

import (
	"backend/app/httpapis"

	"github.com/gin-gonic/gin"
)

func registerUserRouter(api *gin.RouterGroup) {
	userApi := httpapis.UserApi{}
	user := api.Group("/user")
	{
		user.POST("/login", userApi.Login)
		user.GET("/info", userApi.GetUserInfo)
	}
}
