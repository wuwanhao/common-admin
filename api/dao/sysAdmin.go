// 用户dao层
package dao

import (
	"admin-api/api/entity"
	"admin-api/pkg/db"
	"admin-api/util"
	"time"
)

// SysAdminDetail 用户详情
func SysAdminDetail(dto entity.LoginDto) (sysAdmin entity.SysAdmin) {
	username := dto.Username
	db.Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}

// QuerySysAdminDetailByUsername 根据用户名查询用户
func QuerySysAdminDetailByUsername(username string) (sysAdmin entity.SysAdmin) {
	db.Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}

// CreateSysAdmin 新增用户
func CreateSysAdmin(dto entity.AddSysAdminDto) bool {
	sysAdminByUsername := QuerySysAdminDetailByUsername(dto.Username)
	if sysAdminByUsername.ID > 0 {
		return false
	}

	sysAdmin := entity.SysAdmin{
		Username:   dto.Username,
		Password:   util.EncryptionMD5(dto.Password),
		Nickname:   dto.Nickname,
		PostId:     dto.PostId,
		DeptId:     dto.DeptId,
		Phone:      dto.Phone,
		Email:      dto.Email,
		Note:       dto.Note,
		Status:     dto.Status,
		CreateTime: util.HTime{Time: time.Now()},
	}

	tx := db.Db.Create(&sysAdmin)
	sysAdminExist := QuerySysAdminDetailByUsername(dto.Username)
	// 创建admin和Role对应关系
	var sysAdminRole entity.SysAdminRole
	sysAdminRole.AdminId = sysAdminExist.ID
	sysAdminRole.RoleId = dto.RoleId
	db.Db.Create(&sysAdminRole)

	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// QuerySysUserDetailById 根据ID查询用户详情
func QuerySysUserDetailById(id int) (sysAdminInfo entity.SysAdminInfo) {
	// 关联查询
	db.Db.Table("sys_admin").
		Select("sys_admin.*, sys_admin_role.role_id").
		Joins("LEFT JOIN sys_admin_role on sys_admin.id = sys_admin_role.admin_id").
		Joins("LEFT JOIN sys_role on sys_admin_role.role_id = sys_role.id").
		First(&sysAdminInfo, id)
	return sysAdminInfo

}

// UpdateSysAdmin 修改用户
func UpdateSysAdmin(dto entity.UpdateSysAdminInfo) (sysAdmin entity.SysAdmin) {

	// 先根据ID找到
	db.Db.First(&sysAdmin, dto.Id)
	if dto.Username != "" {
		sysAdmin.Username = dto.Username
	}
	sysAdmin.PostId = dto.PostId
	sysAdmin.DeptId = dto.DeptId
	sysAdmin.Status = dto.Status
	if dto.Nickname != "" {
		sysAdmin.Nickname = dto.Nickname
	}
	if dto.Phone != "" {
		sysAdmin.Phone = dto.Phone
	}
	if dto.Email != "" {
		sysAdmin.Email = dto.Email
	}
	if dto.Note != "" {
		sysAdmin.Note = dto.Note
	}
	db.Db.Save(&sysAdmin)
	// 先删除之前的用户角色对应关系
	var sysAdminRole entity.SysAdminRole
	db.Db.Where("admin_id = ?", dto.Id).Delete(entity.SysAdminRole{})
	// 在新增最新的用户角色对应关系
	sysAdminRole.AdminId = dto.Id
	sysAdminRole.RoleId = dto.RoleId
	db.Db.Create(&sysAdminRole)

	return sysAdmin
}

// DeleteSysAdminById 根据ID删除用户
func DeleteSysAdminById(dto entity.SysAdminIdDto) {
	db.Db.First(&entity.SysAdmin{}, dto.Id)
	db.Db.Delete(&entity.SysAdmin{}, dto.Id)
	db.Db.Where("admin_id = ?", dto.Id).Delete(&entity.SysAdminRole{})
}

// UpdateSysAdminStatus 修改用户状态
func UpdateSysAdminStatus(dto entity.UpdateSysAdminStatusDto) {
	var sysAdmin entity.SysAdmin
	db.Db.First(&sysAdmin, dto.Id)
	sysAdmin.Status = dto.Status
	db.Db.Save(&sysAdmin)
}

// ResetSysAdminPassword 重置密码
func ResetSysAdminPassword(dto entity.ResetSysAdminPasswordDto) {
	var sysAdmin entity.SysAdmin
	db.Db.First(&sysAdmin, dto.Id)
	sysAdmin.Password = util.EncryptionMD5(dto.Password)
	db.Db.Save(&sysAdmin)
}

// GetSysAdminList 分页查询用户列表
func GetSysAdminList(PageSize, PageNum int, Username, Status, BeginTime, EndTime string) (sysAdminVo []entity.SysAdminVo, count int64) {
	curDb := db.Db.Table("sys_admin").
		Select("sys_admin.*, sys_post.post_name, sys_role.role_name, sys_dept.dept_name").
		Joins("LEFT JOIN sys_post ON sys_admin.post_id = sys_post.id").
		Joins("LEFT JOIN sys_admin_role ON sys_admin.id = sys_admin_role.admin_id").
		Joins("LEFT JOIN sys_role ON sys_role.id = sys_admin_role.role_id").
		Joins("LEFT JOIN sys_dept ON sys_admin.dept_id = sys_dept.id")

	if Username != "" {
		curDb = curDb.Where("sys_admin.user_name = ?", Username)
	}
	if Status != "" {
		curDb = curDb.Where("sys_admin.status = ?", Status)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("sys_admin.create_time BETWEEN ? AND ?", BeginTime, EndTime)
	}

	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("sys_admin.create_time DESC").Find(&sysAdminVo)
	return sysAdminVo, count

}

// UpdatePersonal 修改个人信息
func UpdatePersonal(dto entity.UpdatePersonalInfoDto) (sysAdmin entity.SysAdmin) {
	db.Db.First(&sysAdmin, dto.Id)
	if dto.Icon != "" {
		sysAdmin.Icon = dto.Icon
	}
	if dto.Username != "" {
		sysAdmin.Username = dto.Username
	}
	if dto.Nickname != "" {
		sysAdmin.Nickname = dto.Nickname
	}
	if dto.Phone != "" {
		sysAdmin.Phone = dto.Phone
	}
	if dto.Email != "" {
		sysAdmin.Email = dto.Email
	}
	db.Db.Save(&sysAdmin)
	return sysAdmin
}

// UpdatePersonalPassword 修改个人密码
func UpdatePersonalPassword(dto entity.UpdatePersonalPasswordDto) (sysAdmin entity.SysAdmin) {
	db.Db.First(&sysAdmin, dto.Id)
	sysAdmin.Password = dto.NewPassword
	db.Db.Save(&sysAdmin)
	return sysAdmin
}