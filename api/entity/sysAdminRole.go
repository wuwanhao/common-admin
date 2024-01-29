package entity

// sysRoleMenu 角色与用户关系模型结构体代码
type SysAdminRole struct {
	AdminId uint `gorm:"column:admin_id;comment:'用户ID';NOT NULL" json:"adminId"`
	RoleId  uint `gorm:"column:role_id;comment:'角色ID';NOT NULL" json:"roleId"`
}

func (SysAdminRole) TableName() string {
	return "sys_admin_role"
}
