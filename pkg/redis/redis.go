package redis

import (
	"admin-api/common/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var RedisDb *redis.Client

func SetUpRedisDb() error {
	var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Address,
		// Password: config.Config.Redis.Password,
		DB:       9,
	})

	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
