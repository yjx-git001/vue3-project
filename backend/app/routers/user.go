package routers

import (
	"backend/app/httpapis"
	"backend/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func init() {
	apiRouters = append(apiRouters, registerUserRouter)
}

func registerUserRouter(apiGroup *gin.RouterGroup) {
	userApi := httpapis.UserApi{}
	courseApi := httpapis.CourseApi{}

	user := apiGroup.Group("/api/user")
	{
		user.POST("/register", userApi.Register) //用户注册
		user.POST("/login", userApi.Login)       //用户登录
	}

	auth := apiGroup.Group("/api/user").Use(jwt.AuthMiddleware())
	{
		auth.GET("/info", userApi.GetUserInfo)       //获取用户信息
		auth.PUT("/profile", userApi.UpdateProfile)  //更新用户资料
		auth.GET("/courses", courseApi.GetMyCourses) //获取我的已购课程
	}

	admin := apiGroup.Group("/api/admin/users")
	{
		admin.GET("/options", userApi.GetAdminUserOptions) // 后台用户选项
	}
}
