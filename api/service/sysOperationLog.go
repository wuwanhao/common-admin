package service

import (
	"admin-api/api/dao"
	"admin-api/api/entity"
	"admin-api/common/result"
	"github.com/gin-gonic/gin"
)

type ISysOpLogService interface {
	GetSysOpLogList(c *gin.Context, Username, BeginTime, EndTime string, PageSize, PageNum int)
	DeleteSysOpLogById(c *gin.Context, dto entity.SysOperationLogIdDto)
	BatchDeleteSysOpLogByIds(c *gin.Context, dto entity.BatchDeleteSysOperationLogIdDto)
	CleanSysOpLog(c *gin.Context)
}

type SysOpLogServiceImpl struct {
}

func (s SysOpLogServiceImpl) GetSysOpLogList(c *gin.Context, Username, BeginTime, EndTime string, PageSize, PageNum int) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysOperationLog, count := dao.GetSysOpLogList(Username, BeginTime, EndTime, PageSize, PageNum)
	result.Success(c, map[string]interface{}{
		"total":    count,
		"pageSize": PageSize,
		"pageNum":  PageNum,
		"list":     sysOperationLog,
	})
}

func (s SysOpLogServiceImpl) DeleteSysOpLogById(c *gin.Context, dto entity.SysOperationLogIdDto) {
	dao.DeleteSysOperationLogById(dto)
	result.Success(c, true)
}

func (s SysOpLogServiceImpl) BatchDeleteSysOpLogByIds(c *gin.Context, dto entity.BatchDeleteSysOperationLogIdDto) {
	dao.BatchDeleteSysOperationLog(dto)
	result.Success(c, true)
}
func (s SysOpLogServiceImpl) CleanSysOpLog(c *gin.Context) {
	dao.CleanSysOperationLog()
	result.Success(c, true)
}

var sysOpLogService = SysOpLogServiceImpl{}

func SysOpLogService() ISysOpLogService {
	return &sysOpLogService
}
