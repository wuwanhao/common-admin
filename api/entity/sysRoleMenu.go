package entity

// sysRoleMenu 角色与菜单关系模型结构体代码
type SysRoleMenu struct {
	RoleId uint `gorm:"column:role_id;comment:'角色ID';NOT NULL" json:"roleId"`
	MenuId uint `gorm:"column:menu_id;comment:'菜单ID';NOT NULL" json:"menuId"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
