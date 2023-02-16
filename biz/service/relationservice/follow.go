package relationservice

import (
	"douyin_backend/biz/dal/redis"
	"douyin_backend/biz/hertz_gen/model/data"
	"douyin_backend/biz/service/userservice"
)

func Follow(userId int64, followedById int64) error {
	// Check if user has already followed
	followed, _, err := redis.FindRelation(userId, followedById)
	if err != nil {
		return err
	}

	if followed {
		return nil
	}

	return redis.AddRelation(userId, followedById)
}

func Unfollow(userId int64, toUserID int64) error {
	followed, _, err := redis.FindRelation(userId, toUserID)
	if err != nil {
		return err
	}

	if !followed {
		return nil
	}

	return redis.RemoveRelation(userId, toUserID)
}

func GetFollowers(userId int64) ([]*data.User, error) {
	users, err := redis.GetFollowerList(userId)
	if err != nil {
		return nil, err
	}

	var res []*data.User

	for _, user := range users {
		usr, err := userservice.GetUserById(user)
		if err != nil {
			return nil, err
		}
		res = append(res, &usr)
	}

	return res, nil
}

func GetFollowings(userId int64) ([]*data.User, error) {
	users, err := redis.GetFollowList(userId)
	if err != nil {
		return nil, err
	}

	var res []*data.User

	for _, user := range users {
		usr, err := userservice.GetUserById(user)
		if err != nil {
			return nil, err
		}
		res = append(res, &usr)
	}

	return res, nil
}

func GetFollowingsCount(userId int64) (int64, error) {
	return redis.GetFollowCount(userId)
}

func GetFollowersCount(userId int64) (int64, error) {
	return redis.GetFollowerCount(userId)
}

// TODO: message
func GetFriends(userId int64) ([]*data.User, error) {
	users, err := redis.GetFriendList(userId)
	if err != nil {
		return nil, err
	}

	var res []*data.User

	for _, user := range users {
		usr, err := userservice.GetUserById(user)
		if err != nil {
			return nil, err
		}
		res = append(res, &usr)
	}

	return res, nil
}
