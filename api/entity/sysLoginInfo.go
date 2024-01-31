// 登录日志相关模型
package entity

import "admin-api/util"

// 登录日志
type SysLoginInfo struct {
	ID            uint       `json:"id" gorm:"column:id;comment:'主键';primaryKey;NOT NULL"`
	Username      string     `json:"username" gorm:"column:username;varchar(50);comment:'用户账号'"`
	IpAddress     string     `json:"ipAddress" gorm:"column:ip_address;varchar(128);comment:'登录IP地址'"`
	LoginLocation string     `json:"loginLocation" gorm:"column:login_location;varchar(255);comment:'登录地点'"`
	Browser       string     `json:"browser" gorm:"column:browser;varchar(50);comment:'浏览器类型'"`
	Os            string     `json:"os" gorm:"column:os;varchar(50);comment:'操作系统'"`
	LoginStatus   int        `json:"loginStatus" gorm:"column:login_status;comment:'登录状态(1-成功 2-失败)'"`
	Message       string     `json:"message" gorm:"column:message;varchar(255);comment:'提示消息'"`
	LoginTime     util.HTime `json:"loginTime" gorm:"column:login_time;comment:'访问时间'"`
}

func (SysLoginInfo) TableName() string {
	return "sys_login_info"
}

// Id参数
type SysLoginInfoIdDto struct {
	Id uint `json:"id"` // ID
}

// 批量删除登录日志Id列表
type DelSysLoginInfoDto struct {
	Ids []uint // Id列表
}
