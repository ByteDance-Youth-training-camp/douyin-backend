package redis

import (
	"douyin_backend/biz/hertz_gen/model/data"
	"encoding/json"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/redis/go-redis/v9"
)

func genVkey(id int64) string {
	return strconv.FormatInt(id, 16)
}

func GetVideoInfo(vid int64) *data.Video {
	video := &data.Video{}
	result := video_cache.Get(Ctx, genVkey(vid))
	if result.Err() != redis.Nil && result.Err() != nil {
		hlog.Debug(result.Err())
		return nil
	}
	data, err := result.Bytes()
	if err != nil {
		hlog.Debug(err)
		return nil
	}
	if err := json.Unmarshal(data, &video); err != nil {
		hlog.Debug(err)
		return nil
	}
	return video
}

func SetVideoInfo(v *data.Video, expired time.Duration) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	result := video_cache.SetEx(Ctx, genVkey(v.ID), data, expired)
	return result.Err()
}
