// sysAdmin service层
package service

import (
	"admin-api/api/dao"
	"admin-api/api/entity"
	"admin-api/common/result"
	"admin-api/middleware/constant"
	"admin-api/pkg/jwt"
	"admin-api/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ISysAdminService 定义接口
type ISysAdminService interface {
	Login(c *gin.Context, dto entity.LoginDto)
	CreateSysAdmin(c *gin.Context, dto entity.AddSysAdminDto)
	GetSysAdminInfo(c *gin.Context, Id int)
	UpdateSysAdmin(c *gin.Context, info entity.UpdateSysAdminInfo)
	DeleteSysAdminById(c *gin.Context, dto entity.SysAdminIdDto)
	UpdateSysAdminStatus(c *gin.Context, dto entity.UpdateSysAdminStatusDto)
	ResetSysAdminPassword(c *gin.Context, dto entity.ResetSysAdminPasswordDto)
	GetSysAdminList(c *gin.Context, PageSize, PageNum int, Username, Status, BeginTime, EndTime string)
	UpdatePersonal(c *gin.Context, dto entity.UpdatePersonalInfoDto)
	UpdatePersonalPassword(c *gin.Context, dto entity.UpdatePersonalPasswordDto)
}



// 接口实现
type SysAdminServiceImpl struct {
}

// UpdatePersonalPassword 修改个人密码
func (s SysAdminServiceImpl)UpdatePersonalPassword(c *gin.Context, dto entity.UpdatePersonalPasswordDto)  {

	// 参数完整性校验
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissingChangePasswordParameter),
			result.ApiCode.GetMessage(result.ApiCode.MissingChangePasswordParameter))
		return
	}

	// 检查旧密码是否输入正确
	admin, _ := jwt.GetAdmin(c)
	dto.Id = admin.ID
	sysAdminExist := dao.QuerySysAdminDetailByUsername(admin.Username)
	if sysAdminExist.Password != util.EncryptionMD5(dto.Password) {
		result.Failed(c, int(result.ApiCode.PASSWORDNOTTRUE),
			result.ApiCode.GetMessage(result.ApiCode.PASSWORDNOTTRUE))
		return
	}

	// 检验新密码两次输入是否正确
	if dto.NewPassword != dto.ConfirmPassword {
		result.Failed(c, int(result.ApiCode.RESETPASSWORD),
			result.ApiCode.GetMessage(result.ApiCode.RESETPASSWORD))
		return
	}

	dto.NewPassword = util.EncryptionMD5(dto.NewPassword)
	sysAdminUpdatedPassword := dao.UpdatePersonalPassword(dto)
	tokenString, _ := jwt.GenerateTokenByAdmin(sysAdminUpdatedPassword)
	result.Success(c, map[string]interface{}{
		"token": tokenString,
		"sysAdmin": sysAdminUpdatedPassword,
	})

	return
}

// Login 登录
func (s SysAdminServiceImpl) Login(c *gin.Context, dto entity.LoginDto) {
	// 登录参数校验
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissingLoginParameter), result.ApiCode.GetMessage(result.ApiCode.MissingLoginParameter))
		return
	}

	// 验证码是否过期
	code := util.RedisStore{}.Get(dto.IdKey, true)
	if len(code) == 0 {
		result.Failed(c, int(result.ApiCode.VerificationCodeHasExpired), result.ApiCode.GetMessage(result.ApiCode.VerificationCodeHasExpired))
		return
	}

	// 校验验证码
	verify := CaptVerify(dto.IdKey, dto.Image)
	if !verify {
		result.Failed(c, int(result.ApiCode.CAPTCHANOTTRUE), result.ApiCode.GetMessage(result.ApiCode.CAPTCHANOTTRUE))
		return
	}

	// 校验密码
	sysAdminDetail := dao.SysAdminDetail(dto)
	if sysAdminDetail.Password != util.EncryptionMD5(dto.Password) {
		result.Failed(c, int(result.ApiCode.PASSWORDNOTTRUE), result.ApiCode.GetMessage(result.ApiCode.PASSWORDNOTTRUE))
		return
	}

	// 用户状态检查
	if sysAdminDetail.Status == constant.SYS_ADMIN_STATUS_DISABLE {
		result.Failed(c, int(result.ApiCode.STATUSISENABLE),
			result.ApiCode.GetMessage(result.ApiCode.STATUSISENABLE))
		return
	}

	// 生成token
	tokenString, _ := jwt.GenerateTokenByAdmin(sysAdminDetail)
	result.Success(c, map[string]interface{}{
		"token":       tokenString,
		"systemAdmin": sysAdminDetail,
	})

}

// UpdatePersonal 更新用户个人信息
func (s SysAdminServiceImpl) UpdatePersonal(c *gin.Context, dto entity.UpdatePersonalInfoDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissingModificationOfPersonalParameters),
			result.ApiCode.GetMessage(result.ApiCode.MissingModificationOfPersonalParameters))
		return
	}
	id, _ := jwt.GetAdminId(c)
	dto.Id = id
	result.Success(c, dao.UpdatePersonal(dto))
}

// GetSysAdminList 分页查询用户列表
func (s SysAdminServiceImpl) GetSysAdminList(c *gin.Context, PageSize, PageNum int, Username, Status, BeginTime, EndTime string) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysAdmin, count := dao.GetSysAdminList(PageSize, PageNum, Username, Status,
		BeginTime, EndTime)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize,
		"pageNum": PageNum, "list": sysAdmin})
	return
}

// ResetSysAdminPassword 重置用户密码
func (s SysAdminServiceImpl) ResetSysAdminPassword(c *gin.Context, dto entity.ResetSysAdminPasswordDto) {
	dao.ResetSysAdminPassword(dto)
	result.Success(c, true)
}

// 修改用户状态
func (s SysAdminServiceImpl) UpdateSysAdminStatus(c *gin.Context, dto entity.UpdateSysAdminStatusDto) {
	dao.UpdateSysAdminStatus(dto)
	result.Success(c, true)
}

// 根据id删除用户
func (s SysAdminServiceImpl) DeleteSysAdminById(c *gin.Context, dto entity.SysAdminIdDto) {
	dao.DeleteSysAdminById(dto)
	result.Success(c, true)
}

// 修改用户
func (s SysAdminServiceImpl) UpdateSysAdmin(c *gin.Context, dto entity.UpdateSysAdminInfo) {
	result.Success(c, dao.UpdateSysAdmin(dto))
}

// 根据id查询用户信息
func (s SysAdminServiceImpl) GetSysAdminInfo(c *gin.Context, Id int) {
	result.Success(c, dao.QuerySysUserDetailById(Id))
}

// 新增用户
func (s SysAdminServiceImpl) CreateSysAdmin(c *gin.Context, dto entity.AddSysAdminDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissingNewAdminParameter),
			result.ApiCode.GetMessage(result.ApiCode.MissingNewAdminParameter))
		return }
	bool := dao.CreateSysAdmin(dto)
	if !bool {
		result.Failed(c, int(result.ApiCode.USERNAMEALREADYEXISTS),
			result.ApiCode.GetMessage(result.ApiCode.USERNAMEALREADYEXISTS))
		return }
	result.Success(c, bool)
	return }


// 创建一个全局的SysAdminServiceImpl实例
var sysAdminService = SysAdminServiceImpl{}

// SysAdminService 该函数返回上述创建的SysAdminServiceImpl实例
func SysAdminService() ISysAdminService {
	return &sysAdminService
}
