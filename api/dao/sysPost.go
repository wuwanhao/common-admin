// 岗位dao层
package dao

import (
	"admin-api/api/entity"
	"admin-api/pkg/db"
	"admin-api/util"
	"time"
)

// GetSysPostByCode 根据编码查询岗位
func GetSysPostByCode(postCode string) (sysPost entity.SysPost) {
	db.Db.Where("post_code = ?", postCode).First(&sysPost)
	return sysPost
}

// GetSysPostByName 根据名称查询
func GetSysPostByName(postName string) (sysPost entity.SysPost) {
	db.Db.Where("post_name = ?", postName).First(&sysPost)
	return sysPost
}

// CreateSysPost 新增岗位
func CreateSysPost(sysPost entity.SysPost) bool {
	// 检查code是否重复
	sysPostByCode := GetSysPostByCode(sysPost.PostCode)
	if sysPostByCode.ID > 0 {
		return false
	}

	// 检查Name是否重复
	sysPostByName := GetSysPostByName(sysPost.PostName)
	if sysPostByName.ID > 0 {
		return false
	}

	// 构建SysPost对象
	addSysPost := entity.SysPost{
		PostCode: sysPost.PostCode,
		PostName: sysPost.PostName,
		PostStatus: sysPost.PostStatus,
		CreateTime: util.HTime{Time: time.Now()},
		Remark: sysPost.Remark,
	}

	// 执行保存
	tx := db.Db.Save(&addSysPost)
	if  tx.RowsAffected > 0 {
		return true
	}

	return false

}

// GetSysPostById 根据ID查询岗位
func GetSysPostById(id int) (sysPost entity.SysPost) {
	db.Db.First(&sysPost, id)
	return sysPost
}

// UpdateSysPost 修改岗位
func UpdateSysPost(post entity.SysPost) (sysPost entity.SysPost) {
	db.Db.First(&sysPost, post.ID)
	sysPost.PostName = post.PostName
	sysPost.PostCode = post.PostCode
	sysPost.PostStatus = post.PostStatus
	if post.Remark != "" {
		sysPost.Remark = post.Remark
	}
	db.Db.Save(&sysPost)
	return sysPost
}

// DeleteSysPostById 删除岗位
func DeleteSysPostById(dto entity.SysPostIdDto) bool {
	db.Db.Delete(&entity.SysPost{}, dto.Id)
	return true
}

// BatchDeleteSysPost 批量删除岗位
func BatchDeleteSysPost(dto entity.DelSysPostDto) bool {
	db.Db.Where("id in (?)", dto.Ids).Delete(&entity.SysPost{})
	return true
}

// UpdateSysPostStatus 修改岗位状态
func UpdateSysPostStatus(dto entity.UpdateSysPostStatusDto) bool {
	var sysPost entity.SysPost
	db.Db.First(&sysPost, dto.Id)
	sysPost.PostStatus = dto.PostStatus
	db.Db.Save(&sysPost)
	return true
}

// GetSysPostList 分页查询岗位列表
func GetSysPostList(pageNum, pageSize int, PostName, PostStatus, BeginTime, EndTime string) (sysPostList []entity.SysPost, count int64) {
	curDb := db.Db.Table("sys_post")
	if PostName != "" {
		curDb = curDb.Where("post_name = ?", PostName)
	}
	if BeginTime != "" && EndTime != ""{
		curDb = curDb.Where("`create_time` BETWEEN ? AND ?", BeginTime, EndTime)
	}
	if PostStatus != "" {
		curDb = curDb.Where("post_status = ?", PostStatus)
	}
	curDb.Count(&count)
	curDb.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("create_time desc").Find(&sysPostList)
	return sysPostList, count
}

// QuerySysPostVoList 岗位下拉列表
func QuerySysPostVoList()(sysPostVo entity.SysPostVo){
	db.Db.Table("sys_post").Select("id, post_name").Scan(&sysPostVo)
	return sysPostVo
}