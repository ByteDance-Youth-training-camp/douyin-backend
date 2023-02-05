package dal

import (
	"douyin_backend/biz/dal/mysql"
	"douyin_backend/biz/dal/redis"
)

func Init() {
	mysql.Init()
	redis.Init()
}
