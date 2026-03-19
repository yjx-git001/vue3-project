package jwt

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"code": 401, "msg": "未登录"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"code": 401, "msg": "token 格式错误"})
			c.Abort()
			return
		}

		claims, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(401, gin.H{"code": 401, "msg": "token 无效或已过期"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("phone", claims.Phone)
		c.Next()
	}
}

// OptionalAuthMiddleware 可选鉴权，有 token 就解析 userID，没有就放行
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) == 2 && parts[0] == "Bearer" {
				if claims, err := ParseToken(parts[1]); err == nil {
					c.Set("userID", claims.UserID)
					c.Set("phone", claims.Phone)
				}
			}
		}
		c.Next()
	}
}

