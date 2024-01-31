// 操作日志相关的结构体
package entity

import "admin-api/util"

// 操作日志
type SysOperationLog struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`         //ID
	AdminId    uint       `gorm:"column:admin_id;comment:'管理员ID';NOT NULL" json:"adminId"`      // admin id
	Username   string     `gorm:"column:username;comment:'管理员用户名';NOT NULL" json:"username"`    // username
	Method     string     `gorm:"column:method;comment:'请求方式';NOT NULL" json:"method"`          // method
	Ip         string     `gorm:"column:ip;comment:'IP';NOT NULL" json:"ip"`                    // ip
	Url        string     `gorm:"column:url;comment:'url';NOT NULL" json:"url"`                 // url
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"` // create time
}

func (SysOperationLog) TableName() string {
	return "sys_operation_log"
}

// Id参数
type SysOperationLogIdDto struct {
	Id uint `json:"id"` // ID
}

// 批量删除ID参数
type BatchDeleteSysOperationLogIdDto struct {
	Ids []uint  // ID列表
}
