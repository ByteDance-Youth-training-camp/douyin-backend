package redis

import (
	"douyin_backend/biz/dal/mysql"
	"strconv"
)

func genFkey(userId int64) string {
	return strconv.FormatInt(userId, 16)
}

func FindRelation(userId int64, followed_by_id int64) (bool, bool, error) {
	var followed bool
	var follow bool
	result := follow_cache.SIsMember(Ctx, genFkey(userId), followed_by_id)
	if result.Err() != nil {
		// Try find in db
		relation, err := mysql.FindFollow(userId, followed_by_id)
		if err != nil {
			return false, false, err
		}
		if relation != nil && !relation.Canceled {
			AddRelation(userId, followed_by_id)
			followed = true
		}
	}
	followed = result.Val()
	result_rev := follower_cache.SIsMember(Ctx, genFkey(followed_by_id), userId)
	if result_rev.Err() != nil {
		// Try find in db
		relation, err := mysql.FindFollow(followed_by_id, userId)
		if err != nil {
			return false, false, err
		}
		if relation != nil && !relation.Canceled {
			AddRelation(followed_by_id, userId)
			follow = true
		}
	}
	follow = result_rev.Val()
	return followed, follow, nil
}

func AddRelation(userId int64, followed_by_id int64) error {
	result := follow_cache.SAdd(Ctx, genFkey(userId), followed_by_id)
	if result.Err() != nil {
		return result.Err()
	}
	result_rev := follower_cache.SAdd(Ctx, genFkey(followed_by_id), userId)
	return result_rev.Err()
}

func RemoveRelation(userId int64, followed_by_id int64) error {
	result := follow_cache.SRem(Ctx, genFkey(userId), followed_by_id)
	if result.Err() != nil {
		return result.Err()
	}
	result_rev := follower_cache.SRem(Ctx, genFkey(followed_by_id), userId)
	return result_rev.Err()
}

func IsFollowed(userId int64, followed_by_id int64) (bool, error) {
	result := follow_cache.SIsMember(Ctx, genFkey(userId), followed_by_id)
	if result.Err() != nil {
		return false, result.Err()
	}
	return result.Val(), nil
}

func GetFollowList(userId int64) ([]int64, error) {
	result := follow_cache.SMembers(Ctx, genFkey(userId))
	if result.Err() != nil {
		return nil, result.Err()
	}
	intList := result.Val()
	var list []int64
	for _, v := range intList {
		value, err := strconv.ParseInt(v, 16, 64)
		if err != nil {
			return nil, err
		}
		list = append(list, value)
	}
	return list, nil
}

func GetFollowerList(userId int64) ([]int64, error) {
	result := follower_cache.SMembers(Ctx, genFkey(userId))
	if result.Err() != nil {
		return nil, result.Err()
	}
	intList := result.Val()
	var list []int64
	for _, v := range intList {
		value, err := strconv.ParseInt(v, 16, 64)
		if err != nil {
			return nil, err
		}
		list = append(list, value)
	}
	return list, nil
}
