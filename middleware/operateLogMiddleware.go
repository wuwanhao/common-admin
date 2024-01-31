package middleware

import (
	"admin-api/api/dao"
	"admin-api/api/entity"
	"admin-api/pkg/jwt"
	"admin-api/util"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := strings.ToLower(c.Request.Method)
		jwtAdmin, _ := jwt.GetAdmin(c)
		// 只记录非Get请求
		if method != "get" {
			log := entity.SysOperationLog{
				AdminId: jwtAdmin.ID,
				Username: jwtAdmin.Username,
				Ip: c.ClientIP(),
				Method: method,
				Url: c.Request.URL.Path,
				CreateTime: util.HTime{Time: time.Now()},
			}
			dao.CreateOpLog(log)
		}
	}
}
