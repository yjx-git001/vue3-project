package main

import (
	"backend/app/models"
	"backend/app/routers"
	"backend/config"
	"log"
)

func main() {
	// 初始化数据库
	if err := config.InitDB(); err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 自动迁移建表
	if err := models.Init(); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	// 启动路由
	r := routers.SetupRouter()
	log.Println("服务启动在 http://localhost:8080")
	r.Run(":8080")
}
