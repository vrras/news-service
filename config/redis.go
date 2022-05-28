package config

import (
	"github.com/go-redis/redis"
)

var redisCon *redis.Client

func InitRedis() {
	redisCon = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer redisCon.Close()
}

func GetRedisConnection() *redis.Client {
	InitRedis()

	return redisCon
}
