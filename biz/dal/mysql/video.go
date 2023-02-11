package mysql

import "douyin_backend/biz/model"

func CreateVideo(video *model.Video) (*model.Video, error) {
	result := DB.Create(video)
	return video, result.Error
}

func UpdateVideoKeys(vid int64, videoKey string, coverKey string) error {
	result := DB.Model(&model.Video{ID: vid}).Updates(model.Video{VideoKey: videoKey, CoverKey: coverKey, Ready: true})
	return result.Error
}

func VideoFeed(latest_time *int64) ([]model.Video, error) {

	videos := make([]model.Video, 0)
	result := DB
	if latest_time != nil {
		result = DB.Where("upload_time < ?", latest_time)
	}
	result = result.Where("ready = true").Order("upload_time desc").Limit(30).Find(&videos)
	return videos, result.Error
}
