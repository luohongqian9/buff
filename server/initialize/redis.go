package initialize

import (
	"os"

	"github.com/go-redis/redis"
	"go.uber.org/zap"

	"server/global"
)

func ConnectRedis() redis.Client {
	redisCfg := global.Config.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Address,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		global.Log.Error("redis连接失败", zap.Error(err))
		os.Exit(1)
	}

	return *client
}
