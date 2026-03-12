package routers

import (
	"backend/app/httpapis"

	"github.com/gin-gonic/gin"
)

func init() {
	apiRouters = append(apiRouters, registerCourseRouter)
}

func registerCourseRouter(apiGroup *gin.RouterGroup) {
	courseApi := httpapis.CourseApi{}

	course := apiGroup.Group("/api/courses")
	{
		course.GET("/list", courseApi.GetPage)
		course.GET("/detail", courseApi.GetDetail)
	}
}
