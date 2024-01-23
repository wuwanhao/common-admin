package controller

import (
	"admin-api/api/service"
	"admin-api/common/result"
	"github.com/gin-gonic/gin"
)

// Captcha 生成验证码接口
// @Summary 验证码接口
// @Produce json
// @Description 验证码接口
// @Success 200 {object} result.Result
// @router /api/captcha [get]
func Captcha(c *gin.Context) {
	id, b64s := service.CaptMake()
	result.Success(c, map[string]interface{}{
		"idKey": id,
		"image": b64s,
	})
}
