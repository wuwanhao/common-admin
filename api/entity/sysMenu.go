package entity

import "admin-api/util"

// sys_menu model
type SysMenu struct {
	ID         uint       `json:"id" gorm:"column:id;comment:'主键';primaryKey;NOT NULL"`            // ID
	ParentId   uint       `json:"parentId" gorm:"column:parent_id;comment:'父菜单id'"`                // 父菜单id
	MenuName   string     `json:"menuName" gorm:"column:menu_name;varchar(100);comment:'菜单名称'"`    // 菜单名称
	Icon       string     `json:"icon" gorm:"column:icon;varchar(100);comment:'菜单图标'"`             // 菜单图标
	Value      string     `json:"value" gorm:"column:value;varchar(100);comment:'权限值'"`            // 权限值
	MenuType   uint       `json:"menuType" gorm:"column:menu_type;comment:'菜单类型:1->目录;2->菜单"`      // 菜单类型:1->目录;2->菜单;3->按钮
	Url        string     `json:"url" gorm:"column:url;varchar(100);comment:'菜单url'"`              // 菜单url
	MenuStatus uint       `json:"menuStatus" gorm:"column:menu_status;comment:'启用状态;1->禁用;2->启用'"` // 启用状态;1->禁用;2->启用
	Sort       uint       `json:"sort" gorm:"column:sort;comment:'排序'"`                            // 排序
	CreateTime util.HTime `json:"createTime" gorm:"column:create_time;comment:'创建时间'"`             // 创建时间
	Children   []SysMenu  `json:"children" gorm:"-"`                                               // 子集
}

func (SysMenu) TableName() string {
	return "sys_menu"
}

// Id参数
type SysMenuIdDto struct {
	Id uint `json:"id"` // ID
}

// SysMenuVo 对象
type SysMenuVo struct {
	Id       uint `json:"id"`
	ParentId uint `json:"parentId"` // 父id Label string `json:"label"` // 名称
}
