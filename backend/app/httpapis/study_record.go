package httpapis

import (
	services "backend/app/server"
	"backend/app/server/dto"
	"backend/pkg/api"
	"backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type StudyRecordApi struct {
	api.Api
}

func (a StudyRecordApi) Add(c *gin.Context) {
	var req dto.AddStudyRecordReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}

	userID := c.GetUint("userID")
	if err := services.StudyRecordService.Add(userID, &req); err != nil {
		logger.Sugar.Errorf("add study record error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(nil, "记录成功")
}

func (a StudyRecordApi) GetStats(c *gin.Context) {
	a.MakeContext(c)

	userID := c.GetUint("userID")
	result, err := services.StudyRecordService.GetStats(userID)
	if err != nil {
		logger.Sugar.Errorf("get study stats error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(result, "ok")
}

func (a StudyRecordApi) GetCourseDuration(c *gin.Context) {
	var req dto.CourseDurationReq
	if err := a.MakeContext(c).Bind(&req, binding.Query).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}
	userID := c.GetUint("userID")
	result, err := services.StudyRecordService.GetCourseDuration(userID, req.CourseEk)
	if err != nil {
		logger.Sugar.Errorf("get course duration error: %s", err.Error())
		a.ErrorApi(err)
		return
	}
	a.OK(result, "ok")
}
