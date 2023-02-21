package videoservice

import (
	"douyin_backend/biz/dal/minio"
	"douyin_backend/biz/dal/mysql"
	"douyin_backend/biz/dal/redis"
	"douyin_backend/biz/hertz_gen/model/data"
	"douyin_backend/biz/model"
	"fmt"
	"io"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type Video struct {
	FReader io.Reader
	Size    int64
	Vid     int64
}

func UploadVideo(video *Video) error {
	key := genVkey(video.Vid)

	err := minio.UploadVideo(key, video.FReader, video.Size)
	if err != nil {
		return err
	}
	coverQue <- coverGenTask{video.Vid, key}
	return nil
}

func genVkey(vid int64) string {
	return fmt.Sprintf("video%d", vid)
}

func PackVideoList(mvlist []model.Video) []*data.Video {
	dvlist := make([]*data.Video, len(mvlist))
	for i := range mvlist {
		dvlist[i] = packVideo(&mvlist[i])
	}
	return dvlist
}

func GetVideoById(vid int64) (*data.Video, error) {
	mv, err := mysql.GetVideoById(vid)
	if err != nil {
		return nil, err
	}
	return packVideo(mv), nil
}

func packVideo(mv *model.Video) *data.Video {
	if v := redis.GetVideoInfo(mv.ID); v != nil {
		return v
	}
	dv := data.Video{
		ID:    mv.ID,
		Title: mv.Title,
	}
	dv.Author = packUser(&mv.User)
	play_url, err := minio.GetVideoUrl(mv.VideoKey, time.Hour)
	if err != nil {
		hlog.Debug(err)
	} else {
		dv.PlayURL = play_url.String()
	}
	cover_url, err := minio.GetImageUrl(mv.CoverKey, time.Hour)
	if err != nil {
		hlog.Debug(err)
	} else {
		dv.CoverURL = cover_url.String()
	}
	dv.Author = packUser(&mv.User)
	// TODO( commentCount, favoriteCount, IsFavorite)
	redis.SetVideoInfo(&dv, time.Minute)
	return &dv

}

func packUser(user *model.User) *data.User {
	duser := data.User{
		ID:            user.ID,
		Name:          user.Username,
		FollowCount:   new(int64),
		FollowerCount: new(int64),
		IsFollow:      false,
	}

	// TODO(Follow & Follower)
	return &duser
}
