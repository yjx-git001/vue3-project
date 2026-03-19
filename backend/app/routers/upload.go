package routers

import (
	"backend/app/httpapis"

	"github.com/gin-gonic/gin"
)

func init() {
	apiRouters = append(apiRouters, registerUploadRouter)
}

func registerUploadRouter(apiGroup *gin.RouterGroup) {
	uploadApi := httpapis.UploadApi{}

	apiGroup.POST("/api/upload", uploadApi.Upload) //上传文件
}
