package config

import (
	"WeiYangWork/global"
	"github.com/go-redis/redis"
	"log"
)

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr: AppConfig.Redis.Address,
		DB:   0,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("failed to connnect to redis, got error: %v ", err)
	}
	global.Redis = RedisClient
}
