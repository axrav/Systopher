package db

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DB_URI"),
		Password: "",
		DB:       0,
	})
}
