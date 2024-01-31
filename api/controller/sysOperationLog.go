package controller

import (
	"admin-api/api/entity"
	"admin-api/api/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 分页获取操作日志列表
// @Summary 分页获取操作日志列表接口
// @Produce json
// @Description 分页获取操作日志列表接口
// @Param PageSize query int false "每页数"
// @Param PageNum query int false "分页数"
// @Param Username query string false "用户名"
// @Param BeginTime query string false "开始时间"
// @Param EndTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/sysOperationLog/list [get]
// @Security ApiKeyAuth
func GetSysOperationLogList(c *gin.Context) {
	Username := c.Query("username")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	service.SysOpLogService().GetSysOpLogList(c, Username, BeginTime, EndTime, PageSize, PageNum)
}

// 根据id删除操作日志
// @Summary 根据id删除操作日志
// @Produce json
// @Description 根据id删除操作日志
// @Param data body entity.SysOperationLogIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysOperationLog/delete [delete]
// @Security ApiKeyAuth
func DeleteSysOperationLogById(c *gin.Context) {
	var dto entity.SysOperationLogIdDto
	_= c.BindJSON(&dto)
	service.SysOpLogService().DeleteSysOpLogById(c, dto)
}

// 批量删除操作日志
// @Summary 批量删除操作日志接口
// @Produce json
// @Description 批量删除操作日志接口
// @Param data body entity.BatchDeleteSysOperationLogIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysOperationLog/batch/delete [delete]
// @Security ApiKeyAuth
func BatchDeleteSysOperationLog(c *gin.Context) {
	var dto entity.BatchDeleteSysOperationLogIdDto
	_ = c.BindJSON(&dto)
	service.SysOpLogService().BatchDeleteSysOpLogByIds(c, dto)
}

// 清空操作日志
// @Summary 清空操作日志接口
// @Produce json
// @Description 清空操作日志接口
// @Success 200 {object} result.Result
// @router /api/sysOperationLog/clean [delete]
// @Security ApiKeyAuth
func CleanSysOperationLog(c *gin.Context) {
	service.SysOpLogService().CleanSysOpLog(c)
}

