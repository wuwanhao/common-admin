// 用户相关结构体
package entity

import "admin-api/util"

// 用户模型对象
type SysAdmin struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL"`                                   // ID
	PostId     int        `gorm:"column:post_id;comment:'岗位ID'" json:"postId"`                                  // postID
	DeptId     int        `gorm:"column:dept_id;comment:'部门ID'" json:"deptId"`                                  // 部门ID
	Username   string     `gorm:"column:username;varchar(64);comment:'用户账号';NOT NULL" json:"username"`          // 用户账号
	Password   string     `gorm:"column:password;varchar(64);comment:'密码';NOT NULL" json:"password"`            // 密码
	Nickname   string     `gorm:"column:nickname;varchar(64);comment:'昵称'" json:"nickname"`                     // 昵称
	Status     int        `gorm:"column:status;default:1;comment:'帐号启用状态:1-> 启用,2->禁用';NOT NULL" json:"status"` // 帐号启用状态:1->启用,2->禁用
	Icon       string     `gorm:"column:icon;varchar(500);comment:'头像'" json:"icon"`                            // 头像
	Email      string     `gorm:"column:email;varchar(64);comment:'邮箱'" json:"email"`                           // 邮箱
	Phone      string     `gorm:"column:phone;varchar(64);comment:'电话'" json:"phone"`                           // 电话
	Note       string     `gorm:"column:note;varchar(500);comment:'备注'" json:"note"`                            // 备注
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`                 // 创建时间
}

func (SysAdmin) TableName() string {
	return "sys_admin"
}

// 鉴权用户结结构体
type JwtAdmin struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Icon     string `json:"icon"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Note     string `json:"note"`
}

// 登录对象
type LoginDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Image    string `json:"image" validate:"required,min=4,max=6"` // 验证码
	IdKey    string `json:"idKey" validate:"required"`             // UUID
}

// AddSysAdminDto 新增参数
type AddSysAdminDto struct {
	PostId   int    `validate:"required"` // 岗位ID
	RoleId   uint   `validate:"required"` // 角色ID
	DeptId   int    `validate:"required"` // 部门ID
	Username string `validate:"required"` // 用户名
	Password string `validate:"required"` // 密码
	Nickname string `validate:"required"` // 昵称
	Phone    string `validate:"required"` // 手机号
	Email    string `validate:"required"` // 电子邮件
	Note     string // 备注
	Status   int    `validate:"required"` // 状态1->启用  2->禁用
}

// 详情视图
type SysAdminInfo struct {
	Id       uint   `json:"Id"` // ID
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Status   int    `json:"status"`
	PostId   int    `json:"postId"`
	DeptId   int    `json:"deptId"`
	RoleId   uint   `json:"roleId"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Note     string `json:"note"`
}

// 修改参数
type UpdateSysAdminInfo struct {
	Id       uint
	PostId   int
	DeptId   int
	RoleId   uint
	Username string
	Nickname string
	Phone    string
	Email    string
	Note     string
	Status   int // 状态 1->启用 2->禁用
}

// Id
type SysAdminIdDto struct {
	Id uint
}

// 设置状态参数
type UpdateSysAdminStatusDto struct {
	Id     uint
	Status int // 状态 1->启用 2->禁用
}

// 重置密码参数
type ResetSysAdminPasswordDto struct {
	Id       uint
	Password string
}

// 修改个人信息参数
type UpdatePersonalInfoDto struct {
	Id       uint
	Icon     string // 头像
	Username string `validate:"required"`
	Nickname string `validate:"required"`
	Phone    string `validate:"required"`
	Email    string `validate:"required"`
	Note     string `validate:"required"`
}

// 修改个人密码
type UpdatePersonalPasswordDto struct {
	Id              uint   // Id
	Password        string `validate:"required"`
	NewPassword     string `validate:"required"`
	ConfirmPassword string `validate:"required"`
}

// 用户列表的vo视图
type SysAdminVo struct {
	ID         uint       `json:"id"`
	Username   string     `json:"username"`
	Nickname   string     `json:"nickname"`
	Status     int        `json:"status"`
	PostId     int        `json:"postId"`
	DeptId     int        `json:"deptId"`
	RoleId     uint       `json:"roleId" `
	PostName   string     `json:"postName"`
	DeptName   string     `json:"deptName"`
	RoleName   string     `json:"roleName"`
	Icon       string     `json:"icon"`
	Email      string     `json:"email"`
	Phone      string     `json:"phone"`
	Note       string     `json:"Note"`
	CreateTime util.HTime `json:"createTime"`
}
