// 岗位服务层
package service

import (
	"admin-api/api/dao"
	"admin-api/api/entity"
	"admin-api/common/result"
	"github.com/gin-gonic/gin"
)

type ISysPostService interface {
	CreateSysPost(c *gin.Context, sysPost entity.SysPost)
	GetSysPostById(c *gin.Context, Id uint)
	UpdateSysPost(c *gin.Context, sysPost entity.SysPost)
	DeleteSysPostById(c *gin.Context, dto entity.SysPostIdDto)
	BatchDeleteSysPost(c *gin.Context, dto entity.DelSysPostDto)
	UpdateSysPostStatus(c *gin.Context, dto entity.UpdateSysPostStatusDto)
	GetSysPostList(c *gin.Context, PageNum, PageSize int, PostName, PostStatus, BeginTime, EndTime string)
	QuerySysPostVoList(c *gin.Context)
}

type SysPostServiceImpl struct {}


// 岗位下拉列表
func (s SysPostServiceImpl)QuerySysPostVoList(c *gin.Context)  {
	result.Success(c, dao.QuerySysPostVoList())
}

// 分页查询岗位列表
func (s SysPostServiceImpl)GetSysPostList(c *gin.Context, PageNum, PageSize int, PostName, PostStatus, BeginTime, EndTime string) {

	if PageSize < 1{
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	list, count := dao.GetSysPostList(PageNum, PageSize, PostName, PostStatus, BeginTime, EndTime)
	result.Success(c, map[string]interface{}{
		"total": count,
		"pageSize": PageSize,
		"pageNum": PageNum,
		"list": list,
	})

}

// 修改岗位状态
func (s SysPostServiceImpl)UpdateSysPostStatus(c *gin.Context,dto entity.UpdateSysPostStatusDto) {
	dao.UpdateSysPostStatus(dto)
	result.Success(c, true)
}

// 批量删除岗位
func (s SysPostServiceImpl) BatchDeleteSysPost(c *gin.Context, dto entity.DelSysPostDto)  {
	dao.BatchDeleteSysPost(dto)
	result.Success(c, true)
}

// 根据ID删除岗位
func (s SysPostServiceImpl) DeleteSysPostById(c *gin.Context, dto entity.SysPostIdDto)  {
	dao.DeleteSysPostById(dto)
	result.Success(c, true)
}

// 修改岗位
func (s SysPostServiceImpl) UpdateSysPost(c *gin.Context, sysPost entity.SysPost)  {
	result.Success(c, dao.UpdateSysPost(sysPost))
}

// 根据ID查询岗位
func (s SysPostServiceImpl)GetSysPostById(c *gin.Context, id uint)  {
	result.Success(c, dao.GetSysPostById(id))
}

// 新增岗位
func (s SysPostServiceImpl)CreateSysPost(c *gin.Context, sysPost entity.SysPost)  {
	createSysPost := dao.CreateSysPost(sysPost)
	if !createSysPost {
		result.Failed(c, int(result.ApiCode.POSTALREADYEXISTS), result.ApiCode.GetMessage(result.ApiCode.POSTALREADYEXISTS))
		return
	}
	result.Success(c, true)
}

var sysPostService = SysPostServiceImpl{}

func SysPostService() ISysPostService {
	return &sysPostService
}



