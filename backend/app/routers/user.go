package routers

import (
	"backend/app/httpapis"

	"github.com/gin-gonic/gin"
)

func init() {
	apiRouters = append(apiRouters, registerUserRouter)
}

func registerUserRouter(apiGroup *gin.RouterGroup) {
	userApi := httpapis.UserApi{}

	user := apiGroup.Group("/api/user")
	{
		user.POST("/register", userApi.Register)
		user.POST("/login", userApi.Login)
		user.GET("/info", userApi.GetUserInfo)
	}
}
