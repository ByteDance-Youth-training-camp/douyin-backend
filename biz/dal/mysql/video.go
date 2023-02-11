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
