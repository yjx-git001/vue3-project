package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var apiRouters = make([]func(group *gin.RouterGroup), 0)

func Init(r *gin.RouterGroup) error {
	v := r.Group("/")

	for _, f := range apiRouters {
		f(v)
	}

	return nil
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// 静态文件服务
	r.Static("/uploads", "./uploads")

	Init(r.Group("/"))

	return r
}
