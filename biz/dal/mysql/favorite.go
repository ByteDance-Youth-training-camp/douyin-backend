package mysql

import "douyin_backend/biz/model"

func AddFavorite(userId int64, videoId int64) error {
	result := DB.Create(&model.Favorite{UserId: userId, VideoId: videoId, Canceled: false})
	return result.Error
}

func RemoveFavorite(userId int64, videoId int64) error {
	result := DB.Model(&model.Favorite{}).Where("user_id = ? and video_id = ? and canceled = ?", userId, videoId, false).Update("canceled", true)
	return result.Error
}

func GetFavoriteList(userId int64) ([]int64, error) {
	var favorites []*model.Favorite
	result := DB.Where("user_id = ? and canceled = ?", userId, false).Find(&favorites)
	if result.Error != nil {
		return nil, result.Error
	}
	var videoIds []int64
	for _, favorite := range favorites {
		videoIds = append(videoIds, favorite.VideoId)
	}
	return videoIds, nil
}

func CheckFavorite(userId int64, videoId int64) (bool, error) {
	var cnt int64
	result := DB.Model(&model.Favorite{}).Where("user_id = ? and video_id = ? and canceled = ?", userId, videoId, false).Count(&cnt)
	if result.Error != nil {
		return false, result.Error
	}
	return cnt == 1, nil
}

func GetFavoriteCount(videoId int64) (int64, error) {
	var cnt int64
	result := DB.Model(&model.Favorite{}).Where("video_id = ? and canceled = ?", videoId, false).Count(&cnt)
	return cnt, result.Error
}
