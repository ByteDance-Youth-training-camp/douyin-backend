package service

import (
	"douyin_backend/biz/dal/mysql"
	"douyin_backend/biz/hertz_gen/model/data"
)

type FavoriteService struct {
}

func (s *FavoriteService) AddFavorite(userId int64, videoId int64) error {
	return mysql.AddFavorite(userId, videoId)
}

func (s *FavoriteService) RemoveFavorite(userId int64, videoId int64) error {
	return mysql.RemoveFavorite(userId, videoId)
}

func (s *FavoriteService) GetFavoriteList(userId int64) ([]*data.Video, error) {
	list, err := mysql.GetFavoriteList(userId)
	if err != nil {
		return nil, err
	}

	videoService := VideoService{}

	var videoList []*data.Video

	for _, video := range list {
		video, err := videoService.GetVideoById(video)
		if err != nil {
			return nil, err
		}
		videoList = append(videoList, &video)
	}

	return videoList, nil
}
