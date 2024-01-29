package dao

import (
	"admin-api/api/entity"
	"admin-api/pkg/db"
	"admin-api/util"
	"time"
)

// 根据名称查询
func GetSysRoleByName(roleName string) (sysRole entity.SysRole) {
	db.Db.Where("role_name = ?", roleName).First(&sysRole)
	return sysRole
}

// 根据角色Key查询
func GetSysRoleByKey(roleKey string) (sysRole entity.SysRole) {
	db.Db.Where("role_key = ?", roleKey).First(&sysRole)
	return sysRole
}

// 新增角色
func CreateSysRole(dto entity.AddSysRoleDto) bool {
	// 检查name是否重复
	sysRoleByName := GetSysRoleByName(dto.RoleName)
	if sysRoleByName.ID > 0 {
		return false
	}
	// 检查key是否重复
	sysRoleByKey := GetSysRoleByKey(dto.RoleKey)
	if sysRoleByKey.ID > 0 {
		return false
	}
	addSysRole := entity.SysRole{
		RoleName:    dto.RoleName,
		RoleKey:     dto.RoleKey,
		Description: dto.Description,
		Status:      dto.Status,
		CreateTime:  util.HTime{Time: time.Now()},
	}
	tx := db.Db.Create(&addSysRole)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// 根据id获取详情
func GetSysRoleById(Id int) (sysRole entity.SysRole) {
	db.Db.First(&sysRole, Id)
	return sysRole
}

// 修改角色
func UpdateSysRole(dto entity.UpdateSysRoleDto) (sysRole entity.SysRole) {
	db.Db.First(&sysRole, dto.Id)
	sysRole.RoleName = dto.RoleName
	sysRole.RoleKey = dto.RoleKey
	sysRole.Status = dto.Status
	if dto.Description != "" {
		sysRole.Description = dto.Description
	}
	db.Db.Save(&sysRole)
	return sysRole
}

// 根据id删除角色
func DeleteSysRoleById(dto entity.SysRoleIdDto) {
	// 删除角色
	db.Db.Table("sys_role").Delete(&entity.SysRole{}, dto.Id)
	// 删除角色和菜单的对应关系
	db.Db.Table("sys_role_menu").Where("role_id = ?",dto.Id).Delete(&entity.SysRoleMenu{})
}

// 角色状态启用/停用
func UpdateSysRoleStatus(dto entity.UpdateSysRoleStatusDto) bool {
	var sysRole entity.SysRole
	db.Db.First(&sysRole, dto.Id)
	sysRole.Status = dto.Status
	tx := db.Db.Save(&sysRole)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// 分页查询角色列表
func GetSysRoleList(PageNum int, PageSize int, RoleName string, status string, BeginTime string, EndTime string) (sysRole []*entity.SysRole, count int64) {
	curDb := db.Db.Table("sys_role")
	if RoleName != "" {
		curDb = curDb.Where("role_name = ?", RoleName)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("create_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	if status != "" {
		curDb = curDb.Where("status = ?", status)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("create_time DESC").Find(&sysRole)
	return sysRole, count
}

// 角色下拉列表
func QuerySysRoleVoList() (sysRoleVo []entity.SysRoleVo) {
	db.Db.Table("sys_role").Select("id, role_name").Scan(&sysRoleVo)
	return sysRoleVo
}

// 根据角色id查询菜单数据
func QueryRoleMenuIdList(Id int) (idVo []entity.IdVo) {
	const menuType int = 3
	db.Db.Table("sys_menu sm").
		Select("sm.id").
		Joins("LEFT JOIN sys_role_menu srm ON srm.menu_id = sm.id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id ").
		Where("sm.menu_type = ?", menuType).
		Where("sr.id = ?", Id).
		Scan(&idVo)
	return idVo
}

// 分配权限
func AssignPermissions(menu entity.RoleMenu) (err error) {
	err = db.Db.Table("sys_role_menu").Where("role_id = ?", menu.Id).Delete(&entity.SysRoleMenu{}).Error
	if err != nil {
		return err
	}
	for _, value := range menu.MenuIds {
		var sysRoleMenu entity.SysRoleMenu
		sysRoleMenu.RoleId = menu.Id
		sysRoleMenu.MenuId = value
		db.Db.Create(&sysRoleMenu)
	}
	return err
}
