package favoriteservice

import (
	"douyin_backend/biz/dal/mysql"
	"douyin_backend/biz/hertz_gen/model/data"
	"douyin_backend/biz/service/videoservice"
)

func AddFavorite(userId int64, videoId int64) error {
	return mysql.AddFavorite(userId, videoId)
}

func RemoveFavorite(userId int64, videoId int64) error {
	return mysql.RemoveFavorite(userId, videoId)
}

func GetFavoriteList(userId int64) ([]*data.Video, error) {
	list, err := mysql.GetFavoriteList(userId)
	if err != nil {
		return nil, err
	}

	var videoList []*data.Video

	for _, video := range list {
		video, err := videoservice.GetVideoById(video)
		if err != nil {
			return nil, err
		}
		videoList = append(videoList, video)
	}

	return videoList, nil
}
