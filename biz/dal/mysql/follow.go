package mysql

import (
	"douyin_backend/biz/hertz_gen/model/data"
	"douyin_backend/biz/model"
)

func Follow(f *model.Follow) error {
	return DB.Create(f).Error
}

func Unfollow(userId int64, followerId int64) error {
	return DB.Model(&model.Follow{}).Where("user_id = ? and follower_id = ?", userId, followerId).Update("canceled", true).Error
}

func FindFollow(userId int64, followerId int64) (*model.Follow, error) {
	var follow model.Follow
	result := DB.Where("user_id = ? and follower_id = ? and canceled = ?", userId, followerId, false).First(&follow)
	if result.Error != nil {
		return nil, result.Error
	}
	return &follow, nil
}

func GetFollowList(userId int64) ([]*data.User, error) {
	var follows []*model.Follow
	result := DB.Where("user_id = ? and canceled = ?", userId, false).Find(&follows)
	if result.Error != nil {
		return nil, result.Error
	}
	var followerIds []int64
	for _, follow := range follows {
		followerIds = append(followerIds, follow.FollowerId)
	}
	var users []*data.User
	result = DB.Where("id in (?)", followerIds).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func GetFollowerList(userId int64) ([]*data.User, error) {
	var follows []*model.Follow
	result := DB.Where("follower_id = ? and canceled = ?", userId, false).Find(&follows)
	if result.Error != nil {
		return nil, result.Error
	}
	var followerIds []int64
	for _, follow := range follows {
		followerIds = append(followerIds, follow.UserId)
	}
	var users []*data.User
	result = DB.Where("id in (?)", followerIds).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
