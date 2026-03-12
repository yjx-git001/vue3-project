package utils

import (
	"fmt"
	"strconv"
	"time"
)

// GetForeignKey 生成业务ID (ek)
func GetForeignKey(id uint) int64 {
	//到分钟
	num, _ := strconv.ParseInt(fmt.Sprintf("%s%04d", time.Now().Format("0601021504"), id%10000), 10, 64)
	return num
}

// GetPlatformCode 生成平台编码前缀
func GetPlatformCode(prefix string) string {
	return fmt.Sprintf("%s%s", prefix, time.Now().Format("0601021504")) //到分钟
}
