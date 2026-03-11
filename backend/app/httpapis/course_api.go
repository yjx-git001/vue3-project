package httpapis

import (
	"backend/app/pkg/api"
	"backend/app/pkg/logger"
	"backend/app/server/dto"
	"backend/app/server/services"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type CourseApi struct {
	api.Api
}

func (a CourseApi) GetPage(c *gin.Context) {
	var req dto.CourseGetPageReq
	if err := a.MakeContext(c).Bind(&req, binding.Query).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}

	result, err := services.CourseService.GetPage(c.Request.Context(), &req)
	if err != nil {
		logger.Sugar.Errorf("CourseApi get page error: %s", err)
		a.Error(500, err, "获取课程失败")
		return
	}

	a.PageOK(result)
}

func (a CourseApi) GetDetail(c *gin.Context) {
	var req dto.CourseDetailReq
	if err := a.MakeContext(c).Bind(&req, binding.Query).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}

	result, err := services.CourseService.GetDetail(req.ID)
	if err != nil {
		logger.Sugar.Errorf("CourseApi get detail error: %s", err)
		a.Error(404, err, "课程不存在")
		return
	}

	a.OK(result, "ok")
}
