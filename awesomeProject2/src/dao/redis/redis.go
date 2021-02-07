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
	var redisConfig = config.GlobalConfig.Redis
	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		DB:       redisConfig.Db,
		Password: redisConfig.Password,
		PoolSize: redisConfig.PoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = Rdb.Ping(ctx).Result()
	defer cancel()
	return
}
