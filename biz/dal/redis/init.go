package redis

import (
	"context"
	"douyin_backend/biz/config"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var video_cache *redis.Client
var video_comments_cnt_cache *redis.Client
var video_user_like_cache *redis.Client
var video_like_count_cache *redis.Client
var follower_cache *redis.Client
var follow_cache *redis.Client

const (
	video_info_db = iota
	video_comments_cnt_db
	video_user_like_db
	video_like_count_db
	follower_info_db
	follow_info_db
)

func Init() {
	rdcfg := &config.Cfg.Redis
	video_cache = redis.NewClient(&redis.Options{
		Addr:     rdcfg.Address,
		Password: rdcfg.Password,
		DB:       video_info_db,
	})
	video_comments_cnt_cache = redis.NewClient(&redis.Options{
		Addr:     rdcfg.Address,
		Password: rdcfg.Password,
		DB:       video_comments_cnt_db,
	})
	video_user_like_cache = redis.NewClient(&redis.Options{
		Addr:     rdcfg.Address,
		Password: rdcfg.Password,
		DB:       video_user_like_db,
	})
	video_like_count_cache = redis.NewClient(&redis.Options{
		Addr:     rdcfg.Address,
		Password: rdcfg.Password,
		DB:       video_like_count_db,
	})

	follow_cache = redis.NewClient(&redis.Options{
		Addr:     rdcfg.Address,
		Password: rdcfg.Password,
		DB:       follow_info_db,
	})

	follower_cache = redis.NewClient(&redis.Options{
		Addr:     rdcfg.Address,
		Password: rdcfg.Password,
		DB:       follower_info_db,
	})

	video_cache.FlushDB(Ctx)
	video_comments_cnt_cache.FlushDB(Ctx)
	video_user_like_cache.FlushDB(Ctx)
	video_like_count_cache.FlushDB(Ctx)
}
