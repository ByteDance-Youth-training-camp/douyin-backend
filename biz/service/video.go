package service

import "douyin_backend/biz/hertz_gen/model/data"

type VideoService struct {
}

func (s *VideoService) GetVideoList(userId int64) ([]*data.Video, error) {
	return nil, nil
}

func (s *VideoService) GetVideoById(videoId int64) (data.Video, error) {
	return data.Video{}, nil
}

func (s *VideoService) CreateVideo(userId int64, video *data.Video) (data.Video, error) {
	return data.Video{}, nil
}

func (s *VideoService) DeleteVideo(videoId int64) error {
	return nil
}
