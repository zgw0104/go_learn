package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"web_app2/settings"
)

var Rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),

		Password: cfg.Pwd,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	_, err = Rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	Rdb.Close()
}
