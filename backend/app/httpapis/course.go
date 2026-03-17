package httpapis

import (
	"backend/pkg/api"
	"backend/pkg/logger"
	services "backend/app/server"
	"backend/app/server/dto"

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

	result, err := services.CourseService.GetPage(&req)
	if err != nil {
		logger.Sugar.Errorf("get page error: %s", err.Error())
		a.ErrorApi(err)
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

	result, err := services.CourseService.GetDetail(&req)
	if err != nil {
		logger.Sugar.Errorf("get detail error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(result, "ok")
}

func (a CourseApi) CreateSystem(c *gin.Context) {
	var req dto.CourseSystemCreateReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}

	err := services.CourseService.CreateSystem(&req)
	if err != nil {
		logger.Sugar.Errorf("create system course error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(nil, "创建成功")
}

func (a CourseApi) CreateSingle(c *gin.Context) {
	var req dto.CourseSingleCreateReq
	if err := a.MakeContext(c).Bind(&req, binding.JSON).Errors; err != nil {
		logger.Sugar.Errorf("binding req data error: %s", err)
		a.ErrorApi(err)
		return
	}

	err := services.CourseService.CreateSingle(&req)
	if err != nil {
		logger.Sugar.Errorf("create single course error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(nil, "创建成功")
}
