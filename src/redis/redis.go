package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	. "tiktok/src/config"
	"time"
)

var RedisCache *redis.Client

func Init() {
	host := AppConfig.Redis.Host
	port := AppConfig.Redis.Port
	password := AppConfig.Redis.Password
	RedisCache = redis.NewClient(&redis.Options{
		DB:       0,
		Addr:     host + ":" + port,
		Password: password,
	})
}
func RCGet(ctx context.Context, key string) (string, error) {
	return RedisCache.Get(ctx, key).Result()
}
func RCExists(ctx context.Context, key string) bool {
	return RedisCache.Exists(ctx, key).Val() != 0
}

func RCSet(ctx context.Context, key string, value interface{}, expiration time.Duration) {
	if RCExists(ctx, key) {
		RedisCache.Expire(ctx, key, expiration)
		return
	}
	RedisCache.Set(ctx, key, value, expiration)
}
