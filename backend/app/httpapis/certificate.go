package httpapis

import (
	services "backend/app/server"
	"backend/pkg/api"
	"backend/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CertificateApi struct {
	api.Api
}

func (a CertificateApi) GetDetail(c *gin.Context) {
	a.MakeContext(c)
	userID := c.GetUint("userID")
	idStr := c.Param("id")
	ek, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		a.ErrorApi(err)
		return
	}
	result, err := services.CertificateService.GetDetail(userID, ek)
	if err != nil {
		logger.Sugar.Errorf("get certificate detail error: %s", err.Error())
		a.ErrorApi(err)
		return
	}
	a.OK(result, "ok")
}
