package middleware

import (
	"admin-api/common/result"
	"admin-api/middleware/constant"
	"admin-api/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

// AuthMiddleware 鉴权
func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			result.Failed(c, int(result.ApiCode.NOAUTH),
				result.ApiCode.GetMessage(result.ApiCode.NOAUTH))
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			result.Failed(c, int(result.ApiCode.AUTHFORMATERROR),
				result.ApiCode.GetMessage(result.ApiCode.AUTHFORMATERROR))
			c.Abort()
			return
		}
		// 检验token
		mc, err := jwt.ValidateToken(parts[1])
		if err != nil {
			result.Failed(c, int(result.ApiCode.INVALIDTOKEN),
				result.ApiCode.GetMessage(result.ApiCode.INVALIDTOKEN))
			c.Abort()
			return
		}
		// 存用户信息
		c.Set(constant.ContextKeyUserObj, mc)
		c.Next()
	}
}