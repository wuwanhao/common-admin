package dao

import (
	"admin-api/api/entity"
	"admin-api/pkg/db"
)

// 新增操作日志
func CreateOpLog(log entity.SysOperationLog) {
	db.Db.Create(&log)
}

// 分页查询操作日志列表
func GetSysOpLogList(Username, BeginTime, EndTime string, PageSize, PageNum int) (opLogList []entity.SysOperationLog, count int64) {
	curDb := db.Db.Table("sys_operation_log")
	if Username != "" {
		curDb = curDb.Where("username = ?", Username)
	}

	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("`create_time` BETWEEN ? AND ?", BeginTime, EndTime)
	}

	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("create_time desc").Find(&opLogList)
	return opLogList, count
}

// 根据id删除操作日志
func DeleteSysOperationLogById(dto entity.SysOperationLogIdDto) {
	db.Db.Delete(&entity.SysOperationLog{}, dto)
}

// 批量删除批量操作日志
func BatchDeleteSysOperationLog(dto entity.BatchDeleteSysOperationLogIdDto) {
	db.Db.Where("id in (?)", dto.Ids).Delete(&entity.SysOperationLog{})
}

// 清空操作日志
func CleanSysOperationLog() {
	db.Db.Exec("truncate table sys_operation_log")
}
