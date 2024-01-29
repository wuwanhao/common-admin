package controller

import (
	"admin-api/api/entity"
	"admin-api/api/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 创建角色
func CreateSysRole(c *gin.Context) {
	var dto entity.AddSysRoleDto
	_ = c.BindJSON(&dto)
	service.SysRoleService().CreateSysRole(c, dto)
}

// / 获取角色详细信息
func GetSysRoleBYId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	service.SysRoleService().GetSysRoleById(c, id)
}

// 修改角色
func UpdateSysRole(c *gin.Context) {
	var dto entity.UpdateSysRoleDto
	_ = c.BindJSON(&dto)
	service.SysRoleService().UpdateSysRole(c, dto)
}

// 根据ID删除角色
func DeleteSysRoleById(c *gin.Context) {
	var dto entity.SysRoleIdDto
	_ = c.BindJSON(&dto)
	service.SysRoleService().DeleteSysRoleById(c, dto)
}

// 更新角色状态
func UpdateSysRoleStatus(c *gin.Context) {
	var dto entity.UpdateSysRoleStatusDto
	_ = c.BindJSON(&dto)
	service.SysRoleService().UpdateSysRoleStatus(c, dto)
}

// 分页查询角色列表
func GetSysRoleList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	RoleName := c.Query("roleName")
	Status := c.Query("status")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysRoleService().GetSysRoleList(c, PageNum, PageSize, RoleName,
		Status, BeginTime, EndTime)
}

// 角色下拉选项列表
func QuerySysRoleVoList(c *gin.Context) {
	service.SysRoleService().QuerySysRoleVoList(c)
}

// 根据角色ID查询菜单数据
func QueryRoleMenuIdList(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("Id"))
	service.SysRoleService().QueryRoleMenuIdList(c, Id)
}

// 分配权限
func AssignPermissions(c *gin.Context) {
	var RoleMenu entity.RoleMenu
	_ = c.BindJSON(&RoleMenu)
	service.SysRoleService().AssignPermissions(c, RoleMenu)
}

