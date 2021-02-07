package redis

import (
	"awesomeProject/awesomeProject2/src/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var Rdb *redis.Client

func Init() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.GlobalConfig.Redis.Host, config.GlobalConfig.Redis.Port),
		DB:       config.GlobalConfig.Redis.Db,
		Password: config.GlobalConfig.Redis.Password,
		PoolSize: config.GlobalConfig.Redis.PoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = Rdb.Ping(ctx).Result()
	defer cancel()
	return
}
