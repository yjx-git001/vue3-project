package httpapis

import (
	services "backend/app/server"
	"backend/app/server/dto"
	"backend/pkg/api"
	"backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type CourseContentApi struct {
	api.Api
}

// ===== 题目 =====

func (a CourseContentApi) CreateQuestion(c *gin.Context) {
	var req dto.QuestionCreateReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}
	if err := services.CourseContentService.CreateQuestion(&req); err != nil {
		logger.Sugar.Errorf("create question error: %s", err)
		a.ErrorApi(err)
		return
	}
	a.OK(nil, "创建成功")
}

func (a CourseContentApi) ListQuestions(c *gin.Context) {
	var req dto.QuestionListReq
	if err := a.MakeContext(c).Bind(&req, binding.Query).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}
	list, err := services.CourseContentService.ListQuestions(req.CourseEk)
	if err != nil {
		logger.Sugar.Errorf("list questions error: %s", err)
		a.ErrorApi(err)
		return
	}
	a.OK(list, "ok")
}

func (a CourseContentApi) DeleteQuestion(c *gin.Context) {
	var req dto.QuestionDeleteReq
	if err := a.MakeContext(c).Bind(&req, binding.Query).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}
	if err := services.CourseContentService.DeleteQuestion(req.ID); err != nil {
		logger.Sugar.Errorf("delete question error: %s", err)
		a.ErrorApi(err)
		return
	}
	a.OK(nil, "删除成功")
}

// ===== 视频 =====

func (a CourseContentApi) GetVideos(c *gin.Context) {
	var req dto.VideoListReq
	if err := a.MakeContext(c).Bind(&req, binding.Query).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}
	list, err := services.CourseContentService.GetVideos(req.CourseEk)
	if err != nil {
		logger.Sugar.Errorf("get videos error: %s", err)
		a.ErrorApi(err)
		return
	}
	a.OK(list, "ok")
}

func (a CourseContentApi) SaveVideos(c *gin.Context) {
	var req dto.VideoSaveReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}
	if err := services.CourseContentService.SaveVideos(&req); err != nil {
		logger.Sugar.Errorf("save videos error: %s", err)
		a.ErrorApi(err)
		return
	}
	a.OK(nil, "保存成功")
}

// ===== 附件 =====

func (a CourseContentApi) GetAttachments(c *gin.Context) {
	var req dto.AttachmentListReq
	if err := a.MakeContext(c).Bind(&req, binding.Query).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}
	list, err := services.CourseContentService.GetAttachments(req.CourseEk)
	if err != nil {
		logger.Sugar.Errorf("get attachments error: %s", err)
		a.ErrorApi(err)
		return
	}
	a.OK(list, "ok")
}

func (a CourseContentApi) SaveAttachments(c *gin.Context) {
	var req dto.AttachmentSaveReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}
	if err := services.CourseContentService.SaveAttachments(&req); err != nil {
		logger.Sugar.Errorf("save attachments error: %s", err)
		a.ErrorApi(err)
		return
	}
	a.OK(nil, "保存成功")
}

// ===== 笔记 =====

func (a CourseContentApi) GetNotes(c *gin.Context) {
	var req dto.NotesGetReq
	if err := a.MakeContext(c).Bind(&req, binding.Query).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}
	result, err := services.CourseContentService.GetNotes(req.CourseEk)
	if err != nil {
		logger.Sugar.Errorf("get notes error: %s", err)
		a.ErrorApi(err)
		return
	}
	a.OK(result, "ok")
}

func (a CourseContentApi) SaveNotes(c *gin.Context) {
	var req dto.NotesSaveReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}
	if err := services.CourseContentService.SaveNotes(&req); err != nil {
		logger.Sugar.Errorf("save notes error: %s", err)
		a.ErrorApi(err)
		return
	}
	a.OK(nil, "保存成功")
}

// ===== 模拟考试配置 =====

func (a CourseContentApi) GetMockExamConfig(c *gin.Context) {
	var req dto.MockExamConfigGetReq
	if err := a.MakeContext(c).Bind(&req, binding.Query).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}
	result, err := services.CourseContentService.GetMockExamConfig(req.CourseEk)
	if err != nil {
		logger.Sugar.Errorf("get mock exam config error: %s", err)
		a.ErrorApi(err)
		return
	}
	a.OK(result, "ok")
}

func (a CourseContentApi) SaveMockExamConfig(c *gin.Context) {
	var req dto.MockExamConfigSaveReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}
	if err := services.CourseContentService.SaveMockExamConfig(&req); err != nil {
		logger.Sugar.Errorf("save mock exam config error: %s", err)
		a.ErrorApi(err)
		return
	}
	a.OK(nil, "保存成功")
}

// ===== 模拟考试记录 =====

func (a CourseContentApi) SaveMockExamRecord(c *gin.Context) {
	var req dto.MockExamRecordSaveReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}
	userID := c.GetUint("userID")
	if err := services.CourseContentService.SaveMockExamRecord(userID, &req); err != nil {
		logger.Sugar.Errorf("save mock exam record error: %s", err)
		a.ErrorApi(err)
		return
	}
	a.OK(nil, "保存成功")
}

func (a CourseContentApi) GetMockExamStats(c *gin.Context) {
	var req dto.MockExamRecordGetReq
	if err := a.MakeContext(c).Bind(&req, binding.Query).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}
	userID := c.GetUint("userID")
	result, err := services.CourseContentService.GetMockExamStats(userID, req.CourseEk)
	if err != nil {
		logger.Sugar.Errorf("get mock exam stats error: %s", err)
		a.ErrorApi(err)
		return
	}
	a.OK(result, "ok")
}

// ===== 错题记录 =====

func (a CourseContentApi) SaveWrongQuestions(c *gin.Context) {
	var req dto.WrongQuestionSaveReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}
	userID := c.GetUint("userID")
	if err := services.CourseContentService.SaveWrongQuestions(userID, &req); err != nil {
		logger.Sugar.Errorf("save wrong questions error: %s", err)
		a.ErrorApi(err)
		return
	}
	a.OK(nil, "保存成功")
}

func (a CourseContentApi) GetWrongQuestionCourseList(c *gin.Context) {
	a.MakeContext(c)
	userID := c.GetUint("userID")
	result, err := services.CourseContentService.GetWrongQuestionCourseList(userID)
	if err != nil {
		logger.Sugar.Errorf("get wrong question course list error: %s", err)
		a.ErrorApi(err)
		return
	}
	a.OK(result, "ok")
}

func (a CourseContentApi) GetWrongQuestionList(c *gin.Context) {
	var req dto.WrongQuestionListReq
	if err := a.MakeContext(c).Bind(&req, binding.Query).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}
	userID := c.GetUint("userID")
	result, err := services.CourseContentService.GetWrongQuestionList(userID, req.CourseEk)
	if err != nil {
		logger.Sugar.Errorf("get wrong question list error: %s", err)
		a.ErrorApi(err)
		return
	}
	a.OK(result, "ok")
}
