// 登录日志服务层
package service

import (
	"admin-api/api/dao"
	"admin-api/api/entity"
	"admin-api/common/result"
	"github.com/gin-gonic/gin"
)

// 接口定义
type ISysLoginInfoService interface {
	GetSysLoginInfoList(c *gin.Context, Username, LoginStatus, BeginTime, EndTime string, PageSize, PageNum int)
	BatchDeleteSysLoginInfo(c *gin.Context, dto entity.DelSysLoginInfoDto)
	DeleteSysLoginInfoById(c *gin.Context, dto entity.SysLoginInfoIdDto)
	CleanSysLoginInfo(c *gin.Context)
}

// 接口实现类
type SysLoginInfoImpl struct{}

func (s SysLoginInfoImpl) GetSysLoginInfoList(c *gin.Context, Username, LoginStatus, BeginTime, EndTime string, PageSize, PageNum int) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysLoginInfoList, count := dao.GetSysLoginInfoList(Username, LoginStatus, BeginTime, EndTime, PageSize, PageNum)
	result.Success(c, map[string]interface{}{
		"list":     sysLoginInfoList,
		"total":    count,
		"pageSize": PageSize,
		"pageNum":  PageNum,
	})
}

func (s SysLoginInfoImpl) BatchDeleteSysLoginInfo(c *gin.Context, dto entity.DelSysLoginInfoDto) {
	dao.BatchDeleteSysLoginInfo(dto)
	result.Success(c, true)
}

func (s SysLoginInfoImpl) DeleteSysLoginInfoById(c *gin.Context, dto entity.SysLoginInfoIdDto) {
	dao.DeleteSysLogById(dto)
	result.Success(c, true)
}

func (s SysLoginInfoImpl) CleanSysLoginInfo(c *gin.Context) {
	dao.CleanSysLoginInfo()
	result.Success(c, true)
}

var sysLoginInfoService = SysLoginInfoImpl{}

func SysLoginInfoService() ISysLoginInfoService {
	return &sysLoginInfoService
}
