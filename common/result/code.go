// 状态码相关
package result

// Codes 定义的状态
type Codes struct {
	SUCCESS uint
	FAILED uint
	Message map[uint]string
	NOAUTH          uint
	AUTHFORMATERROR uint
	INVALIDTOKEN uint
	MissingLoginParameter      uint
	VerificationCodeHasExpired uint
	CAPTCHANOTTRUE             uint
	PASSWORDNOTTRUE            uint
	STATUSISENABLE             uint
}

// ApiCode 状态码
var ApiCode = &Codes{
	SUCCESS: 200,
	FAILED: 501,
	NOAUTH:  403,
	AUTHFORMATERROR: 405,
	INVALIDTOKEN:               406,
	MissingLoginParameter:      407,
	VerificationCodeHasExpired: 408,
	CAPTCHANOTTRUE:             409,
	PASSWORDNOTTRUE:            410,
	STATUSISENABLE:             411,


}

// 状态信息
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS: "成功",
		ApiCode.FAILED: "失败",
		ApiCode.NOAUTH: "请求头中auth为空",
		ApiCode.AUTHFORMATERROR: "请求头中auth格式有误",
		ApiCode.INVALIDTOKEN: "token无效或者登录过期，请重新登录",
		ApiCode.MissingLoginParameter: "缺少登录参数",
		ApiCode.VerificationCodeHasExpired: "验证码过期",
		ApiCode.CAPTCHANOTTRUE:"验证码错误，请重新输入",
		ApiCode.PASSWORDNOTTRUE: "密码错误",
		ApiCode.STATUSISENABLE: "您的账号已被停用，请联系管理员",
	}
}

// GetMessage 供外部调用
func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
