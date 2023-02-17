package service

import (
	"douyin_backend/biz/dal/mysql"
	"douyin_backend/biz/dal/redis"
	"douyin_backend/biz/hertz_gen/model/data"
	"douyin_backend/biz/service/videoservice"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type FavoriteService struct {
}

func (s *FavoriteService) AddFavorite(userId int64, videoId int64) error {
	redis.DeleteUserVideoLike(userId, videoId)
	return mysql.AddFavorite(userId, videoId)
}

func (s *FavoriteService) RemoveFavorite(userId int64, videoId int64) error {
	redis.DeleteUserVideoLike(userId, videoId)
	return mysql.RemoveFavorite(userId, videoId)
}

func (s *FavoriteService) GetFavoriteList(userId int64) ([]*data.Video, error) {
	list, err := mysql.GetFavoriteList(userId)
	if err != nil {
		return nil, err
	}
	vlist, err := mysql.VideoList(list)
	if err!=nil{
		hlog.Debug(err)
		return nil, err
	}

	datalist := videoservice.PackVideoList(userId, vlist)
	return datalist, nil
}
