package routers

import (
	"backend/app/httpapis"
	"backend/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func init() {
	apiRouters = append(apiRouters, registerCertificateRouter)
}

func registerCertificateRouter(apiGroup *gin.RouterGroup) {
	certApi := httpapis.CertificateApi{}

	auth := apiGroup.Group("/api/certificates").Use(jwt.AuthMiddleware())
	{
		auth.GET("/:id", certApi.GetDetail)
	}
}
