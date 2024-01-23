package util

import (
	"admin-api/middleware/constant"
	"admin-api/pkg/redis"
	"context"
	"log"
	"time"
)

var ctx = context.Background()
type RedisStore struct {

}

// 存验证码，这里是set方法的实现
func (r RedisStore)Set(id string, value string) {
	key := constant.LOGIN_CODE + id
	// 验证码5分钟过期
	err := redis.RedisDb.Set(ctx, key, value, time.Minute * 5).Err()
	if err != nil {
		log.Println(err.Error())
	}
}

// 取验证码
func (r RedisStore)Get(id string, clear bool) string {
	key := constant.LOGIN_CODE + id
	result, err := redis.RedisDb.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return result

}


// 校验验证码
func (c RedisStore) Verify(id, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	return v == answer
}
