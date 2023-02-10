package service

import "douyin_backend/biz/hertz_gen/model/data"

type FollowService struct {
}

func (s *FollowService) Follow(userId int64, followId int64) error {
	return nil
}

func (s *FollowService) Unfollow(userId int64, followId int64) error {
	return nil
}

func (s *FollowService) GetFollowers(userId int64) ([]*data.User, error) {
	return nil, nil
}

func (s *FollowService) GetFollowings(userId int64) ([]*data.User, error) {
	return nil, nil
}

func (s *FollowService) GetFollowingsCount(userId int64) (int64, error) {
	return 0, nil
}

func (s *FollowService) GetFollowersCount(userId int64) (int64, error) {
	return 0, nil
}
