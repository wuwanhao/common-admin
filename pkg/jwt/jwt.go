// JWT工具类(生成token,解析token,获取当前登录的用户id及用户信息)
package jwt

import (
	"admin-api/api/entity"
	"admin-api/middleware/constant"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// jwt参数
type userStdClaims struct {
	entity.JwtAdmin
	jwt.StandardClaims
}

// token过期时间
const tokenExpiredDuration = time.Hour * 24

// token秘钥
var Secret = []byte("admin-api")

var (
	ErrAbsent = "token absent"  // 令牌不存在
	ErrInvalid = "token invalid" // 令牌无效
)

// GenerateTokenByAdmin 根据用户信息生成token
func GenerateTokenByAdmin(admin entity.SysAdmin)(string, error)  {
	var jwtAdmin = entity.JwtAdmin{
		ID: admin.ID,
		Username: admin.Username,
		Nickname: admin.Nickname,
		Icon: admin.Icon,
		Email: admin.Email,
		Phone: admin.Phone,
		Note: admin.Note,
	}

	c := userStdClaims{
		jwtAdmin,                                           // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpiredDuration).Unix(),  // 过期时间
			Issuer: "backstage",                                     // 签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// todo set redis
	// token通过秘钥加密
	return token.SignedString(Secret)
}


// ValidateToken 解析token
func ValidateToken(tokenString string) (*entity.JwtAdmin, error) {
	if tokenString == "" {
		return nil, errors.New(ErrAbsent)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})

	if token == nil {
		return nil, errors.New(ErrInvalid)
	}
	claims := userStdClaims{}
	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return Secret, nil
	})

	if err != nil {
		return nil, err
	}

	return &claims.JwtAdmin, err
}

// 返回id
func GetAdminId(c *gin.Context) (uint, error) {
	u, exists := c.Get(constant.ContextKeyUserObj)
	if !exists {
		return 0, errors.New("Can not get user id")
	}
	admin, ok := u.(*entity.JwtAdmin)
	if ok {
		return admin.ID, nil
	}
	return 0, errors.New("Can not convert id to struct")
}

// 返回用户名
func GetAdminName(c *gin.Context) (string, error) {
	u, exists := c.Get(constant.ContextKeyUserObj)
	if !exists {
		return string(string(0)), errors.New("Can not get user name")
	}
	admin, ok := u.(*entity.JwtAdmin)
	if ok {
		return admin.Username, nil
	}
	return string(string(0)), errors.New("Can not convert username to struct")
}

// 返回admin信息
func GetAdmin(c *gin.Context) (*entity.JwtAdmin, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return nil, errors.New("can't get api")
	}
	admin, ok := u.(*entity.JwtAdmin)
	if ok {
		return admin, nil
	}
	return nil, errors.New("can't convert to api struct")
}