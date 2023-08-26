package controller

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 解析和验证身份认证令牌
func ParseTokenMidware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在此处添加解析和验证令牌的逻辑
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		// 解析和验证令牌
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 在此处添加密钥获取逻辑
			// 通常从配置文件、环境变量等获取密钥
			return []byte("your-secret-key"), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// 令牌验证通过，继续处理请求
		c.Next()
	}
}
