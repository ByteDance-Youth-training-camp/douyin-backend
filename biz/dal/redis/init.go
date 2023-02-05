package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var RDB *redis.Client

func Init() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:36379",
		Password: "",
		DB:       0,
	})
}
