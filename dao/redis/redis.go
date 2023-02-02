package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/hxx258456/pyramidel-chain-baas/internal/localconfig"
	"github.com/hxx258456/pyramidel-chain-baas/pkg/utils/logger"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// Init 初始化连接
func Init(redisConfig *localconfig.TopLevel) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			redisConfig.RedisConfig.Host,
			redisConfig.RedisConfig.Port,
		),
		Password: redisConfig.RedisConfig.Password, // no password set
		DB:       redisConfig.RedisConfig.DB,       // use default DB
		PoolSize: redisConfig.RedisConfig.PoolSize,
	})

	if _, err = rdb.Ping().Result(); err != nil {
		logger.Error(err)
	} else {
		logger.Info(">>>Redis连接成功")
	}
	return
}

func Close() {
	_ = rdb.Close()
}
