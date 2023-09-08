package db

import (
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "",
		DB:       0,
	})
}

func GetRedis() *redis.Client {
	return rdb
}
