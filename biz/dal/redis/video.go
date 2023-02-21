package redis

import (
	"douyin_backend/biz/hertz_gen/model/data"
	"encoding/json"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/redis/go-redis/v9"
)

func getVkey(id int64) string {
	return strconv.FormatInt(id, 16)
}
func getUVkey(uid int64, vid int64) string {
	return strconv.FormatInt(uid, 16) + strconv.FormatInt(vid, 16)
}

func GetVideoInfo(vid int64) *data.Video {
	video := &data.Video{}
	result := video_cache.Get(Ctx, getVkey(vid))
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
	result := video_cache.SetEx(Ctx, getVkey(v.ID), data, expired)
	return result.Err()
}

func SetVideoCommentsCnt(vid int64, cnt int64, expired time.Duration) error {
	result := video_comments_cnt_cache.SetEx(Ctx, getVkey(vid), cnt, expired)
	return result.Err()
}
func DeleteVideoCommentsCnt(vid int64) error {
	result := video_comments_cnt_cache.Del(Ctx, getVkey(vid))
	return result.Err()
}
func GetVideoCommentsCnt(vid int64) (int64, error) {
	result := video_comments_cnt_cache.Get(Ctx, getVkey(vid))
	return result.Int64()
}

func GetUserVideoLike(uid int64, vid int64) (bool, error) {
	result := video_user_like_cache.Get(Ctx, getUVkey(uid, vid))
	return result.Bool()
}
func SetUSerVideoLike(uid int64, vid int64, like bool, expired time.Duration) error {
	result := video_user_like_cache.SetEx(Ctx, getUVkey(uid, vid), like, expired)
	return result.Err()
}
func DeleteUserVideoLike(uid int64, vid int64) error {
	result := video_user_like_cache.Del(Ctx, getUVkey(uid, vid))
	return result.Err()
}

func GetVideoLikeCount(vid int64)(int64, error){
	result := video_like_count_cache.Get(Ctx, getVkey(vid))
	return result.Int64()

}
func SetVideoLikeCount(vid int64, count int64, expired time.Duration)error{
	result := video_like_count_cache.SetEx(Ctx, getVkey(vid), count, expired)
	return result.Err()
}
func DeleteVideoLikeCount(vid int64)error{
	result := video_like_count_cache.Del(Ctx, getVkey(vid))
	return result.Err()
}
