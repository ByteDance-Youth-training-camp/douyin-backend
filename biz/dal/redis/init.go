package redis

import (
	"context"
	"douyin_backend/biz/config"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var RDB *redis.Client

func Init() {

	rdcfg := &config.Cfg.Redis
	RDB = redis.NewClient(&redis.Options{
		Addr:     rdcfg.Address,
		Password: rdcfg.Password,
		DB:       0,
	})
}
