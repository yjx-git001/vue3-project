package httpapis

import (
	"backend/pkg/api"
	"backend/pkg/logger"
	"fmt"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UploadApi struct {
	api.Api
}

func (a UploadApi) Upload(c *gin.Context) {
	if err := a.MakeContext(c).Errors; err != nil {
		logger.Sugar.Errorf("make context error: %v", err)
		a.ErrorApi(err)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		logger.Sugar.Errorf("get file error: %s", err)
		a.Error(400, err, "请选择文件")
		return
	}

	// 验证文件大小（限制10MB）
	if file.Size > 10*1024*1024 {
		a.Error(400, nil, "文件大小不能超过10MB")
		return
	}

	// 验证文件类型
	ext := filepath.Ext(file.Filename)
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
		".pdf":  true,
		".doc":  true,
		".docx": true,
		".xls":  true,
		".xlsx": true,
	}
	if !allowedExts[ext] {
		a.Error(400, nil, "不支持的文件类型")
		return
	}

	// 生成唯一文件名
	filename := fmt.Sprintf("%s_%s%s", time.Now().Format("20060102150405"), uuid.New().String()[:8], ext)

	// 保存路径
	savePath := filepath.Join("uploads", filename)

	// 保存文件
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		logger.Sugar.Errorf("save file error: %s", err)
		a.Error(500, err, "文件保存失败")
		return
	}

	// 返回文件访问URL
	fileURL := fmt.Sprintf("/uploads/%s", filename)

	a.OK(gin.H{
		"url":      fileURL,
		"filename": file.Filename,
		"size":     file.Size,
	}, "上传成功")
}
