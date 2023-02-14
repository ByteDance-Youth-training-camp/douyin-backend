package redis

import (
	"context"
	"douyin_backend/biz/config"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var video_cache *redis.Client
var follower_cache *redis.Client
var follow_cache *redis.Client

const (
	video_info    = iota
	follower_info = 1
	follow_info   = 2
)

func Init() {
	rdcfg := &config.Cfg.Redis
	video_cache = redis.NewClient(&redis.Options{
		Addr:     rdcfg.Address,
		Password: rdcfg.Password,
		DB:       video_info,
	})

	follow_cache = redis.NewClient(&redis.Options{
		Addr:     rdcfg.Address,
		Password: rdcfg.Password,
		DB:       follow_info,
	})

	follower_cache = redis.NewClient(&redis.Options{
		Addr:     rdcfg.Address,
		Password: rdcfg.Password,
		DB:       follower_info,
	})

	video_cache.FlushDB(Ctx)
}
