package redis

import (
	"time"

	"github.com/likhithkp/clip/utils/config"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

func NewCache(env *config.Env) *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:         env.RedisAddress,
		Username:     env.RedisUsername,
		Password:     env.RedisPassword,
		DB:           0,
		PoolSize:     10,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	})

	return &Cache{client: rdb}
}

func (c *Cache) GetClient() *redis.Client {
	return c.client
}
