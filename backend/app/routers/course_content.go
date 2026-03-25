package routers

import (
	"backend/app/httpapis"
	"backend/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func init() {
	apiRouters = append(apiRouters, registerCourseContentRouter)
}

func registerCourseContentRouter(apiGroup *gin.RouterGroup) {
	api := httpapis.CourseContentApi{}

	// 题目
	questions := apiGroup.Group("/api/questions")
	{
		questions.GET("/list", api.ListQuestions)
		questions.POST("/create", api.CreateQuestion)
		questions.DELETE("/delete", api.DeleteQuestion)
	}

	// 课程内容（视频、附件、笔记、配置）
	content := apiGroup.Group("/api/course_content")
	{
		content.GET("/videos", api.GetVideos)
		content.POST("/videos", api.SaveVideos)
		content.GET("/attachments", api.GetAttachments)
		content.POST("/attachments", api.SaveAttachments)
		content.GET("/notes", api.GetNotes)
		content.POST("/notes", api.SaveNotes)
		content.GET("/mock_exam_config", api.GetMockExamConfig)
		content.POST("/mock_exam_config", api.SaveMockExamConfig)
	}

	// 用户态课程内容（需要登录）
	contentAuth := apiGroup.Group("/api/course_content").Use(jwt.AuthMiddleware())
	{
		contentAuth.POST("/mock_exam_record", api.SaveMockExamRecord)
		contentAuth.GET("/mock_exam_stats", api.GetMockExamStats)
		contentAuth.GET("/mock_exam_history", api.GetMockExamHistory)
		contentAuth.POST("/wrong_questions", api.SaveWrongQuestions)
		contentAuth.GET("/wrong_questions/courses", api.GetWrongQuestionCourseList)
		contentAuth.GET("/wrong_questions/list", api.GetWrongQuestionList)
	}
}
