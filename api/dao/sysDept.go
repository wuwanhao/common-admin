package dao

import (
	"admin-api/api/entity"
	"admin-api/pkg/db"
	"admin-api/util"
	"time"
)

// GetSysDeptByName 根据部门名称查询
func GetSysDeptByName(deptName string) (sysDept entity.SysDept) {
	db.Db.Where("dept_name = ?", deptName).First(&sysDept)
	return
}

// CreateDept 新增部门
func CreateDept(sysDept entity.SysDept) bool {
	// 检查是否有重名的
	deptByName := GetSysDeptByName(sysDept.DeptName)
	if deptByName.ID > 0 {
		return false
	}

	// 一级部门
	if sysDept.DeptType == 1 {
		dept := entity.SysDept{
			DeptName:   sysDept.DeptName,
			DeptType:   sysDept.DeptType,
			DeptStatus: sysDept.DeptStatus,
			ParentId:   0,
			CreateTime: util.HTime{Time: time.Now()},
		}
		db.Db.Create(&dept)
		return true
	} else {
		// 二级部门
		dept := entity.SysDept{
			DeptStatus: sysDept.DeptStatus,
			ParentId:   sysDept.ParentId,
			DeptType:   sysDept.DeptType,
			DeptName:   sysDept.DeptName,
			CreateTime: util.HTime{Time: time.Now()},
		}
		db.Db.Create(&dept)
		return true
	}
	return false
}

// QuerySysDeptVoList 部门下拉列表
func QuerySysDeptVoList() (sysDeptVo []entity.SysDeptVo) {
	db.Db.Table("sys_dept").Select("id, dept_name AS label, parent_id").Scan(&sysDeptVo)
	return sysDeptVo
}

// GetSysDeptById 根据id查询部门
func GetSysDeptById(Id int) (sysDept entity.SysDept) {
	db.Db.First(&sysDept, Id)
	return sysDept
}

// UpdateSysDept 修改部门
func UpdateSysDept(dept entity.SysDept) (sysDept entity.SysDept) {
	db.Db.First(&sysDept, dept.ID)
	sysDept.ParentId = dept.ParentId
	sysDept.DeptType = dept.DeptType
	sysDept.DeptName = dept.DeptName
	sysDept.DeptStatus = dept.DeptStatus
	db.Db.Save(&sysDept)
	return sysDept
}

// 查询部门下是否有成员
func GetSysAdminDept(id int) (sysAdmin entity.SysAdmin) {
	db.Db.Where("dept_id = ?", id).First(&sysAdmin)
	return sysAdmin
}

// 删除部门
func DeleteSysDeptById(dto entity.SysDeptIdDto) bool {
	sysAdmin := GetSysAdminDept(dto.Id)
	if sysAdmin.ID > 0 {
		return false
	}
	db.Db.Where("parent_id = ?", dto.Id).Delete(&entity.SysDept{})
	db.Db.Delete(&entity.SysDept{}, dto.Id)
	return true
}

// 查询部门列表
func GetSysDeptList(DeptName string, DeptStatus string) (sysDept []entity.SysDept) {
	curDb := db.Db.Table("sys_dept")
	if DeptName != "" {
		curDb = curDb.Where("dept_name = ?", DeptName)
	}
	if DeptStatus != "" {
		curDb = curDb.Where("dept_status = ?", DeptStatus)
	}
	curDb.Find(&sysDept)
	return sysDept
}
