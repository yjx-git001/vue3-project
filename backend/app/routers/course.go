package routers

import (
	"backend/app/httpapis"
	"backend/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func init() {
	apiRouters = append(apiRouters, registerCourseRouter)
}

func registerCourseRouter(apiGroup *gin.RouterGroup) {
	courseApi := httpapis.CourseApi{}

	course := apiGroup.Group("/api/courses")
	{
		course.GET("/list", jwt.OptionalAuthMiddleware(), courseApi.GetPage)     //课程列表
		course.GET("/detail", jwt.OptionalAuthMiddleware(), courseApi.GetDetail) //课程详情
		course.GET("/system_options", courseApi.GetSystemOptions)                //获取系统课程选项
		course.GET("/query_rule", courseApi.QueryRule)                           //查询课程规则
		course.POST("/system", courseApi.CreateSystem)                           //创建系统课程
		course.POST("/single", courseApi.CreateSingle)                           //创建单门课程
	}
}
