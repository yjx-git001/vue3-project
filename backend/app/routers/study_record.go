package routers

import (
	"backend/app/httpapis"
	"backend/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func init() {
	apiRouters = append(apiRouters, registerStudyRecordRouter)
}

func registerStudyRecordRouter(apiGroup *gin.RouterGroup) {
	studyApi := httpapis.StudyRecordApi{}

	auth := apiGroup.Group("/api/study").Use(jwt.AuthMiddleware())
	{
		auth.POST("/record", studyApi.Add)              //添加学习记录
		auth.GET("/stats", studyApi.GetStats)           //获取学习统计数据
		auth.GET("/course_duration", studyApi.GetCourseDuration) //获取某课程学习时长
	}
}
