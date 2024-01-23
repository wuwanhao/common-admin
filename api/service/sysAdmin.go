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
}

// 接口实现
type SysAdminServiceImpl struct {

}

// Login 登录
func (s SysAdminServiceImpl)Login(c *gin.Context, dto entity.LoginDto)  {
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
		"token": tokenString,
		"systemAdmin": sysAdminDetail,
	})


}

// 创建一个全局的SysAdminServiceImpl实例
var sysAdminService = SysAdminServiceImpl{}

// SystemAdminService 该函数返回上述创建的SysAdminServiceImpl实例
func SystemAdminService() ISysAdminService {
	return &sysAdminService
}

