// 岗位Controller层
package controller

import (
	"admin-api/api/entity"
	"admin-api/api/service"
	"github.com/gin-gonic/gin"
)

var sysPost = entity.SysPost{}

// 新增岗位
// @Summary 新增岗位接口
// @Produce json
// @Description 新增岗位接口
// @Param date body entity.SysPost true "data"
// @Success 200 {object} result.Result
// @router /api/post/add [post]
func CreateSysPost(c *gin.Context) {
	_ = c.BindJSON(&sysPost)
	service.SysPostService().CreateSysPost(c, sysPost)
}


