// 岗位相关结构体
package entity

import "admin-api/util"

// 岗位模型
type SysPost struct {
	ID uint `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"` // ID
	PostCode string `gorm:"column:post_code;varchar(64);comment:'岗位编 码';NOT NULL" json:"postCode"` // 岗位编码
	PostName string `gorm:"column:post_name;varchar(50);comment:'岗位名 称';NOT NULL" json:"postName"` // 岗位名称
	PostStatus int `gorm:"column:post_status;default:1;comment:'状态(1->正 常 2->停用)';NOT NULL" json:"postStatus"` // 状态(1->正常 2->停用)
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"` // 创建时间
	Remark string `gorm:"column:remark;varchar(500);comment:'备注'" json:"remark"`
}

func (SysPost) TableName() string {
	return "sys_post"
}

// 修改状态的参数
type UpdateSysPostStatusDto struct {
	Id int // id
	PostStatus int // 状态(1->正常 2->停用)
}

// ID参数
type SysPostIdDto struct {
	Id uint `json:"id"` // ID
}

// 删除岗位参数
type DelSysPostDto struct {
	Ids []uint // Id列表
}

// 岗位下拉列表参数
type SysPostVo struct {
	Id       int    `json:"id"`       // ID
	PostName string `json:"postName"` // 岗位名称
}

