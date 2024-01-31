// 岗位Controller层
package controller

import (
	"admin-api/api/entity"
	"admin-api/api/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

var sysPost = entity.SysPost{}

// 新增岗位
// @Summary 新增岗位接口
// @Produce json
// @Description 新增岗位接口
// @Param date body entity.SysPost true "data"
// @Success 200 {object} result.Result
// @router /api/post/add [post]
// @Security ApiKeyAuth
func CreateSysPost(c *gin.Context) {
	_ = c.BindJSON(&sysPost)
	service.SysPostService().CreateSysPost(c, sysPost)
}

// 根据id查询岗位
// @Summary 根据id查询岗位
// @Produce json
// @Description 根据id查询岗位
// @Param Id query int true "ID"
// @Success 200 {object} result.Result
// @router /api/post/info [get]
// @Security ApiKeyAuth
func GetSysPostById(c *gin.Context)  {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysPostService().GetSysPostById(c, Id)
}

// 修改岗位
// @Summary 修改岗位接口
// @Produce json
// @Description 修改岗位接口
// @Param data body entity.SysPost true "data"
// @Success 200 {object} result.Result
// @router /api/post/update [put]
// @Security ApiKeyAuth
func UpdateSysPost(c *gin.Context) {
	_ = c.BindJSON(&sysPost)
	service.SysPostService().UpdateSysPost(c, sysPost)
}

// 根据ID删除岗位
// @Security ApiKeyAuth
func DeleteSysPostById(c *gin.Context)  {
	var dto entity.SysPostIdDto
	_ = c.BindJSON(&dto)
	service.SysPostService().DeleteSysPostById(c, dto)
}

// 批量删除岗位
// @Security ApiKeyAuth
func BatchDeleteSysPost(c *gin.Context) {
	var dto entity.DelSysPostDto
	_ = c.BindJSON(&dto)
	service.SysPostService().BatchDeleteSysPost(c, dto)
}

// 岗位状态修改
// @Summary 岗位状态修改接口
// @Produce json
// @Description 岗位状态修改接口
// @Param data body entity.UpdateSysPostStatusDto true "data"
// @Success 200 {object} result.Result
// @router /api/post/updateStatus [put]
// @Security ApiKeyAuth
func UpdateSysPostStatus(c *gin.Context) {
	var dto entity.UpdateSysPostStatusDto
	_ = c.BindJSON(&dto)
	service.SysPostService().UpdateSysPostStatus(c, dto)
}
// 分页查询岗位列表
// @Summary 分页查询岗位列表
// @Produce json
// @Description 分页查询岗位列表
// @Param PageNum query int false "分页数"
// @Param PageSize query int false "每页数"
// @Param PostName query string false "岗位名称"
// @Param PostStatus query string false "状态:1->启用,2->禁用"
// @Param BeginTime query string false "开始时间"
// @Param EndTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/post/list [get]
// @Security ApiKeyAuth
func GetSysPostList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	PostName := c.Query("postName")
	PostStatus := c.Query("postStatus")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysPostService().GetSysPostList(c, PageNum, PageSize, PostName,
		PostStatus, BeginTime, EndTime)
}
// 岗位下拉列表
// @Summary 岗位下拉列表
// @Produce json
// @Description 岗位下拉列表
// @Success 200 {object} result.Result
// @router /api/post/vo/list [get]
// @Security ApiKeyAuth
func QuerySysPostVoList(c *gin.Context) {
	service.SysPostService().QuerySysPostVoList(c)
}


