package httpapis

import (
	services "backend/app/server"
	"backend/app/server/dto"
	"backend/pkg/api"
	"backend/pkg/logger"

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

	userID := c.GetUint("userID")
	result, err := services.CourseService.GetPage(&req, userID)
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

	userID := c.GetUint("userID")
	result, err := services.CourseService.GetDetail(&req, userID)
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

func (a CourseApi) GetHotCourses(c *gin.Context) {
	a.MakeContext(c)

	userID := c.GetUint("userID")
	result, err := services.CourseService.GetHotCourses(userID)
	if err != nil {
		logger.Sugar.Errorf("get hot courses error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(result, "ok")
}

func (a CourseApi) GetMyCourses(c *gin.Context) {
	a.MakeContext(c)
	userID := c.GetUint("userID")
	result, err := services.CourseService.GetMyCourses(userID)
	if err != nil {
		logger.Sugar.Errorf("get my courses error: %s", err.Error())
		a.ErrorApi(err)
		return
	}
	a.OK(result, "ok")
}

func (a CourseApi) GetSystemOptions(c *gin.Context) {
	a.MakeContext(c)

	result, err := services.CourseService.GetSystemOptions()
	if err != nil {
		logger.Sugar.Errorf("get system options error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(result, "ok")
}

func (a CourseApi) QueryRule(c *gin.Context) {
	a.MakeContext(c)

	result, err := services.CourseService.QueryRule()
	if err != nil {
		logger.Sugar.Errorf("query rule error: %s", err.Error())
		a.ErrorApi(err)
		return
	}

	a.OK(result, "ok")
}
