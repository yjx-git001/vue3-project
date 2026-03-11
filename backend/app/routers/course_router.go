package routers

import (
	"backend/app/httpapis"

	"github.com/gin-gonic/gin"
)

func registerCourseRouter(api *gin.RouterGroup) {
	courseApi := httpapis.CourseApi{}
	course := api.Group("/courses")
	{
		course.GET("/list", courseApi.GetPage)
		course.GET("/detail", courseApi.GetDetail)
	}
}
